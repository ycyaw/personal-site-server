package main

import (
	"personal-site/controller"
	"personal-site/middleware"
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
			user.POST("/login", controller.UserLogin)

			// 注册接口
			user.POST("/register", controller.UserRegister)

			// 获取个人信息
			user.POST("/info", middleware.AuthMiddleware, controller.UserInfo)

			// 创建文章接口
			user.POST("/article", middleware.AuthMiddleware, controller.ArticleCreate)
		}
	}

	r.Run(utils.Config.Port)
}
