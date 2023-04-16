package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginArtistGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginArtist(c *gin.Context) {
	//firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	artist := ArtistRegister{}

	err := c.Bind(&artist)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if artist.Email == "" && artist.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Authenticate the user with Firebase
	apiKey := "AIzaSyD8wNdQe6T36LKKzfzq01Zx1rwvk6Jty1c"
	signInURL := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", apiKey)
	signInBody := map[string]string{
		"email":             artist.Email,
		"password":          artist.Password,
		"returnSecureToken": "true",
	}
	signInJSON, err := json.Marshal(signInBody)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	signInReq, err := http.NewRequest("POST", signInURL, bytes.NewBuffer(signInJSON))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	signInReq.Header.Set("Content-Type", "application/json")
	signInRes, err := http.DefaultClient.Do(signInReq)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer signInRes.Body.Close()
	signInResJSON := make(map[string]interface{})
	err = json.NewDecoder(signInRes.Body).Decode(&signInResJSON)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if signInRes.StatusCode != http.StatusOK {
		errorMessage := signInResJSON["error"].(map[string]interface{})["message"].(string)
		log.Println(errorMessage)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	idToken := signInResJSON["idToken"].(string)

	log.Println(idToken)

	// Redirect the user to the dashboard or home page
	c.Redirect(http.StatusFound, "/artists")
}
