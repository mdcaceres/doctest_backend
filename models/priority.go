package models

import "gorm.io/gorm"

const (
	Critical = "CRITICAL"
	High     = "HIGH"
	Medium   = "MEDIUM"
	Low      = "LOW"
)

var (
	priorities = map[string]Priority{
		"1": Priority{Value: Critical},
		"2": Priority{Value: High},
		"3": Priority{Value: Medium},
		"4": Priority{Value: Low},
	}

	AvailablePriorities = []string{Critical, High, Medium, Low}
)

type Priority struct {
	gorm.Model
	Value string `json:"value" gorm:"unique"`
}

func getPriority(priority string) Priority {
	return priorities[priority]
}
