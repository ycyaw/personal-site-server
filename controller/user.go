package controller

import (
	"fmt"
	"net/http"
	"personal-site/model"

	"github.com/gin-gonic/gin"
)

// 账号，密码结构体
type Info struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var info Info

	// 获取数据
	fmt.Println(c.BindJSON(&info))

	// 验证用户
	user, err := model.AuthUser(info.Account, info.Password)

	// 返回错误信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	// 返回toekn
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": user.Token,
		})
	}
}
