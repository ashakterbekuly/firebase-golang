package config

import (
	"bytes"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func SetupFirebaseAuth() *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}
	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}
	return auth
}

func SendAuthRequest(email, password string) error {
	// Authenticate the user with Firebase
	apiKey := "AIzaSyD8wNdQe6T36LKKzfzq01Zx1rwvk6Jty1c"
	signInURL := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", apiKey)
	signInBody := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}

	signInJSON, err := json.Marshal(signInBody)
	if err != nil {
		return err
	}

	signInReq, err := http.NewRequest("POST", signInURL, bytes.NewBuffer(signInJSON))
	if err != nil {
		return err
	}

	signInReq.Header.Set("Content-Type", "application/json")
	signInRes, err := http.DefaultClient.Do(signInReq)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println("error closing body: ", err)
			return
		}
	}(signInRes.Body)

	signInResJSON := make(map[string]interface{})
	err = json.NewDecoder(signInRes.Body).Decode(&signInResJSON)
	if err != nil {
		return err
	}

	if signInRes.StatusCode != http.StatusOK {
		errorMessage := signInResJSON["error"].(map[string]interface{})["message"].(string)
		log.Println(errorMessage)
		return err
	}
	//idToken := signInResJSON["idToken"].(string)

	return nil
}
