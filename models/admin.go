package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	// gorm.Model
	ID        uint           `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name      string         `gorm:"not_null" json:"name"`
	Email     string         `gorm:"not_null" json:"email"`
	Password  string         `gorm:"not_null" json:"password"`
	CreatedAt time.Time      `gorm:"not_null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not_null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
