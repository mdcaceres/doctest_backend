package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	Order       uint   `json:"order"`
	Description string `json:"description"`
	Result      string `json:"result"`
	CaseID      uint
}
