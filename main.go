package main

import (
	"personal-site/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载配置文件
	utils.LoadConfig("config.json")
}

func main() {
	r := gin.Default()
	r.Run(utils.Config.Port)
}