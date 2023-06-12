package models

type Template struct {
	Title         string `json:"title"`
	Specification string `json:"specification"`
	ImageUrl      string `json:"imageUrl"`
	Description   string `json:"description"`
	AuthorID      string `json:"authorId"`
	AuthorName    string `json:"authorName"`
}
