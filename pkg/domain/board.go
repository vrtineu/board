package domain

import "github.com/google/uuid"

type Board struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:text" json:"name"`
	Data      []byte    `gorm:"type:jsonb" json:"data"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime" json:"updated_at"`
}
