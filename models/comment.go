package models

import "gorm.io/gorm"

type BugComment struct {
	gorm.Model
	Comment string `json:"comment"`
	BugID   uint   `json:"bug_id"`
	UserID  uint   `json:"user_id"`
}

type TestComment struct {
	gorm.Model
	Comment string `json:"comment"`
	CaseID  uint   `json:"test_id"`
	UserID  uint   `json:"user_id"`
}
