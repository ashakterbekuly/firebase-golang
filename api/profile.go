package api

import (
	"firebase-golang/database/architects"
	"firebase-golang/database/clients"
	"firebase-golang/database/roles"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Profile(c *gin.Context) {
	authored := GetUserState()
	uid := c.Param("uid")

	if roles.GetRoleByUID(uid) == "client" {
		client, err := clients.GetClientByID(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "profile.html", gin.H{
			"IsNonAuthenticated": !authored,
			"Name":               client.Name,
			"Bio":                client.Bio,
			"PhotoUrl":           client.PhotoUrl,
			"Username":           client.Name,
		})
	} else {
		architect, err := architects.GetArchitectByID(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "architecture.html", gin.H{
			"IsNonAuthenticated": !authored,
			"Name":               architect.Name,
			"Bio":                architect.Bio,
			"PhotoUrl":           architect.PhotoUrl,
			"Specialization":     architect.Specialization,
			"Portfolio":          architect.Portfolio,
			"Username":           architect.Name,
		})
	}
}
