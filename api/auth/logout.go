package auth

import (
	"firebase-golang/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	api.SetUserState(false)

	c.Redirect(http.StatusFound, "/")
}
