package test

import (
	"github.com/mdcaceres/doctest/models/media"
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	Result     string        `json:"result"`
	Media      []media.Media `json:"media" gorm:"foreignKey:TestResultID"`
	TestCaseID uint
}
