package cookie

import (
	"github.com/labstack/echo"
)

// GetCookie func
func GetCookie(c echo.Context) string {
	cookieValue := ""
	for _, cookie := range c.Cookies() {
		if cookie.Name == "mkoaUID" {
			cookieValue = cookie.Value
		}
	}

	return cookieValue

}
