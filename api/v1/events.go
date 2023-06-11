package api

import (
	"firebase-golang/database/events"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Events(c *gin.Context) {
	eventsList, err := events.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": eventsList,
	})
}
