package test

import (
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	Result     string         `json:"result"`
	Media      []models.Media `json:"media" gorm:"foreignKey:TestResultID"`
	TestCaseID uint
}
