package api

import (
	"firebase-golang/database"
	"firebase-golang/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MainPage(c *gin.Context) {
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

	var showEvents []models.Event
	if len(events) > 4 {
		for i := 0; i < 4; i++ {
			showEvents = append(showEvents, events[i])
		}
	} else {
		showEvents = events
	}

	projects, err := database.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var showProjects []models.Project
	if len(projects) > 2 {
		for i := 0; i < 2; i++ {
			showProjects = append(showProjects, projects[i])
		}
	} else {
		showProjects = projects
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

	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Events":             showEvents,
		"Projects":           showProjects,
		"ID":                 id,
		"Username":           username,
	})
}
