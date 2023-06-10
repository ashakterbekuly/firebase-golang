package models

import "time"

type Architect struct {
	Name           string `form:"name" json:"name"`
	Specialization string `form:"spec" json:"spec"`
	Portfolio      string `form:"portfolio" json:"portfolio"`
	Phone          string `form:"phone" json:"phone"`
	Email          string `form:"email" json:"email"`
	Password       string `form:"password" json:"password"`
	Bio            string
	PhotoUrl       string
}

func (a Architect) GetBio() string {
	if a.Bio == "" {
		a.Bio = "Short description about yourself, what you do and so on."
	}

	return a.Bio
}

type ArchitectDao struct {
	Name           string    `json:"name"`
	Specialization string    `json:"spec"`
	Portfolio      string    `json:"portfolio"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Bio            string    `json:"bio"`
	PhotoUrl       string    `json:"photoUrl"`
	CreatedAt      time.Time `json:"createdAt"`
}

func DaoFromArchitect(a Architect) ArchitectDao {
	return ArchitectDao{
		Name:           a.Name,
		Specialization: a.Specialization,
		Portfolio:      a.Portfolio,
		Phone:          a.Phone,
		Email:          a.Email,
		PhotoUrl:       a.PhotoUrl,
		Bio:            a.GetBio(),
		CreatedAt:      time.Now(),
	}
}
