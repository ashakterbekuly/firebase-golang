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
	uid := c.Param("uid")

	projectsList, err := projects.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "projects.html", gin.H{
		"IsNonAuthenticated": !authored,
		"Projects":           projectsList,
		"ID":                 uid,
		"Username":           roles.GetUsernameByUID(uid),
	})
}
