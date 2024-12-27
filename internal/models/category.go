package models

import "time"

type Category struct {
	ID        int64     `json:"category_id"`
	Name      string    `json:"category_name"`
	CreatedAt time.Time `json:"category_created_at"`
	UpdatedAt time.Time `json:"category_updated_at"`
	Version   int       `json:"category_version"`
}

func NewCategory() *Category {
	return &Category{
		Name:      "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type CategoryPostJoin struct {
	ID     int64
	CatID  int64
	PostID int64
}

func NewCategoryPostJoin(catID, postID int64) *CategoryPostJoin {
	return &CategoryPostJoin{
		CatID:  catID,
		PostID: postID,
	}
}
