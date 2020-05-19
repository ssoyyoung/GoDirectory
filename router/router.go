package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	h "github.com/ssoyyoung.p/GoDirectory/handler"
)

//Router function
func Router() *echo.Echo {
	e := echo.New()

	//CORS Middleware
	e.Use(middleware.CORS())

	//Router List
	e.GET("/", h.TestPage)
	e.GET("/getStreamers", h.GetStreamers)
	e.GET("/getStreamer/:id", h.GetStreamerByID)
	e.GET("/deleteStreamer/:id", h.DeleteStreamer)
	e.POST("/updateStreamer/:id", h.UpdateStreamer)
	e.POST("/createStreamer", h.CreateStreamer)

	return e
}
