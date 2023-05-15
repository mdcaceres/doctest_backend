package media

import "gorm.io/gorm"

type CaseFile struct {
	gorm.Model
	CaseId uint   `json:"caseId"`
	URL    string `json:"url"`
	file   []byte `json:"file"`
}
