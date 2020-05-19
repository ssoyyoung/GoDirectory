package main

import (
    router "github.com/ssoyyoung.p/tree/router"
)

func main() {
	echo_r := router.Router()

    // Start echo server..
    echo_r.Logger.Fatal(echo_r.Start(":1323"))

}