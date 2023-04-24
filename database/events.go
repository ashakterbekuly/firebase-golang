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

	return formatDateTimeForEvents(events), nil
}

func formatDateTimeForEvents(events []models.Event) []models.Event {
	for i := range events {
		events[i].DateFromString = events[i].DateFrom.Format("2006-01-02")
		events[i].DateToString = events[i].DateTo.Format("2006-01-02")
		events[i].DateTimeFromString = events[i].DateFrom.Format("2006-01-02 15:04:05")
		events[i].DateTimeToString = events[i].DateTo.Format("2006-01-02 15:04:05")
	}

	return events
}
