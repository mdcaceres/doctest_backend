package domains

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	TestSteps  []*TestStep
	TestResult TestResult
}
