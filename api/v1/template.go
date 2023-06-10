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
		"HouseDrafts":     templates.GetHouseDrafts(templatesList),
		"InteriorDesigns": templates.GetInteriorDesigns(templatesList),
		"UrbanProjects":   templates.GetUrbanProjects(templatesList),
		"Reconstructions": templates.GetReconstructionProjects(templatesList),
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
		"Templates": templates.GetHouseDrafts(templatesList),
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
		"Templates": templates.GetInteriorDesigns(templatesList),
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
		"Templates": templates.GetUrbanProjects(templatesList),
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
		"Templates": templates.GetReconstructionProjects(templatesList),
	})
}
