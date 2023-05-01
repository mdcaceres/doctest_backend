package models

import "gorm.io/gorm"

const (
	Admin  = "ADMINISTRATOR"
	Tester = "TESTER"
	Client = "CLIENT"
)

var (
	roles = map[string]Role{
		"1": Role{Name: Admin},
		"2": Role{Name: Tester},
		"3": Role{Name: Client},
	}

	AvailableRoles = []string{Admin, Tester, Client}
)

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_roles"`
}

func GetRoleByName(name string) Role {
	return roles[name]
}
