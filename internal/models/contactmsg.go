package models

import "time"

type ContactMsg struct {
	ID        int64
	Email     string
	Message   string
	CreatedAt time.Time
	Version   int
}
