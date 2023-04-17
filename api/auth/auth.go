package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sendAuthRequest(email, password string) (string, error) {
	// Authenticate the user with Firebase
	apiKey := "AIzaSyBcT-aXVJ41Nbgg0x78wphWkJ2GXDvUHuA"
	signInURL := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", apiKey)
	signInBody := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}

	signInJSON, err := json.Marshal(signInBody)
	if err != nil {
		return "", err
	}

	signInReq, err := http.NewRequest("POST", signInURL, bytes.NewBuffer(signInJSON))
	if err != nil {
		return "", err
	}

	signInReq.Header.Set("Content-Type", "application/json")
	signInRes, err := http.DefaultClient.Do(signInReq)
	if err != nil {
		return "", err
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
		return "", err
	}

	if signInRes.StatusCode != http.StatusOK {
		errorMessage := signInResJSON["error"].(map[string]interface{})["message"].(string)
		return "", errors.New(errorMessage)
	}
	idToken := signInResJSON["idToken"].(string)

	return idToken, nil
}
