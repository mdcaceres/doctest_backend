package TestExecution

type ExecutionStep struct {
	ID              uint `json:"id"`
	TestExecutionID uint
	Order           uint   `json:"order"`
	Description     string `json:"description"`
	Expected        string `json:"expected"`
	CaseID          uint
	Status          string `json:"status"`
	Comment         string `json:"comment"`
}
