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

	// follow & like Request
	e.POST("/updateFollower/:email", handler.UpdateFollower)

	//Login Request
	e.POST("/userInfo", login.GoogleLogin)
	// TODO : chang userInfo > login

	return e
}
