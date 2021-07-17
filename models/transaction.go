package models

import "time"

type Transaction struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	ItemID    uint32    `json:"item_id" binding:"required"`
	TxHash    string    `gorm:"size:255;not null" json:"tx_hash" binding:"required"`
	Buyer     string    `gorm:"size:255;not null" json:"buyer" binding:"required"`
	Seller    string    `gorm:"size:255;not null" json:"seller" binding:"required"`
	Price     uint64    `json:"price" binding:"required"`
	Status    string    `gorm:"size:255;not null;default:Pending" json:"status" binding:"required"`
	Fee       float64   `gorm:"size:255;not null" json:"fee" binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}