package architects

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase-golang/database"
	"firebase-golang/models"
	"log"
)

func CreateArchitect(arch models.ArchitectDao) string {
	ref := database.Firestore.Collection("architects")

	doc, _, err := ref.Add(context.TODO(), &arch)
	if err != nil {
		log.Println(err)
		return ""
	}

	return doc.ID
}

func UpdateArchitect(uid string, architect models.Architect) error {
	_, err := database.Firestore.
		Collection("architects").
		Doc(uid).
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

func GetUsernameByUID(uid string) string {
	doc, err := database.Firestore.
		Collection("architects").
		Doc(uid).Get(context.TODO())
	if err != nil {
		log.Println(err)
		return ""
	}

	return doc.Data()["Name"].(string)
}

func GetArchitectPhotoUrl(uid string) string {
	architect, err := GetArchitectByUID(uid)
	if err != nil {
		log.Println(err)
		return ""
	}

	return architect.PhotoUrl
}

func GetArchitectByUID(uid string) (models.Architect, error) {
	var architect models.Architect

	ref := database.Firestore.Collection("architects")

	docRef, err := ref.Doc(uid).Get(context.Background())
	if err != nil {
		return architect, err
	}

	err = docRef.DataTo(&architect)
	if err != nil {
		return architect, err
	}

	return architect, nil
}
