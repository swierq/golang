package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/headers", printHeaders)
	e.GET("/setcookie", writeCookie)
	e.GET("/cookies", printCookies)
	e.Logger.Fatal(e.Start(":1323"))
}

func printHeaders(c echo.Context) error {
	headers := ""
	for k, v := range c.Request().Header {
		headers = fmt.Sprintf("%s\n%s: %s", headers, k, v)
	}
	return c.String(http.StatusOK, headers)
}

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "test"
	cookie.Expires = time.Now().Add(1 * time.Minute)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "cookie written")
}

func printCookies(c echo.Context) error {
	cookies := c.Cookies()
	if len(cookies) == 0 {
		return c.String(http.StatusOK, "No cookies found")
	}
	cookieList := ""
	for _, cookie := range cookies {
		cookieList = fmt.Sprintf("%s\n%s: %s", cookieList, cookie.Name, cookie.Value)
	}
	return c.String(http.StatusOK, cookieList)
}
