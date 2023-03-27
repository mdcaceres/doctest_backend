package domains

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Users       []*User     `json:"users,omitempty" gorm:"many2many:user_projects;"`
	TestSuit    []*TestSuit `json:"test_suits" gorm:"many2many:testsuit_projects"`
}
