package database

import (
	"context"
	"firebase-golang/models"
	"log"
	"time"
)

func CreateClient(client models.Client) error {
	ref := Firestore.Collection("clients")
	client.CreatedAt = time.Now()

	_, _, err := ref.Add(context.TODO(), &client)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
