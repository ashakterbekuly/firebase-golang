package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
//func AuthMiddleware(c *gin.Context) {
//	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
//	authorizationToken := c.GetHeader("Authorization")
//	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
//	if idToken == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
//		c.Abort()
//		return
//	}
//	//verify token
//	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
//		c.Abort()
//		return
//	}
//	c.Set("UUID", token.UID)
//	c.Next()
//}

//CheckAuth : to verify all authorized operations
func CheckAuth(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	idToken, isExist := c.Get("token")
	if !isExist {
		c.Set("UUID", "")
		c.Next()
		return
	}

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken.(string))
	if err != nil {
		c.Set("UUID", "")
		c.Next()
		return
	}

	c.Set("UUID", token.UID)
	c.Next()
}
