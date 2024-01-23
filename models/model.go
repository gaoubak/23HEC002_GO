package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GetModels() []interface{} {
	return []interface{}{
		&User{},
	}
}
