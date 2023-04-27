package models

import (
	"time"
)

type Project struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	UserId      uint
	Team        []*User  `gorm:"many2many:project_team"`
	Suites      []*Suite `gorm:"many2many:project_suites"`
	Image       []byte
}
