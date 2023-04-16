package api

import (
	"context"
	"firebase-golang/database"
	"firebase-golang/models"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterRoleGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register_role.html", gin.H{})
}

func RegisterArchitectGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register_arch.html", gin.H{})
}

func RegisterArchitect(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	architect := models.Architect{}

	err := c.ShouldBind(&architect)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if architect.Email == "" && architect.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = database.CreateArchitect(architect)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := (&auth.UserToCreate{}).Email(architect.Email).Password(architect.Password)

	_, err = firebaseAuth.CreateUser(context.TODO(), user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
