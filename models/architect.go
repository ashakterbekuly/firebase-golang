package models

import "time"

type Architect struct {
	Name          string `form:"name"`
	Specification string `form:"spec"`
	Portfolio     string `form:"portfolio"`
	Phone         string `form:"phone"`
	Email         string `form:"email"`
	Password      string `form:"password"`
}

type ArchitectDao struct {
	Name          string    `json:"name"`
	Specification string    `json:"spec"`
	Portfolio     string    `json:"portfolio"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"createdAt"`
}

func DaoFromArchitect(a Architect) ArchitectDao {
	return ArchitectDao{
		Name:          a.Name,
		Specification: a.Specification,
		Portfolio:     a.Portfolio,
		Phone:         a.Phone,
		Email:         a.Email,
		CreatedAt:     time.Now(),
	}
}
