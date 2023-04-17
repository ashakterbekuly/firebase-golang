package main

import (
	"firebase-golang/api"
	"firebase-golang/api/auth"
	"firebase-golang/database"
	"firebase-golang/firebase_auth"
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
	firebaseAuth := firebase_auth.SetupFirebaseAuth()

	//set db and auth instance
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	//registration routes
	r.GET("/register/role", auth.RegisterRoleGet)
	r.GET("/register/arch", auth.RegisterArchitectGet)
	r.POST("/register/arch", auth.RegisterArchitect)
	r.GET("/register/client", auth.RegisterClientGet)
	r.POST("/register/client", auth.RegisterClient)

	//login routes
	r.GET("/login", auth.LoginGet)
	r.POST("/login", auth.Login)

	//main page
	r.GET("/", api.MainPage)

	// start the server
	err := r.Run(":5000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
