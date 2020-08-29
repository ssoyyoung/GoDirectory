package cookie

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GetCookie func
func GetCookie(c echo.Context) string {
	cookieValue := ""
	fmt.Println(c.Cookie("mkoaUID"))
	for _, cookie := range c.Cookies() {
		if cookie.Name == "mkoaUID" {
			cookieValue = cookie.Value
		}
	}

	return cookieValue
}

// GetCookieTest func
func GetCookieTest(c echo.Context) error {
	cName := ""
	cValue := ""
	for _, cookie := range c.Cookies() {
		if cookie.Name == "mkoaUID" {
			cName = cookie.Name
			cValue = cookie.Value
		}
	}

	if cName != "" {
		fmt.Println(cName, cValue)
	}

	return c.String(http.StatusOK, "read a cookie")
}
