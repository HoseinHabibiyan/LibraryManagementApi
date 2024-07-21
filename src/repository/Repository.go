package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (repository Repository[T]) Create(entity *T) error {
	return repository.db.Create(&entity).Error
}

func (repository Repository[T]) Update(entity *T) error {
	return repository.db.Updates(&entity).Error
}

func (repository Repository[T]) Delete(id int32) error {
	return repository.db.Delete("id = ?", id).Error
}

func (repository Repository[T]) GetById(id int32) *T {
	var entity T
	err := repository.db.First(&entity, "id = ?", id).Error

	if err != nil {
		return nil
	}

	return &entity
}

func (repository Repository[T]) Get() *T {
	var entities T
	err := repository.db.Find(&entities).Error

	if err != nil {
		return nil
	}

	return &entities
}
