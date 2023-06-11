package api

import (
	"firebase-golang/database/templates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Templates(c *gin.Context) {
	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"houseDrafts":     templates.GetHouseDrafts(templatesList),
		"interiorDesigns": templates.GetInteriorDesigns(templatesList),
		"urbanProjects":   templates.GetUrbanProjects(templatesList),
		"reconstructions": templates.GetReconstructionProjects(templatesList),
	})
}

func HouseDrafts(c *gin.Context) {
	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"templates": templates.GetHouseDrafts(templatesList),
	})
}

func InteriorDesigns(c *gin.Context) {
	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"templates": templates.GetInteriorDesigns(templatesList),
	})
}

func UrbanProjects(c *gin.Context) {
	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"templates": templates.GetUrbanProjects(templatesList),
	})
}

func Reconstructions(c *gin.Context) {
	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"templates": templates.GetReconstructionProjects(templatesList),
	})
}
