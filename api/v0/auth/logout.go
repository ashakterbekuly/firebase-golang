package auth

import (
	apiV0 "firebase-golang/api/v0"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	apiV0.SetUserState(false)

	c.Redirect(http.StatusFound, "/")
}
