package api

import (
	"firebase-golang/database/roles"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArchitecturePage(c *gin.Context) {
	authored := GetUserState()
	uid := c.Param("uid")

	c.HTML(http.StatusOK, "architecture.html", gin.H{
		"IsNonAuthenticated": !authored,
		"ID":                 uid,
		"Username":           roles.GetUsernameByUID(uid),
	})
}
