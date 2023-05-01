package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

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
