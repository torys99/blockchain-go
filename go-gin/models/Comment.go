package models

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text" json:"content"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
