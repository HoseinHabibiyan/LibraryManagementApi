package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(id int32) error
	GetById(id int32) *T
	GetAll() *[]T
}

type repository[T any] struct {
	db *gorm.DB
	Repository[T]
}

func GetRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{
		db: db,
	}
}

func (repository repository[T]) Create(entity *T) error {
	return repository.db.Create(&entity).Error
}

func (repository repository[T]) Update(entity *T) error {
	return repository.db.Updates(&entity).Error
}

func (repository repository[T]) Delete(id int32) error {
	return repository.db.Delete("id = ?", id).Error
}

func (repository repository[T]) GetById(id int32) *T {
	var entity T
	err := repository.db.First(&entity, "id = ?", id).Error

	if err != nil {
		return nil
	}

	return &entity
}

func (repository repository[T]) GetAll() *[]T {
	var entities []T
	err := repository.db.Find(&entities).Error

	if err != nil {
		return nil
	}

	return &entities
}
