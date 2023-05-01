package roles

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase-golang/database"
	"firebase-golang/database/architects"
	"firebase-golang/database/clients"
	"log"
)

type Role struct {
	ID    string
	Email string
	Role  string
}

func SetRoleByEmail(email, role, uid string) error {
	ref := database.Firestore.Collection("roles")

	_, _, err := ref.Add(context.TODO(), &Role{
		ID:    uid,
		Role:  role,
		Email: email,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateRoleByEmail(oldEmail, newEmail, role string) error {
	_, err := database.Firestore.
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

	ref := database.Firestore.Collection("roles")

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

func GetRoleByUID(uid string) string {
	var role string

	ref := database.Firestore.Collection("roles")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		log.Println(err)
		return ""
	}

	for _, doc := range docs {
		rawID, ok := doc.Data()["ID"].(string)
		if !ok {
			log.Println(err)
			continue
		}

		if uid == rawID {
			role, ok = doc.Data()["Role"].(string)
		}
	}

	return role
}

func GetEmailByUID(uid string) string {
	var email string

	ref := database.Firestore.Collection("roles")

	docs, err := ref.Documents(context.TODO()).GetAll()
	if err != nil {
		log.Println(err)
		return ""
	}

	for _, doc := range docs {
		rawUID, ok := doc.Data()["ID"].(string)
		if !ok {
			log.Println(err)
			continue
		}

		if uid == rawUID {
			email, ok = doc.Data()["Email"].(string)
		}
	}

	return email
}

func GetUsernameByUID(uid string) string {
	role := GetRoleByUID(uid)

	var username string
	if role == "client" {
		username = clients.GetUsernameByUID(uid)
	} else {
		username = architects.GetUsernameByUID(uid)
	}

	return username
}

func GetUIDByEmail(email string) string {
	var uid string

	ref := database.Firestore.Collection("roles")

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
			uid, ok = doc.Data()["ID"].(string)
		}
	}

	return uid
}
