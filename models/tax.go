package models

import (
	"gorm.io/gorm"
)

type Tax struct {
	gorm.Model
	HotPrice     float32
	ColdPrice    float32
	EnergyPrice  float32
	DrenagePrice float32
	Email        string
}

type CreateTaxInput struct {
	HotPrice     float32 `json:"HotPrice"`
	ColdPrice    float32 `json:"ColdPrice"`
	EnergyPrice  float32 `json:"EnergyPrice"`
	DrenagePrice float32 `json:"DrenagePrice"`
}

type UpdateTaxInput struct {
	ID           uint    `json:"-"`
	HotPrice     float32 `json:"hot_price"`
	ColdPrice    float32 `json:"cold_price"`
	EnergyPrice  float32 `json:"energy_price"`
	DrenagePrice float32 `json:"drenage_price"`
}
