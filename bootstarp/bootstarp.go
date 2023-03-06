package bootstarp

import (
	"github.com/869413421/chatgpt-web/config"
	"github.com/869413421/chatgpt-web/pkg/logger"
	"net/http"
	"strconv"
)

func StartWebServer() {
	// 注册启动所需各类参数
	SetUpRoute()
	initTemplateDir()
	initStaticServer()

	// 启动服务
	port := config.LoadConfig().Port
	portString := strconv.Itoa(port)
	err := router.Run(":" + portString)
	if err != nil {
		logger.Danger("run webserver error %s", err)
		return
	}
}

// initTemplate 初始化HTML模板加载路径
func initTemplateDir() {
	router.LoadHTMLGlob("resources/view/*")
}

// initStaticServer 初始化静态文件处理
func initStaticServer() {
	router.StaticFS("/assets", http.Dir("static/assets"))
	router.StaticFile("logo192.png","static/logo192.png")
	router.StaticFile("logo512.png","static/logo512.png")
	router.StaticFile("favicon.ico","static/favicon.ico")
	router.StaticFile("manifest.json","static/manifest.json")
}
