package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainPage(c *gin.Context) {
	isAuthored := GetUserState()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsNonAuthenticated": !isAuthored,
	})
}
