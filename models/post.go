package models

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Comment   string `json:"comment"`
	ProjectID uint   `json:"project_id"`
	UserID    uint   `json:"user_id"`
	Image     string
}
