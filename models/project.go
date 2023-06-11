package models

type Project struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	Description     string `json:"description"`
	ProjectImageUrl string `json:"projectImageUrl"`
	AuthorImageUrl  string `json:"authorImageUrl"`
}
