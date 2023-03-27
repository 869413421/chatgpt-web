package main

import (
	"github.com/869413421/chatgpt-web/bootstarp"
	"github.com/869413421/chatgpt-web/config"
	"github.com/alecthomas/kong"
)

func main() {
	kong.Parse(&config.CLI)
	bootstrap.StartWebServer()
}
