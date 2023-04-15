package models

import "gorm.io/gorm"

const (
	Admin  = "ADMINISTRATOR"
	Tester = "TESTER"
	Client = "CLIENT"
)

var AvailableRoles = []string{Admin, Tester, Client}

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_roles"`
}
