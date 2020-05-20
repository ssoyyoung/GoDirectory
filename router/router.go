package router

import (
	"net/http"

	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	h "github.com/ssoyyoung.p/GoDirectory/handler"
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
	e.GET("/", h.TestPage)
	e.GET("/getStreamers", h.GetStreamers)
	e.GET("/getStreamer/:id", h.GetStreamerByID)
	e.GET("/deleteStreamer/:id", h.DeleteStreamer)
	e.POST("/updateStreamer/:id", h.UpdateStreamer)
	e.POST("/createStreamer", h.CreateStreamer)

	return e
}
