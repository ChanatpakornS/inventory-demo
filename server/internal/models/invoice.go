package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	// GORM automatically add id, created_at, updated_at, deleted_at fields
	gorm.Model
	Name   string  `gorm:"size:200;not null" json:"name"`
	Status string  `gorm:"type:text" json:"status"`
	Method string  `gorm:"size:200;not null" json:"method"`
	Amount float64 `gorm:"not null" json:"amount"`
}
