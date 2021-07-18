package controller

import (
	"fmt"
	"net/http"
	"personal-site/model"

	"github.com/gin-gonic/gin"
)

// 用户登录
func UserLogin(c *gin.Context) {
	var user model.User

	// 获取数据
	fmt.Println(c.BindJSON(&user))

	// 验证用户
	user, err := model.QueryUserOfEmailAndPasswd(user.Email, user.Password)

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

// 用户注册
func UserRegister(c *gin.Context) {
	var user model.User

	// 获取json数据
	c.BindJSON(&user)

	// 插入数据到数据库
	err := model.InsertUser(user.Email, user.Name, user.Password)

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}

// 返回用户信息
func UserGet(c *gin.Context) {
	// 获取中间件设置的信息
	id, _ := c.Get("id")
	email, _ := c.Get("email")
	name, _ := c.Get("name")
	token, _ := c.Get("token")

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    id,
			"email": email,
			"name":  name,
			"token": token,
		},
	})
}

// 更新用户信息
func UserPatch(c *gin.Context) {
	// 获取用户id
	id, _ := c.Get("id")
	idStr := id.(string)

	// 获取json数据
	user := model.User{}
	c.BindJSON(&user)

	err := model.UpdateUser(idStr, user.Email, user.Name)

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}
