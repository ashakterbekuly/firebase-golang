package api

import (
	"firebase-golang/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	login := LoginForm{}

	err := c.Bind(&login)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if login.Email == "" && login.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = config.SendAuthRequest(login.Email, login.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect the user to the dashboard or home page
	//c.Redirect(http.StatusFound, "/artists")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
