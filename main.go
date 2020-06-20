package main

import (
	router "github.com/ssoyyoung.p/GoDirectory/router"
	"fmt"
)

// @title MeerkatOnAir API Server
// @version 0.0.1
// @description We Are Team Sparker

// @contact.name soyoung Park
// @contact.url http://www.sparker.kr
// @contact.email cracker.weare@gmail.com

func main() {
	debug := true

	echoR := router.Router()

	fmt.Println("Start echo server..")

	if debug {
		echoR.Logger.Fatal(echoR.Start(":1324"))
	} else {
	echoR.Logger.Fatal(echoR.StartTLS(":1323", "cert.pem", "privkey.pem"))
	}
}
