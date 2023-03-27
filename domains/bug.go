package domains

import "gorm.io/gorm"

type Bug struct {
	gorm.Model
	TestSuitID uint
	TestSuit   TestSuit
}
