package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	auth "github.com/ssoyyoung.p/GoDirectory/auth"

	//_ "github.com/ssoyyoung.p/GoDirectory/docs"
	handler "github.com/ssoyyoung.p/GoDirectory/handler"
)

// Router function
// GET은 가져오는 것, POST는 수행하는 것
// GET은 주소 뒤에 parameter가 붙고, POST는 body에 감싸져서 데이터를 보낸다.
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

	search := e.Group("/search")
	{
		search.GET("/:query", handler.SearchBar)
	}

	return e
}
