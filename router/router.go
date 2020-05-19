package router

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    mw "github.com/ssoyyoung.p/tree/middleware"
)

func Router( ) *echo.Echo {
    e := echo.New()

    //CORS Middleware
    e.Use(middleware.CORS())

    //Router List
    e.GET("/", mw.TestPage)

    return e
}