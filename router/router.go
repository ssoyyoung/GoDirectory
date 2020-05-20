package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	handler "github.com/ssoyyoung.p/GoDirectory/handler"
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
	e.GET("/getStreamers", handler.GetStreamers)
	e.GET("/getStreamer/:id", handler.GetStreamerByID)
	e.GET("/deleteStreamer/:id", handler.DeleteStreamer)
	e.POST("/updateStreamer/:id", handler.UpdateStreamer)
	e.POST("/createStreamer", handler.CreateStreamer)

	return e
}
