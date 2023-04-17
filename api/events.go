package api

import (
	"firebase-golang/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EventsPage(c *gin.Context) {
	events, err := database.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "events.html", gin.H{
		"Events": events,
	})
}
