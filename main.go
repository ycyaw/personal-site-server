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

		// 登录接口
		api.POST("/login", controller.UserLogin)

		// 注册接口
		api.POST("/register", controller.UserRegister)

		// 用户操作路由组
		user := api.Group("/user")
		// 使用token验证中间件
		user.Use(middleware.AuthMiddleware)
		{
			// 获取指定的用户信息
			user.GET("/user", controller.UserGet)

			// 更新指定的用户信息
			user.PATCH("/user", controller.UserPatch)

			// 获取指定的文章
			user.GET("/article/:id", controller.ArticleGet)

			// 创建文章接口
			user.POST("/article", controller.ArticleCreate)

			// 修改指定的文章
			user.PATCH("/article/:id", controller.ArticlePatch)

			// 删除指定的文章
			user.DELETE("/article/:id", controller.ArticleDelete)
		}
	}

	r.Run(utils.Config.Port)
}
