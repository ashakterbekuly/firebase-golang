package main

import (
	"firebase-golang/api"
	"firebase-golang/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// initialize new gin engine (for server)
	r := gin.Default()
	// create/configure database instance
	db := config.CreateDatabase()
	//firebaseAuth := config.SetupFirebase()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		//c.Set("firebaseAuth", firebaseAuth)
	})

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
