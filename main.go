package main

import (
	"firebase-golang/api"
	"firebase-golang/api/auth"
	"firebase-golang/database"
	"firebase-golang/firebase_auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
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
	r.LoadHTMLGlob("./static/*.html")
	r.Static("/css", "static/css")
	r.Static("/img", "static/img")
	r.Static("/js", "static/js")

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
	r.GET("/", api.MainPage)

	//events page
	r.GET("/events", api.EventsPage)

	//architecture page
	r.GET("/architecture", api.ArchitecturePage)

	//send mail
	r.GET("/sendmail", api.SendMail)

	//sign out
	r.GET("/logout", auth.LogoutGet)

	//projects
	r.GET("/projects", api.ProjectsGet)

	//profile
	r.GET("/profile", api.Profile)

	//edit profile
	r.POST("/edit-profile", auth.EditProfile)

	//edit architect profile
	r.POST("/edit-architect", auth.EditArchitectProfile)

	// start the server
	err := r.Run(":5002")
	if err != nil {
		log.Fatal(err)
		return
	}
}
