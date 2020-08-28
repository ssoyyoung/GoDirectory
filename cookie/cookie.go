package cookie

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// WriteCookie func
func WriteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "uid"
	cookie.Value = "test"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie")
}

// ReadCookie func
func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("uid")
	if err != nil {
		return err
	}
	fmt.Println(cookie)
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	return c.String(http.StatusOK, "read a cookie")

}