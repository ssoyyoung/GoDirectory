package cookie

import (
	"fmt"
	"net/http"
	"time"
	"github.com/labstack/echo"
)

// GetCookie func
func GetCookie(c echo.Context) error {
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
