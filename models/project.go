package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string
	Description string
	UserId      uint
	Team        []*User  `gorm:"many2many:project_team"`
	Suites      []*Suite `gorm:"many2many:project_suites"`
	Image       []byte
}
