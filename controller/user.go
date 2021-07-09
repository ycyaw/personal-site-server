package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user User

	fmt.Println(c.BindJSON(&user))

	if user.Account == "123456" && user.Password == "qwerty" {
		c.JSON(http.StatusOK, gin.H{
			"token": "etq3rgbartw45y34at",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	}
}
