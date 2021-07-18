package middleware

import (
	"net/http"
	"personal-site/model"

	"github.com/gin-gonic/gin"
)

// 验证用户token中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取http头部信息
		authorization := c.Request.Header.Get("Authorization")

		// 验证获取的token
		err := model.QueryUserToken(authorization)

		if err != nil {
			// 返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			// 终止后续执行
			c.Abort()
		} else {
			// 验证成功继续执行后续操作
			c.Next()
		}
	}
}
