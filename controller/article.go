package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

func Latest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"title": "this is a test title",
			"text":  "hwehgAWGah\nwgwagWG\ngWG\n",
		},
	})
}

func Create(c *gin.Context) {
	article := Article{}
	c.BindJSON(&article)

	fmt.Println(article.Title)
	fmt.Println(article.Category)
	fmt.Println(article.Content)
}