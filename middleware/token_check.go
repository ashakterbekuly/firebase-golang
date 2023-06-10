package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const SignKey = "archneo"

func TokenCheck(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// Проверка формата токена
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Извлечение токена из "Bearer <token>"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка используемого алгоритма подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Возвращаем ключ подписи
		return []byte(SignKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Проверка валидности токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Выполнение дополнительных действий при успешной валидации токена
		// Например, сохранение идентификатора пользователя из токена в контексте запроса
		c.Set("ID", claims["username"])
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
