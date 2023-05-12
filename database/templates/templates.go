package templates

import (
	"context"
	"firebase-golang/database"
	"firebase-golang/models"
)

func GetTemplatesList() ([]models.Template, error) {
	var templates []models.Template
	coll := database.Firestore.Collection("templates")

	documents, err := coll.Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}

	for _, document := range documents {
		var template models.Template
		err = document.DataTo(&template)
		if err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}

	return templates, nil
}

func SortTemplates(templates []models.Template) []models.Template {
	var sortedTemplates []models.Template

	isHouseDraftFound := false
	isInteriorDesignFound := false
	isUrbanProjectFound := false
	isReconstructionFound := false

	for _, template := range templates {
		if !isHouseDraftFound && template.Specification == "House Draft" {
			sortedTemplates = append(sortedTemplates, template)
			isHouseDraftFound = true
		}
		if !isInteriorDesignFound && template.Specification == "Interior Design" {
			sortedTemplates = append(sortedTemplates, template)
			isInteriorDesignFound = true
		}
		if !isUrbanProjectFound && template.Specification == "Urban Project" {
			sortedTemplates = append(sortedTemplates, template)
			isUrbanProjectFound = true
		}
		if !isReconstructionFound && template.Specification == "Reconstruction" {
			sortedTemplates = append(sortedTemplates, template)
			isReconstructionFound = true
		}
	}

	return sortedTemplates
}

func GetHouseDrafts(templates []models.Template) []models.Template {
	var houseDrafts []models.Template
	for _, template := range templates {
		if template.Specification == "House Draft" {
			houseDrafts = append(houseDrafts, template)
		}
	}
	return houseDrafts
}

func GetInteriorDesigns(templates []models.Template) []models.Template {
	var interiorDesigns []models.Template
	for _, template := range templates {
		if template.Specification == "Interior Design" {
			interiorDesigns = append(interiorDesigns, template)
		}
	}
	return interiorDesigns
}

func GetUrbanProjects(templates []models.Template) []models.Template {
	var urbanProjects []models.Template
	for _, template := range templates {
		if template.Specification == "Urban Project" {
			urbanProjects = append(urbanProjects, template)
		}
	}
	return urbanProjects
}

func GetReconstructionProjects(templates []models.Template) []models.Template {
	var reconstructionProjects []models.Template
	for _, template := range templates {
		if template.Specification == "Reconstruction" {
			reconstructionProjects = append(reconstructionProjects, template)
		}
	}
	return reconstructionProjects
}
