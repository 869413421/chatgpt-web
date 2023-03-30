package bootstrap

import (
	"github.com/869413421/chatgpt-web/config"
	"github.com/869413421/chatgpt-web/pkg/logger"
	"github.com/gin-gonic/gin"
	"mime"
	"strconv"
	"strings"
)

func StartWebServer() {
	// 注册启动所需各类参数
	SetUpRoute()
	SetupDB()
	initTemplateDir()
	initStaticServer()

	// 启动服务
	port := config.LoadConfig().Port
	portString := strconv.Itoa(port)
	// 自定义监听地址
	listen := config.LoadConfig().Listen
	err := router.Run(listen + ":" + portString)
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
	router.GET("/assets/:filename", func(c *gin.Context) {
		fileName := c.Param("filename")
		nameSlice := strings.Split(fileName, ".")
		ext := nameSlice[len(nameSlice)-1]
		if ext == "js" {
			c.Header("Content-Type", "application/javascript")
		} else {
			c.Header("Content-Type", mime.TypeByExtension(ext))
		}
		c.File("static/assets/" + c.Param("filename"))
	})

	//router.StaticFS("/assets", http.Dir("static/assets"))
	router.StaticFile("logo192.png", "static/logo192.png")
	router.StaticFile("logo512.png", "static/logo512.png")
	router.StaticFile("favicon.ico", "static/favicon.ico")
	router.StaticFile("manifest.json", "static/manifest.json")
}
