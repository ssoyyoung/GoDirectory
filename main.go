package main

import (
	router "github.com/ssoyyoung.p/GoDirectory/router"
)

// @title MeerkatOnAir API Server
// @version 0.0.1
// @description We Are Team Sparker

// @contact.name soyoung Park
// @contact.url http://www.sparker.kr
// @contact.email cracker.weare@gmail.com

func main() {
	echoR := router.Router()

	// Start echo server..
	echoR.Logger.Fatal(echoR.Start(":1323"))

}
