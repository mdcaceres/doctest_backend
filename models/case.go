package models

import (
	"gorm.io/gorm"
	"time"
)

type Case struct {
	gorm.Model
	Title  string
	UserId uint
	//Priority    Priority `gorm:"foreignKey:PriorityID"`
	Description string
	Duration    time.Time
	Steps       []*Step
	ProjectId   uint
	SuiteId     uint
}
