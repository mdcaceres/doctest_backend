package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	Order          uint   `json:"order"`
	Step           string `json:"name"`
	ExpectedResult string `json:"expected_result"`
	CaseID         uint
}
