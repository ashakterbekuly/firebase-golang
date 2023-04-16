package models

import "time"

type Client struct {
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type ClientDao struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func DaoFromClient(c Client) ClientDao {
	return ClientDao{
		Name:      c.Name,
		Phone:     c.Phone,
		Email:     c.Email,
		CreatedAt: time.Now(),
	}
}
