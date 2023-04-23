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
	id := c.Query("id")

	var email string
	if data != nil {
		email = data.(string)
		if email == "" {
			isAuthored = false
		} else {
			log.Println(email)
		}
	}

	client, err := database.GetClientByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Name":               client.Name,
		"Bio":                client.Bio,
		"PhotoUrl":           client.PhotoUrl,
		"Username":           database.GetUsername(email),
	})
}
