package container

import (
	"lib_management/repository"
)

type Container[T any] interface {
	GetRepository() repository.Repository[T]
}

type container[T any] struct {
	repo repository.Repository[T]
}

func GetContainer[T any](repo repository.Repository[T]) Container[T] {
	return &container[T]{
		repo: repo,
	}
}

func (c *container[T]) GetRepository() repository.Repository[T] {
	return c.repo
}
