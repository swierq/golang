package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/swierq/golang/internal/uihtmx"
	"github.com/swierq/golang/internal/uihtmx/ui"
)

func main() {
	assets, _ := fs.Sub(ui.Assets, "assets")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/headers", printHeaders)
	e.GET("/setcookie", writeCookie)
	e.GET("/cookies", printCookies)
	e.StaticFS("/assets/", assets)
	e.Logger.Fatal(e.Start(":1323"))
}

func printHeaders(c echo.Context) error {
	headers := ""
	for k, v := range c.Request().Header {
		headers = fmt.Sprintf("%s\n%s: %s\n", headers, k, v)
	}
	return uihtmx.RenderPage(c.Response().Writer, TextPage(nl2br(headers)), "Cookies", "Description")
}

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "test"
	cookie.Expires = time.Now().Add(1 * time.Minute)
	c.SetCookie(cookie)
	return uihtmx.RenderPage(c.Response().Writer, TextPage("cookie written"), "Cookies", "Description")
}

func printCookies(c echo.Context) error {
	cookies := c.Cookies()
	if len(cookies) == 0 {
		return uihtmx.RenderPage(c.Response().Writer, TextPage("no cookies set"), "Cookies", "Description")
	}
	cookieList := ""
	for _, cookie := range cookies {
		cookieList = fmt.Sprintf("%s\n%s: %s", cookieList, cookie.Name, cookie.Value)
	}
	return uihtmx.RenderPage(c.Response().Writer, TextPage(cookieList), "Cookies", "Description")
}

func nl2br(text string) string {
	return strings.Replace(text, "\n", "<br>", -1)
}
