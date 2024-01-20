package handler

import (
	"free-community/logic"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	resp, err := logic.NewHandler().UserRegister()
	if err != nil {
		c.AbortWithError(200, err)
	}
	c.JSONP(200, resp)
}

func UserInfo(c *gin.Context) {
	resp, err := logic.NewHandler().UserInfo()
	if err != nil {
		c.AbortWithError(200, err)
	}
	c.JSONP(200, resp)
}
