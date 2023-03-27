package domains

import "gorm.io/gorm"

type TestSuit struct {
	gorm.Model
	Name        string
	Description string
	UserID      uint
	User        User
	Target      string
	TestCases   []*TestCase `json:"test_cases" gorm:"many2many:testcases_testsuit"`
	Bugs        []*Bug      `gorm:"foreignKey:TestSuitID"`
}
