package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primary_key;auto_increment"`
	Name      string     `json:"name" gorm:"type:varchar(150);not null""`
	LastName  string     `json:"last_name" gorm:"type:varchar(150);not null""`
	Roles     []Role     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:user_roles;"`
	Photo     *string    `json:"photo_url"`
	Email     string     `json:"email" gorm:"unique"`
	Password  []byte     `json:"-"`
	Projects  []*Project `json:"projects" gorm:"many2many:user_projects;"`
	CreatedAt *time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
