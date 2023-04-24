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

	projects, err := database.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	templates, err := database.GetTemplatesList()
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
		"Events":             sortEventsByDateFrom(events),
		"Projects":           showProjects,
		"Templates":          database.SortTemplates(templates),
		"ID":                 id,
		"Username":           username,
	})
}

func sortEventsByDateFrom(events []models.Event) []models.Event {
	var threeEvents []models.Event

	for i := 0; i < len(events); i++ {
		for j := 0; j < len(events)-1; j++ {
			if events[j].DateFrom.After(events[j+1].DateFrom) {
				events[j], events[j+1] = events[j+1], events[j]
			}
		}
	}

	if len(events) > 3 {
		for i := 0; i < 3; i++ {
			threeEvents = append(threeEvents, events[i])
		}
		return threeEvents
	}

	return threeEvents
}
