package models

import (
	"time"
)

type Auction struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	ItemID       uint32    `json:"item_id" binding:"required"`
	InitialPrice float64   `json:"initial_price" binding:"required"`
	FinalPrice   float64   `json:"final_price" binding:"required"`
	Status       string    `gorm:"default:Pending" json:"status" binding:"required"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	EndAt        time.Time `json:"end_at" binding:"required"`
}
