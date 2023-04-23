package database

import (
	"cloud.google.com/go/firestore"
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

func UpdateClient(oldEmail string, client models.Client) error {
	_, err := Firestore.
		Collection("clients").
		Doc(GetDocumentIDByEmail(oldEmail)).
		Set(context.TODO(), map[string]interface{}{
			"Email":    client.Email,
			"Bio":      client.Bio,
			"Name":     client.Name,
			"PhotoUrl": client.PhotoUrl,
		}, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func GetClientByEmail(email string) (models.Client, error) {
	ref := Firestore.Collection("clients")
	doc := ref.Doc(GetDocumentIDByEmail(email))

	snapshot, err := doc.Get(context.TODO())
	if err != nil {
		return models.Client{}, err
	}

	var client models.Client

	err = snapshot.DataTo(&client)
	if err != nil {
		return models.Client{}, err
	}

	return client, nil
}

func GetUsername(email string) string {
	client, err := GetClientByEmail(email)
	if err != nil {
		log.Println(err)
		return ""
	}

	return client.Name
}

func GetID(email string) string {
	client, err := GetClientByEmail(email)
	if err != nil {
		log.Println(err)
		return ""
	}

	return client.ID
}

func GetPhotoUrl(email string) string {
	client, err := GetClientByEmail(email)
	if err != nil {
		log.Println(err)
		return ""
	}

	return client.PhotoUrl
}

func GetDocumentIDByEmail(email string) string {
	var id string

	ref := Firestore.Collection("clients")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		log.Println(err)
		return ""
	}

	for _, doc := range docs {
		rawEmail, ok := doc.Data()["Email"].(string)
		if !ok {
			log.Println(err)
			continue
		}

		if email == rawEmail {
			id = doc.Ref.ID
		}
	}

	return id
}

func GetClientByID(id string) (models.Client, error) {
	var client models.Client

	ref := Firestore.Collection("clients")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		return client, err
	}

	for _, doc := range docs {
		rawID, ok := doc.Data()["ID"].(string)
		if !ok {
			return client, err
		}

		if id == rawID {
			err = doc.DataTo(&client)
			if err != nil {
				return client, err
			}
		}
	}

	return client, nil
}
