package entity

import "github.com/google/uuid"

type Person struct {
	ID    uuid.UUID
	Name  string
	Email string
	Age   int
}
