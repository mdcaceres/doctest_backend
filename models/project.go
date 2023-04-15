package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UserId      uint     `json:"owner_id"`
	Team        []*User  `json:"users,omitempty" gorm:"many2many:project_team"`
	Suites      []*Suite `json:"suites,omitempty" gorm:"many2many:project_suites"`
	Image       []byte   `json:"image"`
}
