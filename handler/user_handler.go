package handler

import (
	"free-community/entity"
	"free-community/logic"
	"github.com/gin-gonic/gin"
	"log"
)

func UserRegister(c *gin.Context) {
	var req entity.UserRegisterReq
	if err := c.ShouldBind(&req); err != nil {
		log.Panicln(err)
	}
	resp, err := logic.NewHandler().UserRegister(req)
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
