package controller

import (
	"net/http"
	"personal-site/model"

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