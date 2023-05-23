package models

import (
	"gorm.io/gorm"
)

type ProjectClient struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string `json:"phone"`
	UserId      uint
}
