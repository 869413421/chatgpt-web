package bootstarp

import "net/http"

func StartWebServer() {
	// 注册启动所需各类参数
	SetUpRoute()
	initTemplateDir()
	initStaticServer()

	// 启动服务
	router.Run(":8080")
}

// initTemplate 初始化HTML模板加载路径
func initTemplateDir() {
	router.LoadHTMLGlob("resources/view/*")
}

// initStaticServer 初始化静态文件处理
func initStaticServer() {
	router.StaticFS("/static", http.Dir("public"))
}
