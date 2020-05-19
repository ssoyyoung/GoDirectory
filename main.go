package main

import (
	router "github.com/ssoyyoung.p/GoDirectory/router"
)

func main() {
	echoR := router.Router()

	// Start echo server..
	echoR.Logger.Fatal(echoR.Start(":1323"))

}
