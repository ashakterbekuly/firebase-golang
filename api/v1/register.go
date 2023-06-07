package api

import (
	"context"
	"firebase-golang/api/v0"
	"firebase-golang/database"
	"firebase-golang/database/architects"
	"firebase-golang/database/clients"
	"firebase-golang/database/roles"
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

	form, _ := c.MultipartForm()
	files := form.File["photo"]
	for _, file := range files {
		architect.PhotoUrl, err = database.CreateArchitectPhoto(file)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if architect.Email == "" && architect.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid := architects.CreateArchitect(models.DaoFromArchitect(architect))
	user := (&auth.UserToCreate{}).Email(architect.Email).Password(architect.Password)

	dbUser, err := firebaseAuth.CreateUser(context.TODO(), user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = roles.SetRoleByEmail(dbUser.Email, "architects", uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiV0.SetUserState(true)

	c.Redirect(http.StatusFound, "/uid="+uid)
}

func RegisterClientGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register_client.html", gin.H{})
}

func RegisterClient(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	client := models.Client{}

	err := c.ShouldBind(&client)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["photo"]
	for _, file := range files {
		client.PhotoUrl, err = database.CreateClientPhoto(file)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if client.Email == "" && client.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid := clients.CreateClient(models.DaoFromClient(client))
	user := (&auth.UserToCreate{}).Email(client.Email).Password(client.Password)

	_, err = firebaseAuth.CreateUser(context.TODO(), user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = roles.SetRoleByEmail(client.Email, "clients", uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiV0.SetUserState(true)

	c.Redirect(http.StatusFound, "/?uid="+uid)
}
