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