package projects

import (
	"context"
	"firebase-golang/database"
	"firebase-golang/models"
)

func GetProjectsList() ([]models.Project, error) {
	var projects []models.Project
	coll := database.Firestore.Collection("projects")

	documents, err := coll.Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}

	for _, document := range documents {
		var project models.Project
		err = document.DataTo(&project)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}
