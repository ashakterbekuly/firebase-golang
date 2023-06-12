package main

import (
	apiV0 "firebase-golang/api/v0"
	"firebase-golang/api/v0/auth"
	"firebase-golang/api/v0/edit_profile"
	apiV1 "firebase-golang/api/v1"
	"firebase-golang/database"
	"firebase-golang/firebase_auth"
	"firebase-golang/middleware"
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
	r.Use(middleware.CorsMiddleware(), func(c *gin.Context) {
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

	v1 := r.Group("/v1")

	v1.GET("/", apiV1.Main)

	v1.POST("/login", apiV1.Login)

	v1.POST("/register-arch", apiV1.RegisterArchitect)

	v1.POST("/register-client", apiV1.RegisterClient)

	v1.GET("/events", apiV1.Events)

	v1.GET("/projects", apiV1.Projects)

	v1.GET("/templates", apiV1.Templates)

	v1.GET("/interior-designs", apiV1.InteriorDesigns)

	v1.GET("/house-drafts", apiV1.HouseDrafts)

	v1.GET("/urban-projects", apiV1.UrbanProjects)

	v1.GET("/reconstructions", apiV1.Reconstructions)

	v1.POST("/me", middleware.TokenCheck, apiV1.Profile)

	v1.GET("/architect", apiV1.Architect)

	v1.POST("/edit-arch", middleware.TokenCheck, apiV1.EditArchitectProfile)

	v1.POST("/edit-client", middleware.TokenCheck, apiV1.EditClient)

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
