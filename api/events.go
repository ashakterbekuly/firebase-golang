package api

import (
	"firebase-golang/database"
	"firebase-golang/firebase_auth"
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/sessions"
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

func SendMail(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	code := c.Query("code")

	session := sessions.Default(c)
	data := session.Get("email")
	var email string
	if data != nil {
		email = data.(string)
		if email != "" {
			log.Println(email)
		}
	}

	tokenRaw := session.Get("token")
	var token string
	if data != nil {
		token = tokenRaw.(string)
	}

	events, err := database.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, event := range events {
		if event.Code == code {
			firebase_auth.SendMessage(firebaseAuth, email, event.Title, token)
		}
	}

}
