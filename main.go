package main

import (
	apiV0 "firebase-golang/api/v0"
	"firebase-golang/api/v0/auth"
	"firebase-golang/api/v0/edit_profile"
	apiV1 "firebase-golang/api/v1"
	"firebase-golang/database"
	"firebase-golang/firebase_auth"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	apiV0.InitUserState()
	database.InitFirebaseApp()

	// initialize new gin engine (for server)
	r := gin.Default()

	// load static files
	r.LoadHTMLGlob("./web/*.html")
	r.Static("/css", "web/css")
	r.Static("/img", "web/img")
	r.Static("/js", "web/js")

	// create/configure database instance
	db := database.InitFirestore()
	storage := database.InitStorage()
	firebaseAuth := firebase_auth.SetupFirebaseAuth()

	//set db and auth instance
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("storage", storage)
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
	r.GET("/", apiV0.MainPage)

	//events page
	r.GET("/events", apiV0.EventsPage)

	//architecture page
	r.GET("/architecture", apiV0.ArchitecturePage)

	//sign out
	r.GET("/logout", auth.Logout)

	//projects
	r.GET("/projects", apiV0.ProjectsGet)

	//profile
	r.GET("/profile", apiV0.Profile)

	//edit profile
	r.POST("/edit-profile", edit_profile.EditClient)

	//edit architects profile
	r.POST("/edit-architects", edit_profile.EditArchitectProfile)

	//templates
	r.GET("/templates", apiV0.Template)

	r.POST("/v1/login", apiV1.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}

	// start the server
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
