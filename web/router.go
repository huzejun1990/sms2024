// @Author huzejun 2024/1/13 13:25:00
package web

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	smsGroup := r.Group("/sms")
	smsGroup.POST("/send", Send)
}
