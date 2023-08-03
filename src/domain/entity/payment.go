package entity

import "github.com/google/uuid"

type Payment struct {
	ID   uuid.UUID
	Name string
	Card string
	Cash float64
}
