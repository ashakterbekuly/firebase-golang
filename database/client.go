package database

import (
	"context"
	"firebase-golang/models"
	"log"
)

func CreateClient(client models.ClientDao) error {
	ref := Firestore.Collection("clients")

	_, _, err := ref.Add(context.TODO(), &client)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
