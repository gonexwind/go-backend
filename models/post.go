package models

import "time"

type Post struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	OwnerID   uint      `json:"ownerId"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PostID"`
}
