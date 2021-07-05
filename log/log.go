package log

import (
	"log"
	"os"
)

// 日志对像，用于输出日志
var logger *log.Logger

func init() {
	// 初始化日志文件格式
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// [INFO]级别日志
func Info(args ...interface{}) {
	logger.SetPrefix("[INFO]    ")
	logger.Println(args...)
}

// [WARNING]级别日志
func Warning(args ...interface{}) {
	logger.SetPrefix("[WARNING] ")
	logger.Println(args...)
}

// [ERROR]级别日志
func Error(args ...interface{}) {
	logger.SetPrefix("[ERROR]   ")
	logger.Println(args...)
}
