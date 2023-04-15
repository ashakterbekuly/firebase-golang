package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Artist struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type Document struct {
	ID string `json:"id" binding:"required"`
}

// CreateArtistInput : struct for create art post request
type CreateArtistInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// FindArtists : Controller for getting all artists
func FindArtists(c *gin.Context) {
	client := c.MustGet("db").(*firestore.Client)
	var artists []Artist
	coll := client.Collection("artist")
	documents, err := coll.Documents(context.TODO()).GetAll()
	if err != nil {
		log.Println("error in reading from firebase DB: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, document := range documents {
		var artist Artist
		err = document.DataTo(&artist)
		if err != nil {
			log.Println("error data to: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		artists = append(artists, artist)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "golang-firebase",
		"Artists": artists,
	})
}

// FindArtistsByDocID : Controller for getting all artists
func FindArtistsByDocID(c *gin.Context) {
	client := c.MustGet("db").(*firestore.Client)
	var document Document
	err := c.ShouldBindJSON(&document)
	if err != nil {
		log.Println("error in reading from firebase DB: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var artist Artist
	coll := client.Collection("artist")
	res, err := coll.Doc(document.ID).Get(context.TODO())
	if err != nil {
		log.Println("error in reading from firebase DB: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = res.DataTo(&artist)
	if err != nil {
		log.Println("error data to: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": artist})
}

// GetDocIDs : Controller for getting all artists
func GetDocIDs(c *gin.Context) {
	client := c.MustGet("db").(*firestore.Client)
	coll := client.Collection("artist")
	documents := coll.Documents(context.TODO())
	res, err := documents.GetAll()
	if err != nil {
		log.Println("error data to: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var ids []Document

	for _, document := range res {
		ids = append(ids, Document{ID: document.Ref.ID})
	}

	c.JSON(http.StatusOK, gin.H{"data": ids})
}

// CreateArtist : controller for creating new artists
func CreateArtist(c *gin.Context) {
	client := c.MustGet("db").(*firestore.Client)
	// Validate input
	var input CreateArtistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create artist
	artist := Artist{Name: input.Name, Email: input.Email, CreatedAt: time.Now()}
	ref := client.Collection("artist")

	var doc *firestore.DocumentRef
	var err error

	if doc, _, err = ref.Add(context.TODO(), &artist); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{
		"docID":  doc.ID,
		"artist": artist,
	}})
}
