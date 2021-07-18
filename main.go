package main

import (
	"personal-site/controller"
	"personal-site/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载配置文件
	utils.LoadConfig("config.json")
}

func main() {
	r := gin.Default()

	// API路由组
	api := r.Group("/api")
	{
		// 返回文章接口
		api.GET("/article", controller.Article)

		// 用户操作路由组
		user := api.Group("/user")
		{
			// 登录接口
			user.POST("/login", controller.Login)

			// 创建文章接口
			user.POST("/article", utils.AuthMiddleware(), controller.Create)
		}
	}

	r.Run(utils.Config.Port)
}
