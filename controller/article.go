package controller

import (
	"net/http"
	"personal-site/log"
	"personal-site/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Latest(c *gin.Context) {

	articles, err := model.LatestArticle()

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

func ArticleOfId(c *gin.Context) {
	ids := c.Query("id")


	id, err := strconv.ParseInt(ids, 10, 32)
	if err != nil {
		log.Info(err.Error())
	}

	article, err := model.QueryRowArticle(id)

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

func ArticleByCategory(c *gin.Context) {
	category := c.Query("category")

	articles, err := model.QueryByArticleCategory(category)

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

func SearchArticle(c *gin.Context) {
	title := c.Query("title")

	articles, err := model.QueryTitle(title)

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

func Create(c *gin.Context) {
	article := model.Article{}
	c.BindJSON(&article)

	err := model.InsertArticle(article.Title, article.Category, article.Content)

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