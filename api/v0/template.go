package apiV0

import (
	"firebase-golang/database/roles"
	"firebase-golang/database/templates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Template(c *gin.Context) {
	templatesList, err := templates.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authored := GetUserState()

	res := map[string]interface{}{
		"IsNonAuthenticated": !authored,
	}

	if authored {
		uid := c.Query("uid")
		res["ID"] = uid
		res["Username"] = roles.GetUsernameByUID(uid)
	}

	switch c.Query("spec") {
	case "House Draft":
		res["Templates"] = templates.GetHouseDrafts(templatesList)
		c.HTML(http.StatusOK, "cottage_drafts.html", res)
	case "Interior Design":
		res["Templates"] = templates.GetInteriorDesigns(templatesList)
		c.HTML(http.StatusOK, "interior_designs.html", res)
	case "Urban Project":
		res["Templates"] = templates.GetUrbanProjects(templatesList)
		c.HTML(http.StatusOK, "urban_projects.html", res)
	case "Reconstruction":
		res["Templates"] = templates.GetReconstructionProjects(templatesList)
		c.HTML(http.StatusOK, "restoration_projects.html", res)
	default:
		c.HTML(http.StatusOK, "index.html", res)
	}
}
