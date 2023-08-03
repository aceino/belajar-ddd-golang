package aggregate

import (
	"errors"

	"github.com/aceino/belajar-ddd/src/domain/entity"
	"github.com/google/uuid"
)

var (
	ErrMisingValue = errors.New("missig value")
)

type Product struct {
	item     *entity.Item
	payment  *entity.Payment
	year     int
	color    string
	quantity int
}

func NewProduct(name string, year int, color string, price float64) (Product, error) {
	if name == " " {
		return Product{}, ErrMisingValue
	}

	return Product{
		item: &entity.Item{
			ID:    uuid.New(),
			Name:  name,
			Price: price,
		},
		year:     year,
		color:    color,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetYear() int {
	return p.year
}

func (p Product) GetColor() string {
	return p.color
}

func (p Product) GetPrice() float64 {
	return p.item.Price
}
