package models

import "time"

type Architect struct {
	Name          string `form:"name"`
	Specification string `form:"spec"`
	Portfolio     string `form:"portfolio"`
	Phone         string `form:"phone"`
	Email         string `form:"email"`
	Password      string `form:"password"`
	Bio           string
}

func (a Architect) GetBio() string {
	if a.Bio == "" {
		a.Bio = "Short description about yourself, what you do and so on."
	}

	return a.Bio
}

type ArchitectDao struct {
	Name          string    `json:"name"`
	Specification string    `json:"spec"`
	Portfolio     string    `json:"portfolio"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Bio           string    `json:"bio"`
	CreatedAt     time.Time `json:"createdAt"`
}

func DaoFromArchitect(a Architect) ArchitectDao {
	return ArchitectDao{
		Name:          a.Name,
		Specification: a.Specification,
		Portfolio:     a.Portfolio,
		Phone:         a.Phone,
		Email:         a.Email,
		Bio:           a.GetBio(),
		CreatedAt:     time.Now(),
	}
}
