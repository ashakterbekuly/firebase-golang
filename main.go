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
	//r.LoadHTMLGlob("web/templates/*")
	//r.Static("/static", "web/static/")
	r.LoadHTMLGlob("./static/*.html")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")

	//r.StaticFile("/logo.png", "./static/img/logo.png")
	//r.StaticFile("/hero.png", "./static/img/hero.png")
	//r.StaticFile("/Sign-in.png", "./static/img/Sign-in.png")
	//r.StaticFile("/Sign-up.png", "./static/img/Sign-up.png")

	// create/configure database instance
	db := database.CreateDatabase()
	auth := config.SetupFirebase()
	//firebaseAuth := config.SetupFirebase()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", auth)
	})

	// routes definition for finding and creating artists
	r.GET("/artist", api.FindArtistsByDocID)
	r.GET("/artists", api.FindArtists)
	r.GET("/documents", api.GetDocIDs)
	r.POST("/artist", api.CreateArtist)

	r.GET("/register/arch", api.RegisterArchitectGet)
	r.POST("/register/arch", api.RegisterArchitect)
	r.GET("/register/role", api.RegisterRoleGet)

	r.GET("/login", api.LoginArtistGet)
	r.POST("/login", api.LoginArtist)
	// start the server
	err := r.Run(":5000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
