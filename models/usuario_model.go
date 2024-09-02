package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the database model
type User struct {
	gorm.Model
	Name      string    `gorm:"size:255;not null" json:"name" form:"name" validate:"required"`
	Email     string    `gorm:"unique;size:255" json:"email" form:"email" validate:"required,email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
