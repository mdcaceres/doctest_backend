package models

type Invitation struct {
	ID           uint `gorm:"primary_key"`
	InviterID    uint
	InvitedID    uint
	InvitedEmail string
	ProjectID    uint
	Accepted     bool `gorm:"default:false"`
}
