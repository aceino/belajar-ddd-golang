package repository

import (
	"errors"

	"github.com/aceino/belajar-ddd/src/domain/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("Customer is not found")
	ErrFailedToAddCustomer = errors.New("Failed to add customer")
	ErrUpdateCustomer      = errors.New("Failed to update customer")
)

type RepositoryCustomer interface {
	GetAll(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
