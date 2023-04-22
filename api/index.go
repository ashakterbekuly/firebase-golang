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
	data := session.Get("email")
	var email string
	if data != nil {
		email = data.(string)
		if email == "" {
			isAuthored = false
		} else {
			log.Println(email)
		}
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

	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
		"Events":             showEvents,
		"Projects":           showProjects,
		"Username":           database.GetUsername(email),
	})
}
