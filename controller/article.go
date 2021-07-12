package controller

import (
	"net/http"
	"personal-site/log"
	"personal-site/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 响应最近的文章数据
func Latest(c *gin.Context) {
	// 从数据库获取数据
	articles, err := model.LatestArticle()

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": articles,
		})
	}
}

// 通过文章Id查询文章
func ArticleOfId(c *gin.Context) {
	// 获取参数
	ids := c.Query("id")

	// 类型转换
	id, err := strconv.ParseInt(ids, 10, 32)
	if err != nil {
		log.Info(err.Error())
	}

	// 从数据库获取数据
	article, err := model.QueryRowArticle(id)

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": article,
		})
	}
}

// 通过文章类别查询文章
func ArticleByCategory(c *gin.Context) {
	// 获取参数
	category := c.Query("category")

	// 从数据库获取数据
	articles, err := model.QueryByArticleCategory(category)

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": articles,
		})
	}
}

// 通过文章关键字搜索文章
func SearchArticle(c *gin.Context) {
	// 获取参数
	title := c.Query("title")

	// 从数据库获取数据
	articles, err := model.QueryTitle(title)

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": articles,
		})
	}
}

// 将用户的文章插入数据库中
func Create(c *gin.Context) {
	// 获取http头部信息
	authorization := c.Request.Header.Get("Authorization")

	// 验证获取的token
	name, err := model.AuthToken(authorization)

	// 返回数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	} else {
		article := model.Article{}
		// 获取josn数据
		c.BindJSON(&article)

		// 将数据插入表中
		err = model.InsertArticle(article.Title, name, article.Category, article.Content)

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
}
