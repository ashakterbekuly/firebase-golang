package models

import "time"

type Event struct {
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	ImageUrl           string    `json:"imageUrl"`
	Code               string    `json:"code"`
	DateFrom           time.Time `json:"dateFrom"`
	DateTo             time.Time `json:"dateTo"`
	DateFromString     string    `json:"dateFromString"`
	DateToString       string    `json:"dateToString"`
	DateTimeFromString string    `json:"dateTimeFromString"`
	DateTimeToString   string    `json:"dateTimeToString"`
	Venue              string    `json:"venue"`
	Subjects           string    `json:"subjects"`
}
