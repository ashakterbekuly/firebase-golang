package api

import (
	"firebase-golang/database"
	"firebase-golang/database/architects"
	"firebase-golang/database/roles"
	"firebase-golang/firebase_auth"
	"firebase-golang/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EditArchitectProfile(c *gin.Context) {
	uid := c.MustGet("ID").(string)
	currentEmail := roles.GetEmailByUID(uid)

	// Получаем данные из формы
	currentPassword := c.PostForm("current-password")
	if currentPassword == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password field is empty"})
		return
	}
	newEmail := c.PostForm("new-email")
	if newEmail == "" {
		newEmail = currentEmail
	}
	newPassword := c.PostForm("new-password")
	if newPassword == "" {
		newPassword = currentPassword
	}
	newName := c.PostForm("name")
	newBio := c.PostForm("bio")
	newSpecialization := c.PostForm("specialization")
	newPortfolio := c.PostForm("portfolio")

	form, _ := c.MultipartForm()
	files := form.File["photo"]
	var newPhotoUrl string
	if len(files) == 0 {
		newPhotoUrl = architects.GetArchitectPhotoUrl(uid)
	} else {
		err := database.DeleteArchitectImage(architects.GetArchitectPhotoUrl(uid))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		for _, file := range files {
			newPhotoUrl, _ = database.CreateArchitectPhoto(file)
		}
	}

	// Проверяем аутентификацию пользователя
	token, err := firebase_auth.GetUserToken(currentEmail, currentPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Обновляем учетную запись пользователя
	err = firebase_auth.UpdateUserRequest(token, newEmail, newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Обновляем его в базе
	err = architects.UpdateArchitect(uid, models.Architect{
		Email:          newEmail,
		Name:           newName,
		Bio:            newBio,
		PhotoUrl:       newPhotoUrl,
		Specialization: newSpecialization,
		Portfolio:      newPortfolio,
	})

	err = roles.UpdateRoleByEmail(currentEmail, newEmail, roles.Architect)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
