package models

import "time"

type Architect struct {
	Name          string `form:"name"`
	Specification string `form:"spec"`
	Portfolio     string `form:"portfolio"`
	Phone         string `form:"phone"`
	Email         string `form:"email"`
	Password      string `form:"password"`
	CreatedAt     time.Time
}
