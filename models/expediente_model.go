package models

import "gorm.io/gorm"

// Define the Expediente model
type Expediente struct {
	gorm.Model
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
}
