package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Body      string
}
