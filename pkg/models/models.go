package models

import "time"

type Note struct {
	ID        int       `json: "id"`
	Content   string    `json: "content"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	DeletedAt time.Time `json: "deleted_at"`
}

var Notes []Note

type Config struct {
	Host string
	Port string
}
