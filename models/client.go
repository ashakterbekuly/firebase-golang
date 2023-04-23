package models

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID       string
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Bio      string `form:"bio"`
	Password string `form:"password"`
	PhotoUrl string
}

func (c Client) GetBio() string {
	if c.Bio == "" {
		c.Bio = "Short description about yourself, what you do and so on."
	}

	return c.Bio
}

func (c Client) CreateID() string {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return c.ID
}

type ClientDao struct {
	ID        string    `json:"ID"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	PhotoUrl  string    `json:"photoUrl"`
	CreatedAt time.Time `json:"createdAt"`
}

func DaoFromClient(c Client) ClientDao {
	return ClientDao{
		ID:        c.CreateID(),
		Name:      c.Name,
		Phone:     c.Phone,
		Email:     c.Email,
		Bio:       c.GetBio(),
		PhotoUrl:  c.PhotoUrl,
		CreatedAt: time.Now(),
	}
}
