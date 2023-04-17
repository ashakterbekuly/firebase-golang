package firebase_auth

import (
	"context"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"log"
)

func SendMessage(authClient *auth.Client, email, event string) {
	user, err := authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Fatalf("error getting user: %v\n", err)
	}
	email = user.Email

	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Registration to the event",
			Body:  "You have invitation to the event named: " + event,
		},
		Token: email,
	}

	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	response, err := messagingClient.Send(context.Background(), msg)
	if err != nil {
		log.Fatalf("error sending message: %v\n", err)
	}

	log.Printf("message sent: %v\n", response)

}
