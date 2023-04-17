package database

import (
	"context"
	"firebase-golang/models"
)

func GetEventsList() ([]models.Event, error) {
	var events []models.Event
	coll := Firestore.Collection("events")

	documents, err := coll.Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}

	for _, document := range documents {
		var event models.Event
		err = document.DataTo(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
