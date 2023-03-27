package models

import (
	"github.com/mdcaceres/doctest/models/test"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Users       []*User      `json:"users,omitempty" gorm:"many2many:user_projects;"`
	TestSuit    []*test.Suit `json:"test_suits" gorm:"many2many:testsuit_projects"`
}
