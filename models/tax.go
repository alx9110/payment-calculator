package models

import (
	"gorm.io/gorm"
)

type Tax struct {
	gorm.Model
	Name  string
	Value float32
}

type CreateTaxInput struct {
	Name  string  `json:"name" binding:"required"`
	Value float32 `json:"value" binding:"required"`
}

type UpdateTaxInput struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}
