package routes

import (
	. "github.com/869413421/chatgpt-web/app/http/controllers"
	"github.com/869413421/chatgpt-web/app/middlewares"
	"github.com/gin-gonic/gin"
)

var chatController = NewChatController()
var authController = NewAuthController()

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {

	router.Use(middlewares.Cors())
	router.GET("", chatController.Index)
	router.POST("user/auth", authController.Auth)
	chat := router.Group("/chat").Use(middlewares.Jwt())
	{
		chat.POST("/completion", chatController.Completion)
	}
	auth := router.Group("/auth").Use(middlewares.Jwt())
	{
		auth.POST("/info", authController.Info)
	}
}
