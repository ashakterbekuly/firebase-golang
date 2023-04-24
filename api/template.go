package api

import (
	"firebase-golang/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Template(c *gin.Context) {
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

	templates, err := database.GetTemplatesList()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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

	res := map[string]interface{}{
		"IsNonAuthenticated": !isAuthored,
		"ID":                 id,
		"Username":           username,
	}

	switch c.Query("spec") {
	case "House Draft":
		res["Templates"] = database.GetHouseDrafts(templates)
		c.HTML(http.StatusOK, "cottage_drafts.html", res)
	case "Interior Design":
		res["Templates"] = database.GetInteriorDesigns(templates)
		c.HTML(http.StatusOK, "interior_designs.html", res)
	case "Urban Project":
		res["Templates"] = database.GetUrbanProjects(templates)
		c.HTML(http.StatusOK, "urban_projects.html", res)
	case "Reconstruction":
		res["Templates"] = database.GetReconstructionProjects(templates)
		c.HTML(http.StatusOK, "restoration_projects.html", res)
	default:
		c.HTML(http.StatusOK, "index.html", res)
	}
}
