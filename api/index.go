package api

import (
	"firebase-golang/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MainPage(c *gin.Context) {
	isAuthored := GetUserState()

	events, err := database.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Events":             events,
	})
}
