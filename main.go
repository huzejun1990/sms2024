package main

import (
	//"net/http"
	"sms2024/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	web.InitRouter(r)
	/*	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})*/
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
