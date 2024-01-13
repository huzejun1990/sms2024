package main

import (
	"github.com/gin-gonic/gin"
	_ "sms2024/redis"
	//"net/http"
	"sms2024/web"
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
