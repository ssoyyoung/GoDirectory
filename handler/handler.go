package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	mongodb "github.com/ssoyyoung.p/GoDirectory/mongo"
	"github.com/ssoyyoung.p/GoDirectory/utils"
)

// GetLiveStreamers func
// @Summary get all live true data
// @name GetLiveStreamers
// @Router /getList [get]
// @Success 200
func GetLiveStreamers(c echo.Context) error {
	res := mongodb.LiveTrueList() //Get live data (desc)

	return c.String(http.StatusOK, res)
}

// GetLiveStreamersByFollower func
func GetLiveStreamersByFollower(c echo.Context) error {
	email := c.Param("email")
	follower := mongodb.GetFollowing(email)
	// follower가 없는경우
	if follower == nil {
		return c.String(http.StatusOK, "")
	}
	res := mongodb.LiveTrueListByFollower(follower) //Get Follower data

	return c.String(http.StatusOK, res)
	//return c.JSON(http.StatusOK, res)
}

// GetLiveStreamersByBlocking func
func GetLiveStreamersByBlocking(c echo.Context) error {
	email := c.Param("email")
	blocking := mongodb.GetBlocking(email)
	// blocking가 없는경우
	if blocking == nil {
		return c.String(http.StatusOK, "")
	}
	res := mongodb.LiveAllListByBlocking(blocking)

	return c.String(http.StatusOK, res)
}

// GetStreamers func
func GetStreamers(c echo.Context) error {
	res := mongodb.AllList() //Get all data  > crawl to live

	return c.String(http.StatusOK, res)
}

// GetAllStreamers func
func GetAllStreamers(c echo.Context) error {
	res := mongodb.AllList() //Get all data  > crawl to live

	return c.String(http.StatusOK, res)
}

// GetLiveStreamersByCate func
func GetLiveStreamersByCate(c echo.Context) error {
	cate := c.Param("category")
	res := mongodb.GetCategoryList(cate)

	return c.String(http.StatusOK, res)
}

// GetStreamersByPlatform func
func GetStreamersByPlatform(c echo.Context) error {
	platform := c.Param("platform")
	res := mongodb.GetPlatformList(platform)

	return c.String(http.StatusOK, res)
}

// GetAllCategory func
func GetAllCategory(c echo.Context) error {
	cateList := []string{"GAME", "MUSIC", "CHATTING", "SHOPPING", "NEWS & INFO", "AIR & RADIO", "SPORTS & EXERCISE", "COOKING"}

	result := map[string]int{}
	for _, cate := range cateList {
		res := mongodb.GetCategoryCount(cate)
		result[cate] = res
	}

	return c.JSON(http.StatusOK, result)
}

// GetStreamerByID func
func GetStreamerByID(c echo.Context) error {
	id := c.Param("id")
	res := mongodb.SearchDBbyID(id) //  > crawl to live

	return c.String(http.StatusOK, res)
}

// DeleteStreamer func
func DeleteStreamer(c echo.Context) error {
	id := c.Param("id")
	res := mongodb.DeleteDBbyID(id) //  > crawl to live

	return c.String(http.StatusOK, res)
}

// UpdateStreamer func
func UpdateStreamer(c echo.Context) error {
	id := c.Param("id") //  > crawl to live
	res := mongodb.UpdateDBbyID(id, c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))

	return c.String(http.StatusOK, res)
}

// CreateStreamer func
func CreateStreamer(c echo.Context) error {

	res := mongodb.CreateDB(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"), c.FormValue("category"))

	return c.String(http.StatusOK, res)
}

// ExistStreamer func
func ExistStreamer(c echo.Context) error {

	res := mongodb.CheckDB(c.FormValue("platform"), c.FormValue("channelID"))

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

	result := strings.Join(mongodb.GetFollowing(email), ",")

	return c.String(http.StatusOK, result)
}

// PushBlocking func
func PushBlocking(c echo.Context) error {
	email := c.Param("email")
	res := mongodb.PushBlocking(c.FormValue("block"), email)

	return c.String(http.StatusOK, res)
}

// PullBlocking func
func PullBlocking(c echo.Context) error {
	email := c.Param("email")
	res := mongodb.PullBlocking(c.FormValue("block"), email)

	return c.String(http.StatusOK, res)
}

// GetBlocking func
func GetBlocking(c echo.Context) error {
	email := c.Param("email")

	result := strings.Join(mongodb.GetBlocking(email), ",")

	return c.String(http.StatusOK, result)
}

// SearchBar func
func SearchBar(c echo.Context) error {
	query := c.Param("query")
	res := mongodb.SearchBar(query)

	return c.String(http.StatusOK, res)
}

// GetScheduleListByDays func
func GetScheduleListByDays(c echo.Context) error {
	month := c.QueryParam("month")
	day := c.QueryParam("day")
	hours := c.QueryParam("hour")

	// month := c.FormValue("month")
	// day := c.FormValue("days")
	// hours := c.FormValue("hours")

	if month == "" && day == "" {
		return c.String(http.StatusOK, "Query Parameter is empty")
	}

	m, error := strconv.Atoi(month)
	utils.CheckErr(error)
	d, error := strconv.Atoi(day)
	utils.CheckErr(error)

	res := mongodb.GetScheduleListByDay(m, d, hours)

	return c.String(http.StatusOK, res)
}

// GetScheduleList func
func GetScheduleList(c echo.Context) error {
	res := mongodb.AllScheduleList()

	return c.String(http.StatusOK, res)
}

// InsertFeedback func
func InsertFeedback(c echo.Context) error {
	//cookie := cookie.GetCookie(c)
	cookie := "cookie"
	title := c.FormValue("title")
	email := c.FormValue("email")
	message := c.FormValue("message")

	res := mongodb.InsertFeedback(cookie, title, email, message)
	return c.String(http.StatusOK, res)
}

// InsertUserHistory func
func InsertUserHistory(c echo.Context) error {
	//cookie := cookie.GetCookie(c)
	loginType := c.FormValue("loginType")
	username := c.FormValue("username")
	pathname := c.FormValue("pathname")
	residencetime, _ := strconv.Atoi(c.FormValue("residencetime"))

	res := mongodb.InsertUserHistory(loginType, username, pathname, residencetime)
	return c.String(http.StatusOK, res)
}

// InsertViewHistory func
func InsertViewHistory(c echo.Context) error {
	//cookie := cookie.GetCookie(c)
	loginType := c.FormValue("loginType")
	username := c.FormValue("username")
	streaming := c.FormValue("streaming")
	platform := c.FormValue("platform")
	_uniq := c.FormValue("_uniq")

	res := mongodb.InsertViewHistory(loginType, username, streaming, platform, _uniq)
	return c.String(http.StatusOK, res)
}

// CheckID func
func CheckID(c echo.Context) error {
	type returnData struct {
		Status   string
		SerialNo string
	}

	userID := c.FormValue("userID")

	status, res := mongodb.CheckID(userID)
	fmt.Println(status, res)

	rData := returnData{
		Status:   status,
		SerialNo: res,
	}
	//return c.String(http.StatusOK, status+res)
	return c.JSON(http.StatusOK, rData)
}
