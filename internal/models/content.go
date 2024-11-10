package models

import (
	"time"
)

type Content struct {
	ID          int64
	UserId      int64
	UserEmail   string
	Title       string
	Slug        string
	Description string
	Content     string
	Delta       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Version     int64
	Status      string
}
