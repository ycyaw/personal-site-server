package controller

import (
	"net/http"
	"personal-site/log"
	"personal-site/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 返回文章接口
// 通过category参数查询文章分类
// 通过search参数查询文章
// 通过id参数查询指定文章
// 无参数返回最新文章
func Article(c *gin.Context) {
	// 获id取参数
	if ids := c.Query("id"); ids != "" {
		articleOfId(c)

	} else if category := c.Query("category"); category != "" {
		// 获取类别参数
		articleByCategory(c)

	} else if search := c.Query("search"); search != "" {
		// 获取搜索参数
		articleByKey(c)

	} else {
		// 无参数时返回最新文章
		articleByOrder(c)
	}
}

// 响应最近的文章数据
func articleByOrder(c *gin.Context) {
	// 从数据库获取数据
	articles, err := model.QueryArticleOfLatest()

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
func articleOfId(c *gin.Context) {
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
func articleByCategory(c *gin.Context) {
	// 获取参数
	category := c.Query("category")

	// 从数据库获取数据
	articles, err := model.QueryArticleByCategory(category)

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
func articleByKey(c *gin.Context) {
	// 获取参数
	search := c.Query("search")

	// 从数据库获取数据
	articles, err := model.QueryArticleOfTitle(search)

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

	// 通过token获取用户名
	name := model.QueryUserOfToken(authorization)

	// 插入的数据数据
	article := model.Article{}
	// 获取josn数据
	c.BindJSON(&article)

	// 将数据插入表中
	err := model.InsertArticle(article.Title, name, article.Category, article.Content)

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
