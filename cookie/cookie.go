package cookie

import (
	"fmt"
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
