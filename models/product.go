package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type CreateProductInput struct {
	Code  string `json:"code" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}

type UpdateProductInput struct {
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
