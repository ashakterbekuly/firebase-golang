package api

import (
	apiV0 "firebase-golang/api/v0"
	"firebase-golang/database/roles"
	"firebase-golang/firebase_auth"
	"firebase-golang/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type LoginForm struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	login := LoginForm{}

	err := c.Bind(&login)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if login.Email == "" || login.Password == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is empty"})
		return
	}

	_, err = firebase_auth.GetUserToken(login.Email, login.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	uid := roles.GetUIDByEmail(login.Email)

	tokenString, err := GetToken(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	apiV0.SetUserState(true)

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": strings.ToUpper(roles.GetUsernameByUID(uid)),
	})
}

func GetToken(uid string) (string, error) {
	// Генерация токена
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Время жизни токена

	// Подпись токена
	tokenString, err := token.SignedString([]byte(middleware.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
