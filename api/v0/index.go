package apiV0

import (
	"firebase-golang/database/events"
	"firebase-golang/database/projects"
	"firebase-golang/database/roles"
	"firebase-golang/database/templates"
	"firebase-golang/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MainPage(c *gin.Context) {
	eventsList, err := events.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	projectsList, err := projects.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var showProjects []models.Project
	if len(projectsList) > 2 {
		for i := 0; i < 2; i++ {
			showProjects = append(showProjects, projectsList[i])
		}
	} else {
		showProjects = projectsList
	}

	authored := GetUserState()

	res := map[string]interface{}{
		"IsNonAuthenticated": !authored,
		"Events":             sortEventsByDateFrom(eventsList),
		"Projects":           showProjects,
		"Templates":          templates.SortTemplates(templatesList),
	}

	if !authored {
		c.HTML(http.StatusOK, "index.html", res)
		return
	}

	uid := c.Query("uid")
	res["ID"] = uid
	res["Username"] = roles.GetUsernameByUID(uid)

	c.HTML(http.StatusOK, "index.html", res)
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
