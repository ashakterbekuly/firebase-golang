package auth

import (
	"firebase-golang/api"
	"firebase-golang/database"
	"firebase-golang/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EditArchitectProfile(c *gin.Context) {
	session := sessions.Default(c)
	currentEmail := session.Get("email").(string)

	// Получаем данные из формы
	currentPassword := c.PostForm("current-password")
	if currentPassword == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
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
		newPhotoUrl = database.GetArchitectPhotoUrl(currentEmail)
	} else {
		err := database.DeleteArchitectImage(database.GetArchitectPhotoUrl(currentEmail))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка удаления фото"})
			return
		}
		for _, file := range files {
			newPhotoUrl, _ = database.CreateArchitectPhoto(file)
		}
	}

	// Проверяем аутентификацию пользователя
	token, err := sendAuthRequest(currentEmail, currentPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	// Обновляем учетную запись пользователя
	err = sendUpdateUserRequest(token, newEmail, newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Обновляем его в базе
	err = database.UpdateArchitect(currentEmail, models.Architect{
		Email:          newEmail,
		Name:           newName,
		Bio:            newBio,
		PhotoUrl:       newPhotoUrl,
		Specialization: newSpecialization,
		Portfolio:      newPortfolio,
	})

	err = database.UpdateRoleByEmail(currentEmail, newEmail, "architect")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SetUserState(true)
	session.Set("token", token)
	session.Set("email", newEmail)
	err = session.Save()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем сообщение об успешном обновлении учетной записи
	c.Redirect(http.StatusFound, fmt.Sprintf("/profile?id=%s", database.GetArchitectDocumentIDByEmail(newEmail)))
}
