package api

import (
	"firebase-golang/database/projects"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Projects(c *gin.Context) {
	projectsList, err := projects.GetProjectsList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projectsList,
	})
}
