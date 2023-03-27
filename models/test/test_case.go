package test

import (
	"gorm.io/gorm"
)

type Case struct {
	gorm.Model
	TestSteps  []*Step
	TestResult Result
}
