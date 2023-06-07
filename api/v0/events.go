package apiV0

import (
	"firebase-golang/database/events"
	"firebase-golang/database/roles"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EventsPage(c *gin.Context) {
	authored := GetUserState()

	eventsList, err := events.GetEventsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := map[string]interface{}{
		"IsNonAuthenticated": !authored,
		"Events":             eventsList,
	}

	if !authored {
		c.HTML(http.StatusOK, "events.html", res)
		return
	}

	uid := c.Query("uid")

	res["ID"] = uid
	res["Username"] = roles.GetUsernameByUID(uid)

	c.HTML(http.StatusOK, "events.html", res)
}
