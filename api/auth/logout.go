package auth

import (
	"firebase-golang/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutGet(c *gin.Context) {
	api.SetUserState(false)
	session := sessions.Default(c)
	session.Delete("email")
	session.Delete("token")

	c.Redirect(http.StatusFound, "/")
}
