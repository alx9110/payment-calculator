package models

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Name  string
	Value float32
	Cost  float32
}

type CreateRecordInput struct {
	Name  string  `json:"name" binding:"required"`
	Value float32 `json:"value" binding:"required"`
}

type UpdateRecordInput struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}
