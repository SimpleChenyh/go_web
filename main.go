package main

import (
	"chenyh.com/go_web/router"
	"github.com/gin-gonic/gin"
	"log"
)


func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	router.InitUserRouter(r)

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
