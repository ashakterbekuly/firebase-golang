package api

import (
	"errors"
	apiV0 "firebase-golang/api/v0"
	"firebase-golang/database/roles"
	"firebase-golang/firebase_auth"
	"firebase-golang/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("email or password is empty")})
		return
	}

	_, err = firebase_auth.GetUserToken(login.Email, login.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Генерация токена
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = roles.GetUIDByEmail(login.Email)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Время жизни токена

	// Подпись токена
	tokenString, err := token.SignedString([]byte(middleware.SignKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	apiV0.SetUserState(true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
