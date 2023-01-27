package models

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	HotValue     float32
	HotCost      float32
	ColdValue    float32
	ColdCost     float32
	EnergyValue  float32
	EnergyCost   float32
	DrenageValue float32
	DrenageCost  float32
	Email        string
}

type CreateRecordInput struct {
	HotValue    float32 `json:"HotValue"`
	ColdValue   float32 `json:"ColdValue"`
	EnergyValue float32 `json:"EnergyValue"`
}

type UpdateRecordInput struct {
	ID           uint    `json:"-"`
	HotValue     float32 `json:"hot_value"`
	HotCost      float32 `json:"hot_cost"`
	ColdValue    float32 `json:"cold_value"`
	ColdCost     float32 `json:"cold_cost"`
	EnergyValue  float32 `json:"energy_value"`
	EnergyCost   float32 `json:"energy_cost"`
	DrenageValue float32 `json:"drenage_value"`
	DrenageCost  float32 `json:"drenage_cost"`
}
