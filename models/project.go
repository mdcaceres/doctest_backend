package models

import "time"

type Project struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	StartDate       time.Time
	EndDate         time.Time
	Name            string
	Description     string
	UserId          uint
	Team            []*User `gorm:"many2many:project_team"`
	Suites          []*Suite
	Cases           []*Case
	Bugs            []*Bug `gorm:"foreignKey:ProjectID"`
	Image           string
	ProjectClientID uint
}
