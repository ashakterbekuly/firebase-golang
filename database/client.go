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

func GetUsername(email string) string {
	var username string

	ref := Firestore.Collection("clients")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		log.Println(err)
		return email
	}

	for _, doc := range docs {
		rawEmail, ok := doc.Data()["Email"].(string)
		if !ok {
			log.Println(err)
			continue
		}

		if email == rawEmail {
			username, ok = doc.Data()["Name"].(string)
			if !ok {
				log.Println()
				continue
			}
		}
	}

	if username == "" {
		return email
	}

	return username
}
