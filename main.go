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

	// 登录接口
	r.POST("/api/login", controller.Login)

	// 获取最新文章
	r.GET("/api/latest", controller.Latest)

	// 依据id获取指定文章
	r.GET("/api/article", controller.ArticleOfId)

	// 依据分类获取文章
	r.GET("/api/articleyycategory", controller.ArticleByCategory)

	// 搜索文章
	r.GET("/api/search", controller.SearchArticle)

	// 新建文章
	r.POST("/api/create", controller.Create)

	r.Run(utils.Config.Port)
}
