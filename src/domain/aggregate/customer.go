package aggregate

import (
	"errors"

	"github.com/aceino/belajar-ddd/src"
	"github.com/aceino/belajar-ddd/src/domain/entity"
	"github.com/google/uuid"
)

var (
	ErrInvalidCustomer = errors.New("a customer has to valid person")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []*src.Transaction
}

func NewCustomer(name string, email string) (Customer, error) {
	if name == "" && email == "" {
		return Customer{}, ErrInvalidCustomer
	}

	person := &entity.Person{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*src.Transaction, 0),
	}, nil
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
