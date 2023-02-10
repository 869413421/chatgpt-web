package logger

import (
	"log"
	"os"
	"sync"
)

var Logger *log.Logger
var once sync.Once

func init() {
	once.Do(func() {
		Logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// Info 详情
func Info(args ...interface{}) {
	Logger.SetPrefix("[INFO]")
	Logger.Println(args...)
}

// Danger 错误 为什么不命名为 error？避免和 error 类型重名
func Danger(args ...interface{}) {
	Logger.SetPrefix("[ERROR]")
	Logger.Fatal(args...)
}

// Warning 警告
func Warning(args ...interface{}) {
	Logger.SetPrefix("[WARNING]")
	Logger.Println(args...)
}

// DeBug debug
func DeBug(args ...interface{}) {
	Logger.SetPrefix("[DeBug]")
	Logger.Println(args...)
}
