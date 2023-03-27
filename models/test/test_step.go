package test

type Step struct {
	ID          uint   `json:"id"`
	Order       uint   `json:"order"`
	Description string `json:"description"`
	TestCaseID  uint
}
