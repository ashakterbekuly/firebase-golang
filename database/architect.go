package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase-golang/models"
	"log"
)

func CreateArchitect(arch models.ArchitectDao) error {
	ref := Firestore.Collection("architects")

	_, _, err := ref.Add(context.TODO(), &arch)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateArchitect(oldEmail string, architect models.Architect) error {
	_, err := Firestore.
		Collection("architects").
		Doc(GetArchitectDocumentIDByEmail(oldEmail)).
		Set(context.TODO(), map[string]interface{}{
			"Email":          architect.Email,
			"Bio":            architect.Bio,
			"Name":           architect.Name,
			"PhotoUrl":       architect.PhotoUrl,
			"Portfolio":      architect.Portfolio,
			"Specialization": architect.Specialization,
		}, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func GetArchitectByEmail(email string) (models.Architect, error) {
	ref := Firestore.Collection("architects")
	doc := ref.Doc(GetArchitectDocumentIDByEmail(email))

	snapshot, err := doc.Get(context.TODO())
	if err != nil {
		return models.Architect{}, err
	}

	var architect models.Architect

	err = snapshot.DataTo(&architect)
	if err != nil {
		return models.Architect{}, err
	}

	return architect, nil
}

func GetArchitectUsername(email string) string {
	architect, err := GetArchitectByEmail(email)
	if err != nil {
		log.Println(err)
		return ""
	}

	return architect.Name
}

func GetArchitectPhotoUrl(email string) string {
	architect, err := GetArchitectByEmail(email)
	if err != nil {
		log.Println(err)
		return ""
	}

	return architect.PhotoUrl
}

func GetArchitectDocumentIDByEmail(email string) string {
	var id string

	ref := Firestore.Collection("architects")

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

func GetArchitectByID(id string) (models.Architect, error) {
	var architect models.Architect

	ref := Firestore.Collection("architects")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		return architect, err
	}

	for _, doc := range docs {
		rawID, ok := doc.Data()["ID"].(string)
		if !ok {
			return architect, err
		}

		if id == rawID {
			err = doc.DataTo(&architect)
			if err != nil {
				return architect, err
			}
		}
	}

	return architect, nil
}
