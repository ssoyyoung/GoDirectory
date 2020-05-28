package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	handler "github.com/ssoyyoung.p/GoDirectory/handler"
	login "github.com/ssoyyoung.p/GoDirectory/login"
)

//Router function
func Router() *echo.Echo {
	e := echo.New()

	//Setting logger
	e.Use(middleware.Logger())
	//Recover from panics anywhere in the chain
	e.Use(middleware.Recover())

	//CORS Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//Router List
	e.GET("/", handler.TestPage)

	//Main Page Request(order by views)
	e.GET("/getList", handler.GetLiveStreamers)
	//Main Page Request(order by follwer)
	e.GET("/getList/:email", handler.GetLiveStreamersByFollower)

	//Admin Page Request
	e.GET("/getStreamers", handler.GetStreamers)
	e.GET("/getStreamer/:id", handler.GetStreamerByID)
	e.GET("/deleteStreamer/:id", handler.DeleteStreamer, login.IsLoggedIn)
	e.POST("/updateStreamer/:id", handler.UpdateStreamer, login.IsLoggedIn)
	e.POST("/createStreamer", handler.CreateStreamer, login.IsLoggedIn)

	// following Request
	e.POST("/pushFollowing/:email", handler.PushFollowing)
	e.POST("/pullFollowing/:email", handler.PullFollowing)
	e.GET("/getFollowing/:email", handler.GetFollowing)

	//Login Request > return으로 following 정보 보내기
	e.POST("/login", login.GoogleLogin)

	// TODO
	// userInfo > login
	// updateFollower > pushFollowing, param 이름 follower > following
	// pullFollowing 새로 생성함

	return e
}
