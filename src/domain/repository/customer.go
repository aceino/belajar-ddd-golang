package repository

import (
	"errors"

	"github.com/aceino/belajar-ddd/src/domain/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("Product is not found")
	ErrFailedToAddProduct = errors.New("Failed to add Product")
	ErrUpdateProduct      = errors.New("Failed to update Product")
)

type RepositoryProduct interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
