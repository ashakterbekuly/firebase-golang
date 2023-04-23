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
	var email string
	rawEmail := session.Get("email")
	if rawEmail != nil {
		email = rawEmail.(string)
	}
	var role string
	rawRole := session.Get("role")
	if rawRole != nil {
		role = rawRole.(string)
	}

	if role == "client" {
		client, err := database.GetClientByEmail(email)
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
	} else {
		architect, err := database.GetArchitectByEmail(email)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "architecture.html", gin.H{
			"IsNonAuthenticated": !isAuthored,
			"Name":               architect.Name,
			"Bio":                architect.Bio,
			"PhotoUrl":           architect.PhotoUrl,
			"Specialization":     architect.Specification,
			"Portfolio":          architect.Portfolio,
			"Username":           database.GetArchitectUsername(email),
		})
	}
}
