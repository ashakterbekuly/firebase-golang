package models

import "time"

type Event struct {
	Title              string
	Description        string
	ImageUrl           string
	Code               string
	DateFrom           time.Time
	DateTo             time.Time
	DateFromString     string
	DateToString       string
	DateTimeFromString string
	DateTimeToString   string
	Venue              string
	Subjects           string
}
