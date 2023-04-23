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

	events, err := database.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var id string
	var username string
	if role == "client" {
		id = database.GetID(email)
		username = database.GetUsername(email)
	} else {
		id = database.GetArchitectDocumentIDByEmail(email)
		username = database.GetArchitectUsername(email)
	}

	c.HTML(http.StatusOK, "events.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Events":             events,
		"ID":                 id,
		"Username":           username,
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
