package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"url_shortner/lib"
)

var goPort = "80"

func main() {
	fmt.Println("Hello, world.")
	fmt.Println("Running on host:" + lib.CurrentDns + ":" + goPort + "/")

	router := gin.Default()
	router.POST("/api/shortenUrl", lib.ShortenNewUrl)

	router.GET("/url", lib.GetLongUrl)

	router.GET("/", lib.Hello)

	// Run the server on port GO_PORT
	err1 := router.Run(":" + goPort)
	if err1 != nil {
		panic(err1)
	}
}
