package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

type Role struct {
	Email string
	Role  string
}

func SetRoleByEmail(email, role string) error {
	ref := Firestore.Collection("roles")

	_, _, err := ref.Add(context.TODO(), &Role{
		Role:  role,
		Email: email,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateRoleByEmail(oldEmail, newEmail, role string) error {
	_, err := Firestore.
		Collection("roles").
		Doc(GetRoleDocumentIDByEmail(oldEmail)).
		Set(context.TODO(), map[string]interface{}{
			"Email": newEmail,
			"Role":  role,
		}, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func GetRoleDocumentIDByEmail(email string) string {
	var docID string

	ref := Firestore.Collection("roles")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		log.Println(err)
		return ""
	}

	for _, doc := range docs {
		if email == doc.Data()["Email"].(string) {
			docID = doc.Ref.ID
		}
	}

	return docID
}

func GetRoleByEmail(email string) string {
	var role string

	ref := Firestore.Collection("roles")

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
			role, ok = doc.Data()["Role"].(string)
		}
	}

	return role
}
