package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mongodb "github.com/ssoyyoung.p/GoDirectory/mongo"
	"github.com/ssoyyoung.p/GoDirectory/utils"
)

// IsLoggedIn FUNC
var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

type admin struct {
	AdminUser []string //대문자로 써야 인식이 된다.
}

func getAdminUser() admin {
	data, err := os.Open("auth/auth.json")
	utils.CheckErr(err)

	var ad admin
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &ad)

	return ad
}

func contains(v string, a []string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

// SignUp func
func SignUp(c echo.Context) error {
	id, password, nickname, birthday, tags := c.FormValue("id"), c.FormValue("password"), c.FormValue("nickname"), c.FormValue("birthday"), c.FormValue("tags")
	tagList := strings.Split(tags, ",")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["nickname"] = nickname
	claims["id"] = id
	claims["admin"] = false

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	mongodb.SignUp(id, password, nickname, birthday, t, tagList)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":     t,
		"token_exp": claims["exp"],
		"nickname":  nickname,
		"id":        id,
	})
}

// GoogleLogin func
func GoogleLogin(c echo.Context) error {
	googleID, name, email := c.FormValue("googleId"), c.FormValue("name"), c.FormValue("email")

	if mongodb.CheckUser(googleID, name, email) {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = name
		claims["googleId"] = googleID
		claims["admin"] = false

		admin := getAdminUser()
		switch {
		case contains(email, admin.AdminUser):
			claims["admin"] = true
		}

		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		mongodb.UpdateUser(googleID, t)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":     t,
			"token_exp": claims["exp"],
			"name":      name,
			"email":     email,
			"googleID":  googleID,
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
