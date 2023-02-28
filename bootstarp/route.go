package bootstarp

import (
	"github.com/869413421/chatgpt-web/routes"
	"github.com/gin-gonic/gin"
	"sync"
)

var router *gin.Engine
var once sync.Once

func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		// 添加basic认证
		router.Use(gin.BasicAuth(gin.Accounts{
		"qaxnb": "qaxnb",
		}))
		routes.RegisterWebRoutes(router)
	})
}
