package middleware

import (
	"net/http"
	"personal-site/model"

	"github.com/gin-gonic/gin"
)

// 验证用户token中间件
func AuthMiddleware(c *gin.Context) {
	// 获取http头部信息
	authorization := c.Request.Header.Get("Authorization")

	// 验证获取的token
	responseUser, err := model.QueryUserOfToken(authorization)

	// 验证失败时终止后续执行操作
	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		// 终止后续执行
		c.Abort()
	} else {
		// 设置需要后续使用的数据
		c.Set("email", responseUser.Email)
		c.Set("name", responseUser.Name)
		c.Set("token", responseUser.Token)
	}
}
