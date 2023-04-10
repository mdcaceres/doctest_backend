package models

import "gorm.io/gorm"

const (
	Admin  = "ADMINISTRATOR"
	user   = "USER"
	tester = "TESTER"
	client = "CLIENT"
)

var AvailableRoles = []string{Admin, user, tester, client}

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_roles"`
}
