package domain

import (
	"database/sql"
)

type Author struct {
	Id        int32
	FirstName string
	LastName  string
	Birthdate sql.NullTime
	Email     string
	Phone     string
	Address   string
}
