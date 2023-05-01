package api

import (
	"firebase-golang/database/events"
	"firebase-golang/database/roles"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EventsPage(c *gin.Context) {
	authored := GetUserState()
	uid := c.Param("uid")

	eventsList, err := events.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "events.html", gin.H{
		"IsNonAuthenticated": !authored,
		"Events":             eventsList,
		"ID":                 uid,
		"Username":           roles.GetUsernameByUID(uid),
	})
}
