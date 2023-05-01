package models

type Invitation struct {
	ID           uint `gorm:"primary_key"`
	InviterID    uint
	InvitedID    uint
	InvitedEmail string
	Role         string
	ProjectID    uint
	Accepted     bool `gorm:"default:false"`
}
