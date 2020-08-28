package cookie

import (
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