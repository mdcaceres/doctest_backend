package models

import (
	"github.com/mdcaceres/doctest/models/test"
	"gorm.io/gorm"
)

type Bug struct {
	gorm.Model
	TestSuitID uint
	TestSuit   test.Suit
}
