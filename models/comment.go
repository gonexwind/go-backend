package models

import "time"

type Comment struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
}
