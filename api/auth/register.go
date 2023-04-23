package auth

import (
	"context"
	"firebase-golang/api"
	"firebase-golang/database"
	"firebase-golang/models"
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/sessions"
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

	err = database.CreateArchitect(models.DaoFromArchitect(architect))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := (&auth.UserToCreate{}).Email(architect.Email).Password(architect.Password)

	dbUser, err := firebaseAuth.CreateUser(context.TODO(), user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api.SetUserState(true)
	session := sessions.Default(c)
	session.Set("token", "")
	session.Set("email", dbUser.Email)
	err = session.Save()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/")
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

	err = database.CreateClient(models.DaoFromClient(client))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := (&auth.UserToCreate{}).Email(client.Email).Password(client.Password)

	_, err = firebaseAuth.CreateUser(context.TODO(), user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api.SetUserState(true)
	session := sessions.Default(c)
	session.Set("email", client.Email)
	err = session.Save()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
