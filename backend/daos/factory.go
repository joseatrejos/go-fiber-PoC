package daos

import (
	"gorm.io/gorm"
)

// NewBaseDAO creates a new instance of BaseDAO for a specific model
func NewBaseDAO[T any](db *gorm.DB, modelClass T) *BaseDAO[T] {
	return &BaseDAO[T]{
		db:         db,
		modelClass: modelClass,
	}
}
