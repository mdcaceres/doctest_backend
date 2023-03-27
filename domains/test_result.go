package domains

import "gorm.io/gorm"

type TestResult struct {
	gorm.Model
	Result     string  `json:"result"`
	Media      []Media `json:"media" gorm:"foreignKey:TestResultID"`
	TestCaseID uint
}
