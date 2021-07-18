package log

import (
	"log"
	"os"
)

// 日志对像，用于输出日志
var logger *log.Logger

func init() {
	// 打开(新建)日志文件
	logfile, err := os.OpenFile("personal-site.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 初始化日志文件格式
	logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
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
