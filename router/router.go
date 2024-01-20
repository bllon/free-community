package router

import (
	"free-community/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping")
	})

	router.POST("user/register", handler.UserRegister)
	router.GET("user/info", handler.UserInfo)
}
