package main

import (
	"firebase-golang/api"
	"firebase-golang/config"
	"firebase-golang/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// initialize new gin engine (for server)
	r := gin.Default()

	// load static files
	r.LoadHTMLGlob("./static/*.html")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")

	// create/configure database instance
	db := database.CreateDatabase()
	auth := config.SetupFirebaseAuth()

	//set db and auth instance
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", auth)
	})

	//registration routes
	r.GET("/register/role", api.RegisterRoleGet)
	r.GET("/register/arch", api.RegisterArchitectGet)
	r.POST("/register/arch", api.RegisterArchitect)
	r.GET("/register/client", api.RegisterClientGet)
	r.POST("/register/client", api.RegisterClient)

	//login routes
	r.GET("/login", api.LoginGet)
	r.POST("/login", api.Login)

	// start the server
	err := r.Run(":5000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
