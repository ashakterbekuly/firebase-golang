package api

import (
	"firebase-golang/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ProjectsGet(c *gin.Context) {
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

	projects, err := database.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "projects.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Projects":           projects,
		"Username":           database.GetUsername(email),
	})
}
