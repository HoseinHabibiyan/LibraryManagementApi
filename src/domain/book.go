package domain

import "time"

type Book struct {
	Id          int32
	Title       string
	Summery     string
	Description string
	ISBN        string
	Author      Author
	Publisher   Publisher
	PublishDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
