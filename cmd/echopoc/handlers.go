package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/goforj/godump"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/swierq/golang/internal/uihtmx"
	"github.com/swierq/golang/internal/uihtmx/ui/layout"
	"golang.org/x/oauth2"
)

var ( // app is the main application instance
	menu = layout.Menu{Items: []layout.MenuItem{
		layout.MenuItem{Title: "Set Cookie", Path: "/setcookie"},
		layout.MenuItem{Title: "Cookies", Path: "/cookies"},
		layout.MenuItem{Title: "Headers", Path: "/headers"},
		layout.MenuItem{Title: "Login", Path: "/login"},
		layout.MenuItem{Title: "Token", Path: "/token"},
	}}
)

func (a *app) printHeaders(c echo.Context) error {
	headers := ""
	for k, v := range c.Request().Header {
		headers = fmt.Sprintf("%s\n%s: %s\n", headers, k, v)
	}
	return uihtmx.RenderPage(c.Response().Writer, TextPage(nl2br(headers)), menu, "Cookies", "Description")
}

func (a *app) homeHandler(c echo.Context) error {
	return uihtmx.RenderPage(c.Response().Writer, TextPage("Home"), menu, "Cookies", "Description")
}

func (a *app) writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "test"
	cookie.Expires = time.Now().Add(1 * time.Minute)
	c.SetCookie(cookie)
	return uihtmx.RenderPage(c.Response().Writer, TextPage("cookie written"), menu, "Cookies", "Description")
}

func (a *app) printCookies(c echo.Context) error {
	cookies := c.Cookies()
	if len(cookies) == 0 {
		return uihtmx.RenderPage(c.Response().Writer, TextPage("no cookies set"), menu, "Cookies", "Description")
	}
	cookieList := ""
	for _, cookie := range cookies {
		cookieList = fmt.Sprintf("%s\n%s: %s", cookieList, cookie.Name, cookie.Value)
	}
	return uihtmx.RenderPage(c.Response().Writer, TextPage(cookieList), menu, "Cookies", "Description")
}

func nl2br(text string) string {
	return strings.ReplaceAll(text, "\n", "<br>")
}

func (a *app) loginHandler(c echo.Context) error {
	url := a.entra.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *app) logoutHandler(c echo.Context) error {
	// Clear cookies
	cookies := c.Cookies()
	for _, cookie := range cookies {
		deleteCookie := &http.Cookie{
			Name:    cookie.Name,
			Value:   "",
			Expires: time.Unix(0, 0),
			Path:    "/",
		}
		c.SetCookie(deleteCookie)
	}

	// Redirect to landing page
	return c.Redirect(http.StatusSeeOther, "/")
}

func (a *app) callbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "code parameter missing"})
	}

	token, err := a.entra.Exchange(c.Request().Context(), code)
	if err != nil {
		a.e.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to exchange code"})
	}
	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = token.AccessToken
	cookie.Expires = time.Now().Add(60 * time.Minute)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusTemporaryRedirect, "/token")

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"access_token": token.AccessToken,
	// 	"token_type":   token.TokenType,
	// 	"expiry":       token.Expiry,
	// 	"valid":        token.Valid(),
	// })
}

func (a *app) tokenHandler(c echo.Context) error {
	token, err := c.Cookie("auth")
	if err != nil {
		return uihtmx.RenderPage(c.Response().Writer, TextPage("no token in cookie"), menu, "Token", "Description")
	}

	// Get the JWKS URL.
	jwksURL := fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/v2.0/keys", a.config.TenantID)

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Create the keyfunc options. Use an error handler that logs. Refresh the JWKS when a JWT signed by an unknown KID
	// is found or at the specified interval. Rate limit these refreshes. Timeout the initial JWKS refresh request after
	// 10 seconds. This timeout is also used to create the initial context.Context for keyfunc.Get.
	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		a.e.Logger.Error("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}
	defer jwks.EndBackground()
	var vtoken *jwt.Token
	if vtoken, err = jwt.Parse(token.Value, jwks.Keyfunc); err != nil {
		a.e.Logger.Error("Failed to parse the JWT.\nError: %s", err.Error())
	}

	// Check if the token is valid.
	if !vtoken.Valid {
		a.e.Logger.Error("The token is not valid.")
	}

	data := godump.DumpHTML(vtoken)

	return uihtmx.RenderPage(c.Response().Writer, TextPage(data), menu, "Cookies", "Description")
}
