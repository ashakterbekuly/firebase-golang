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

	r.LoadHTMLGlob("./static/*.html")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")

	// create/configure database instance
	db := database.CreateDatabase()
	auth := config.SetupFirebaseAuth()

	//firebaseAuth := config.SetupFirebaseAuth()
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

	// routes definition for finding and creating artists
	r.GET("/artist", api.FindArtistsByDocID)
	r.GET("/artists", api.FindArtists)
	r.GET("/documents", api.GetDocIDs)
	r.POST("/artist", api.CreateArtist)

	// start the server
	err := r.Run(":5000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
