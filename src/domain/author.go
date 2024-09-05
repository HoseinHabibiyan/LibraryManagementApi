package domain

import (
	"database/sql"
)

type Author struct {
	Id        int32        `json:"id" example:"1"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Birthdate sql.NullTime `json:"birthdate"`
	Email     string       `json:"email,omitempty"`
	Phone     string       `json:"phone,omitempty"`
	Address   string       `json:"address,omitempty"`
}
