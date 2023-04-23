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

func EditProfile(c *gin.Context) {
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

	form, _ := c.MultipartForm()
	files := form.File["photo"]
	var newPhotoUrl string
	if len(files) == 0 {
		newPhotoUrl = database.GetPhotoUrl(currentEmail)
	} else {
		err := database.DeleteClientImage(database.GetPhotoUrl(currentEmail))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка удаления фото"})
			return
		}
		for _, file := range files {
			newPhotoUrl, _ = database.CreateClientPhoto(file)
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
	err = database.UpdateClient(currentEmail, models.Client{
		Email:    newEmail,
		Name:     newName,
		Bio:      newBio,
		PhotoUrl: newPhotoUrl,
	})

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
	c.Redirect(http.StatusFound, fmt.Sprintf("/profile?id=%s", database.GetID(newEmail)))
}
