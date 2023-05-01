package main

import (
	"firebase-golang/api"
	"firebase-golang/api/auth"
	"firebase-golang/api/edit_profile"
	"firebase-golang/database"
	"firebase-golang/firebase_auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	api.InitUserState()
	database.InitFirebaseApp()

	// initialize new gin engine (for server)
	r := gin.Default()

	// Инициализация сессий
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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
	r.GET("/register/roles", auth.RegisterRoleGet)
	r.GET("/register/arch", auth.RegisterArchitectGet)
	r.POST("/register/arch", auth.RegisterArchitect)
	r.GET("/register/clients", auth.RegisterClientGet)
	r.POST("/register/clients", auth.RegisterClient)

	//login routes
	r.GET("/login", auth.LoginGet)
	r.POST("/login", auth.Login)

	//main page
	r.GET("/:uid", api.MainPage)

	//events page
	r.GET("/events", api.EventsPage)

	//architecture page
	r.GET("/architecture", api.ArchitecturePage)

	//sign out
	r.GET("/logout", auth.Logout)

	//projects
	r.GET("/projects", api.ProjectsGet)

	//profile
	r.GET("/profile", api.Profile)

	//edit profile
	r.POST("/edit-profile", edit_profile.EditClient)

	//edit architects profile
	r.POST("/edit-architects", edit_profile.EditArchitectProfile)

	//templates
	r.GET("/templates", api.Template)

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
