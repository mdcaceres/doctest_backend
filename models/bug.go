package models

import (
	"github.com/mdcaceres/doctest/models/media"
	"gorm.io/gorm"
	"time"
)

type Bug struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	ProjectID   uint   `json:"project_id"`
	//TestCaseID  uint        `json:"test_case_id"`
	Status     string      `json:"status"`
	Priority   string      `json:"priority"`
	Severity   string      `json:"severity"`
	AssignedId uint        `json:"assigned_id"`
	Files      media.Files `gorm:"type:VARCHAR(255)"`
	Due        time.Time   `json:"due"`
}
