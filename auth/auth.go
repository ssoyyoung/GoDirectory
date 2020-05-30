package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mongodb "github.com/ssoyyoung.p/GoDirectory/mongo"
)

// IsLoggedIn FUNC
var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

// GoogleLogin func
func GoogleLogin(c echo.Context) error {
	googleID, name, email := c.FormValue("googleId"), c.FormValue("name"), c.FormValue("email")

	if mongodb.CheckUser(googleID, name, email) {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = name
		claims["googleId"] = googleID
		claims["admin"] = false
		switch email {
		case
			"jwhyun2215@gmail.com",
			"truenorthj@gmail.com",
			"inajung.korea@gmail.com",
			"ssoyyoung.p@gmail.com",
			"cracker.weare@gmail.com":
			claims["admin"] = true
		}

		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		mongodb.UpdateUser(googleID, t)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":    t,
			"name":     name,
			"email":    email,
			"googleID": googleID,
		})
	}
	return echo.ErrUnauthorized
}

// IsAdmin func
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		fmt.Println(isAdmin)
		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
