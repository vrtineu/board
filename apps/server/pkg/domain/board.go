package domain

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Name string `gorm:"type:text" json:"name"`
	Data []byte `gorm:"type:jsonb" json:"data"`
}
