package model

import (
	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Products   []Product `json:"products"`
	Status     status    `json:"status"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}

type status int
