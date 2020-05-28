package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	auth "github.com/ssoyyoung.p/GoDirectory/auth"

	//_ "github.com/ssoyyoung.p/GoDirectory/docs"
	handler "github.com/ssoyyoung.p/GoDirectory/handler"
)

/* //Router function
func Router() *echo.Echo {
	e := echo.New()

	e.Debug = true

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
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Main Page Request(order by views)
	e.GET("/getList", handler.GetLiveStreamers)
	//Main Page Request(order by follwer)
	e.GET("/getList/:email", handler.GetLiveStreamersByFollower)

	//Admin Page Request
	e.GET("/getStreamers", handler.GetStreamers)
	e.GET("/getStreamer/:id", handler.GetStreamerByID)
	e.GET("/deleteStreamer/:id", handler.DeleteStreamer, auth.IsLoggedIn)
	e.POST("/updateStreamer/:id", handler.UpdateStreamer, auth.IsLoggedIn)
	e.POST("/createStreamer", handler.CreateStreamer, auth.IsLoggedIn)

	// following Request
	e.POST("/pushFollowing/:email", handler.PushFollowing)
	e.POST("/pullFollowing/:email", handler.PullFollowing)
	e.GET("/getFollowing/:email", handler.GetFollowing)

	//Login Request > return으로 following 정보 보내기
	e.POST("/login", auth.GoogleLogin)

	return e
}
*/

// TODO
// userInfo > login
// updateFollower > pushFollowing, param 이름 follower > following
// pullFollowing 새로 생성함

// Router function
func Router() *echo.Echo {
	e := echo.New()
	e.Debug = true

	// echo middleware func
	e.Use(middleware.Logger())                             //Setting logger
	e.Use(middleware.Recover())                            //Recover from panics anywhere in the chain
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //CORS Middleware
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Router List
	getList := e.Group("getList")
	{
		getList.GET("/live", handler.GetLiveStreamers)
		getList.GET("/following/:email", handler.GetLiveStreamersByFollower)
	}

	admin := e.Group("/admin")
	{
		admin.GET("/getStremers", handler.GetStreamers)
		admin.GET("/getStreamer/:id", handler.GetStreamerByID)
		admin.GET("/deleteStreamer/:id", handler.DeleteStreamer, auth.IsLoggedIn)
		admin.POST("/updateStreamer/:id", handler.UpdateStreamer, auth.IsLoggedIn)
		admin.POST("/createStreamer", handler.CreateStreamer, auth.IsLoggedIn)
	}

	following := e.Group("/following")
	{
		following.POST("/updateUserInfo/:email", handler.PushFollowing)
		following.POST("/deleteUserInfo/:email", handler.PullFollowing)
		following.GET("/getUserInfo/:email", handler.GetFollowing)
	}

	login := e.Group("/login")
	{
		login.POST("", auth.GoogleLogin)
	}

	return e
}
