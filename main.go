package main

import (
	"fmt"

	router "github.com/ssoyyoung.p/GoDirectory/router"
)

// @title MeerkatOnAir API Server
// @version 0.0.1
// @description We Are Team Sparker

// @contact.name soyoung Park
// @contact.url http://www.sparker.kr
// @contact.email cracker.weare@gmail.com

func main() {
	debug := false
	echoR := router.Router()

	fmt.Println("Start echo server..")

	if debug {
		echoR.Logger.Fatal(echoR.Start(":1323"))
	} else {
		echoR.Logger.Fatal(echoR.StartTLS(":1323", "cert.pem", "privkey.pem"))
	}
}
