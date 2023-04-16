package database

import (
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
