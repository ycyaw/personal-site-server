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
	Account  string
	Password string
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

	r.Run(utils.Config.Port)
}
