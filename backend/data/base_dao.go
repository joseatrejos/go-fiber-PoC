package daos

import "gorm.io/gorm"

// BaseDAO is a generic DAO that provides CRUD operations
type BaseDAO[T any] struct {
	db         *gorm.DB
	modelClass T
}

// Get retrieves a record by ID
func (dao *BaseDAO[T]) Get(id uint) (*T, error) {
	var obj T
	err := dao.db.First(&obj, id).Error
	return &obj, err
}

// GetAll retrieves all records
func (dao *BaseDAO[T]) GetAll() ([]T, error) {
	var objs []T
	err := dao.db.Find(&objs).Error
	return objs, err
}

// Create adds a new record to the database
func (dao *BaseDAO[T]) Create(obj *T) error {
	return dao.db.Create(&obj).Error
}

// Update updates an existing record
func (dao *BaseDAO[T]) Update(obj **T) error {
	return dao.db.Save(&obj).Error
}

// Delete removes a record from the database
func (dao *BaseDAO[T]) Delete(id uint) error {
	return dao.db.Delete(&dao.modelClass, id).Error
}
