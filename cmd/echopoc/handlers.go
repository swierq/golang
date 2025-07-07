package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/swierq/golang/internal/uihtmx"
	"github.com/swierq/golang/internal/uihtmx/ui/layout"
)

var ( // app is the main application instance
	menu = layout.Menu{Items: []layout.MenuItem{
		layout.MenuItem{Title: "Set Cookie", Path: "/setcookie"},
		layout.MenuItem{Title: "Cookies", Path: "/cookies"},
		layout.MenuItem{Title: "Headers", Path: "/headers"},
	}}
)

func (a *app) printHeaders(c echo.Context) error {
	headers := ""
	for k, v := range c.Request().Header {
		headers = fmt.Sprintf("%s\n%s: %s\n", headers, k, v)
	}
	return uihtmx.RenderPage(c.Response().Writer, TextPage(nl2br(headers)), menu, "Cookies", "Description")
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
	return strings.Replace(text, "\n", "<br>", -1)
}
