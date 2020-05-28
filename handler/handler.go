package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	mongodb "github.com/ssoyyoung.p/GoDirectory/mongo"
)

//TestPage func
func TestPage(c echo.Context) error {
	fmt.Println("Connection Complete")
	return nil
}

// GetSteamer func
func GetLiveStreamers(c echo.Context) error {
	res := mongodb.LiveTrueList() //Get live data (desc)

	return c.String(http.StatusOK, res)
}

// GetLiveStreamersByFollower func
func GetLiveStreamersByFollower(c echo.Context) error {
	email := c.Param("email")
	follower := mongodb.SearchDBbyEmail(email)
	res := mongodb.LiveTrueListByFollower(follower) //Get Follower data

	return c.String(http.StatusOK, res)
}

// GetStreamers func
func GetStreamers(c echo.Context) error {
	res := mongodb.LiveAllList() //Get all data

	return c.String(http.StatusOK, res)
}

// GetStreamerByID func
func GetStreamerByID(c echo.Context) error {
	id := c.Param("id")
	res := mongodb.SearchDBbyID(id)

	return c.String(http.StatusOK, res)
}

// DeleteStreamer func
func DeleteStreamer(c echo.Context) error {
	id := c.Param("id")
	res := mongodb.DeleteDBbyID(id)

	return c.String(http.StatusOK, res)
}

// UpdateStreamer func
func UpdateStreamer(c echo.Context) error {
	id := c.Param("id")
	res := mongodb.UpdateDBbyID(id, c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))

	return c.String(http.StatusOK, res)
}

// CreateStreamer func
func CreateStreamer(c echo.Context) error {
	res := mongodb.CreateDB(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))

	return c.String(http.StatusOK, res)
}

// PushFollowing func
func PushFollowing(c echo.Context) error {
	email := c.Param("email")
	res := mongodb.PushFollowing(c.FormValue("following"), email)

	return c.String(http.StatusOK, res)
}

// PullFollowing func
func PullFollowing(c echo.Context) error {
	email := c.Param("email")
	res := mongodb.PullFollowing(c.FormValue("following"), email)

	return c.String(http.StatusOK, res)
}

// GetFollowing func
func GetFollowing(c echo.Context) error {
	email := c.Param("email")

	result := strings.Join(mongodb.SearchDBbyEmail(email), ",")

	return c.String(http.StatusOK, result)
}
