package bootstarp

import (
	"github.com/869413421/chatgpt/routes"
	"github.com/gin-gonic/gin"
	"sync"
)

var router *gin.Engine
var once sync.Once

func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		routes.RegisterWebRoutes(router)
	})
}
