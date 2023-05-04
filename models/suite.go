package models

import (
	"gorm.io/gorm"
)

type Suite struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Target      string `json:"target"`
	UserId      uint   `json:"user_id"`
	ProjectId   uint   `json:"project_id"`
	TestCases   []*Case
}
