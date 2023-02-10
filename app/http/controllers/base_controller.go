package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (*BaseController) ResponseJson(ctx *gin.Context, code int, errorMsg string, data interface{}) {

	ctx.JSON(code, gin.H{
		"code":     code,
		"errorMsg": errorMsg,
		"data":     data,
	})
	ctx.Abort()
}
