package models

import "time"

type Client struct {
	Name      string `form:"name"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	CreatedAt time.Time
}
