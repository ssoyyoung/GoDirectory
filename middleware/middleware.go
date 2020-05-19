package middleware

import (
    "fmt"
    "github.com/labstack/echo"
)

//TestPage func
func TestPage(c echo.Context) error {
    fmt.Println("okay tree folder")
    return nil
}
