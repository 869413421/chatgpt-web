package routes

import (
	. "github.com/869413421/chatgpt-web/app/http/controllers"
	"github.com/869413421/chatgpt-web/app/middlewares"
	"github.com/869413421/chatgpt-web/config"
	"github.com/gin-gonic/gin"
)

var chatController = NewChatController()

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {

	router.Use(middlewares.Cors())

	cnf := config.LoadConfig()
	if len(cnf.AuthUser) > 0 {
		gin.BasicAuth(gin.Accounts{
			cnf.AuthUser: cnf.AuthPassword,
		})
	}
	router.GET("/", chatController.Index)
	router.POST("/completion", chatController.Completion)
}
