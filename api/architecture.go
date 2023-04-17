package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArchitecturePage(c *gin.Context) {
	c.HTML(http.StatusOK, "architecture.html", gin.H{})
}
