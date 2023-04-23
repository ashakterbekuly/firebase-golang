package api

import (
	"firebase-golang/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArchitecturePage(c *gin.Context) {
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

	var id string
	var username string
	if role == "client" {
		id = database.GetID(email)
		username = database.GetUsername(email)
	} else {
		id = database.GetArchitectDocumentIDByEmail(email)
		username = database.GetArchitectUsername(email)
	}

	c.HTML(http.StatusOK, "architecture.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"ID":                 id,
		"Username":           username,
	})
}
