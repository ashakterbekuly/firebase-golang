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
			"Role":     roles.Client,
			"Name":     client.Name,
			"Bio":      client.Bio,
			"PhotoUrl": client.PhotoUrl,
		})
	} else if role == roles.Architect {
		architect, err := architects.GetArchitectByUID(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Role":           roles.Architect,
			"Name":           architect.Name,
			"Bio":            architect.Bio,
			"PhotoUrl":       architect.PhotoUrl,
			"Specialization": architect.Specialization,
			"Portfolio":      architect.Portfolio,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown role"})
		return
	}
}
