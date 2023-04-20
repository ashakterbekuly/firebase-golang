package firebase_auth

import (
	"bytes"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Recipient struct {
	Email string `json:"email"`
}

type Message struct {
	Text    string      `json:"text"`
	Subject string      `json:"subject"`
	From    string      `json:"from_email"`
	To      []Recipient `json:"to"`
}

type APIResponse struct {
	Status string `json:"status"`
	Id     string `json:"id"`
	Error  string `json:"error"`
}

func SendMessage(authClient *auth.Client, email, event, token string) {
	// Настройки сообщения
	message := Message{
		Text:    "This is a test email",
		Subject: "Test Email",
		From:    "sender@example.com",
		To: []Recipient{
			{Email: email},
		},
	}

	// Кодируем сообщение в формат JSON
	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v\n", err)
		return
	}

	// Создаем запрос к API Mailchimp
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/messages", "ba44a9569197bb47f98b0a48135e49f7-us10", "")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageJson))
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос и получаем ответ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v\n", err)
		}
	}(resp.Body)

	// Декодируем ответ в формат JSON
	var apiResponse APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		log.Printf("Error decoding response: %v\n", err)
		return
	}

	// Проверяем статус ответа
	if apiResponse.Status == "error" {
		log.Printf("Error sending email: %s\n", apiResponse.Error)
		return
	}

	log.Printf("Email sent with ID %s\n", apiResponse.Id)
}
