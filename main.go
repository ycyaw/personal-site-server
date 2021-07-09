package main

import (
	"fmt"
	"net/http"
	"personal-site/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载配置文件
	utils.LoadConfig("config.json")
}

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Article struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

func main() {
	r := gin.Default()

	// 登录接口
	r.POST("/api/login", func(c *gin.Context) {
		var user User

		fmt.Println(c.BindJSON(&user))

		if user.Account == "123456" && user.Password == "qwerty" {
			c.JSON(http.StatusOK, gin.H{
				"token": "etq3rgbartw45y34at",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
		}
	})

	// 获取最新文章
	r.GET("/api/latest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"title": "this is a test title",
				"text":  "hwehgAWGah\nwgwagWG\ngWG\n",
			},
		})
	})

	r.POST("/api/create", func(c *gin.Context) {
		article := Article{}
		c.BindJSON(&article)

		fmt.Println(article.Title)
		fmt.Println(article.Category)
		fmt.Println(article.Content)
	})

	r.Run(utils.Config.Port)
}
