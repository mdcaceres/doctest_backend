package models

import (
	"database/sql/driver"
	"errors"
	"strings"
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
	ProjectId   uint
	SuiteId     uint
	Status      string
	Files       Files `gorm:"type:VARCHAR(255)"`
}

type Files []string

func (f *Files) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*f = strings.Split(string(bytes), ",")
	return nil
}

func (f Files) Value() (driver.Value, error) {
	if len(f) == 0 {
		return nil, nil
	}
	return strings.Join(f, ","), nil
}
