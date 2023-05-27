package models

import (
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"github.com/mdcaceres/doctest/models/media"
	"time"
)

type Case struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	UserID      uint
	Type        string
	Priority    string
	Description string
	Duration    time.Duration
	Steps       []Step
	ProjectID   uint
	SuiteID     uint
	Status      string
	Files       media.Files `gorm:"type:VARCHAR(255)"`
	Execution   []TestExecution.TestExecution
	Comments    []TestComment
}
