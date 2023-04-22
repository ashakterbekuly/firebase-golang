package api

import (
	"firebase-golang/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Profile(c *gin.Context) {
	isAuthored := GetUserState()
	session := sessions.Default(c)
	data := session.Get("email")
	var email string
	if data != nil {
		email = data.(string)
		if email == "" {
			isAuthored = false
		} else {
			log.Println(email)
		}
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Username":           database.GetUsername(email),
	})
}
