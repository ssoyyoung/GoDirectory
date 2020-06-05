package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	auth "github.com/ssoyyoung.p/GoDirectory/auth"
	//"golang.org/x/crypto/acme/autocert"

	//_ "github.com/ssoyyoung.p/GoDirectory/docs"
	handler "github.com/ssoyyoung.p/GoDirectory/handler"
)

// Router function
func Router() *echo.Echo {
	e := echo.New()
	e.Debug = true

	// echo middleware func
	//e.Pre(middleware.HTTPSRedirect())
	//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("mkoa.sparker.kr")
	//e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Logger())                             //Setting logger
	e.Use(middleware.Recover())                            //Recover from panics anywhere in the chain
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //CORS Middleware
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Health check!
	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	// Router List
	getList := e.Group("getList")
	{
		getList.GET("/live", handler.GetLiveStreamers)
		getList.GET("/live/:category", handler.GetLiveStreamersByCate)
		getList.GET("/following/:email", handler.GetLiveStreamersByFollower)
	}

	admin := e.Group("/admin")
	{
		admin.GET("/getStreamers", handler.GetStreamers)
		admin.GET("/getStreamer/:id", handler.GetStreamerByID)
		admin.GET("/deleteStreamer/:id", handler.DeleteStreamer, auth.IsLoggedIn, auth.IsAdmin)
		admin.POST("/updateStreamer/:id", handler.UpdateStreamer, auth.IsLoggedIn, auth.IsAdmin)
		admin.POST("/createStreamer", handler.CreateStreamer, auth.IsLoggedIn, auth.IsAdmin)
	}

	following := e.Group("/following")
	{
		following.POST("/updateUserInfo/:email", handler.PushFollowing)
		following.POST("/deleteUserInfo/:email", handler.PullFollowing)
		following.GET("/getUserInfo/:email", handler.GetFollowing)
	}

	unfollowing := e.Group("/block")
	{
		unfollowing.POST("/updateUserInfo/:email", handler.PushBlocking)
		unfollowing.POST("/deleteUserInfo/:email", handler.PullBlocking)
		unfollowing.GET("/getUserInfo/:email", handler.GetBlocking)
	}

	login := e.Group("/login")
	{
		login.POST("", auth.GoogleLogin)
	}

	return e
}
