package api

import (
	"firebase-golang/database/projects"
	"firebase-golang/database/roles"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ProjectsGet(c *gin.Context) {
	authored := GetUserState()

	projectsList, err := projects.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := map[string]interface{}{
		"IsNonAuthenticated": !authored,
		"Projects":           projectsList,
	}

	if !authored {
		c.HTML(http.StatusOK, "projects.html", res)
		return
	}

	uid := c.Query("uid")

	res["ID"] = uid
	res["Username"] = roles.GetUsernameByUID(uid)

	c.HTML(http.StatusOK, "projects.html", res)
}
