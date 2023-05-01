package firebase_auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

var authClient *auth.Client

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
	authClient, err = app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}
	return authClient
}

func GetUserToken(email, password string) (string, error) {
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

func UpdateUserRequest(token, newEmail, newPassword string) error {
	// Формируем запрос на изменение учетной записи пользователя
	body := map[string]interface{}{
		"idToken":           token,
		"email":             newEmail,
		"password":          newPassword,
		"returnSecureToken": true,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	apiKey := "AIzaSyBcT-aXVJ41Nbgg0x78wphWkJ2GXDvUHuA"

	// отправляем запрос на изменение учетной записи пользователя
	resp, err := http.Post(fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:update?key=%s", apiKey), "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	// Обрабатываем ответ сервера
	if resp.StatusCode != http.StatusOK {
		var errorResponse struct {
			Error struct {
				Message string
			}
		}
		if err = json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return err
		}
		return err
	}

	return nil
}
