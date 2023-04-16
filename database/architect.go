package database

import (
	"context"
	"firebase-golang/models"
	"log"
	"time"
)

func CreateArchitect(arch models.Architect) error {
	ref := Firestore.Collection("architects")
	arch.CreatedAt = time.Now()

	_, _, err := ref.Add(context.TODO(), &arch)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
