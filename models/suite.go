package models

import (
	"github.com/mdcaceres/doctest/models/test"
	"gorm.io/gorm"
)

type Suite struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Target      string       `json:"target"`
	UserId      uint         `json:"user_id"`
	ProjectId   uint         `json:"project_id"`
	TestCases   []*test.Case `json:"test_cases,omitempty" gorm:"many2many:testcases_testsuite"`
}
