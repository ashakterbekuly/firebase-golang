package api

import (
	"firebase-golang/database/architects"
	"firebase-golang/database/clients"
	"firebase-golang/database/roles"
	"firebase-golang/database/templates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Profile(c *gin.Context) {
	uid := c.MustGet("ID").(string)
	role := roles.GetRoleByUID(uid)

	if role == roles.Client {
		client, err := clients.GetClientByUID(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"role":     roles.Client,
			"name":     client.Name,
			"bio":      client.Bio,
			"photoUrl": client.PhotoUrl,
		})
	} else if role == roles.Architect {
		architect, err := architects.GetArchitectByUID(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"role":           roles.Architect,
			"name":           architect.Name,
			"bio":            architect.Bio,
			"photoUrl":       architect.PhotoUrl,
			"specialization": architect.Specialization,
			"portfolio":      architect.Portfolio,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown role"})
		return
	}
}

func Architect(c *gin.Context) {
	uid := c.Query("ID")

	architect, err := architects.GetArchitectByUID(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmps, err := templates.GetTemplatesByAuthorID(uid, architect.Name)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"role":           roles.Architect,
		"name":           architect.Name,
		"bio":            architect.Bio,
		"photoUrl":       architect.PhotoUrl,
		"specialization": architect.Specialization,
		"portfolio":      architect.Portfolio,
		"templates":      tmps,
	})
}
