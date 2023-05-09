package media

import "gorm.io/gorm"

type Img struct {
	gorm.Model
	ownerId uint   `json:"ownerId"`
	URL     string `json:"url"`
	file    []byte `json:"file"`
}
