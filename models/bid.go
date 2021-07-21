package models

import "time"

type Bid struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	AuctionID uint32    `json:"auction_id" binding:"required"`
	TxHash    string    `gorm:"size:255;not null" json:"tx_hash" binding:"required"`
	Bidder    string    `gorm:"size:255;not null" json:"bidder" binding:"required"`
	Price     float64    `json:"price" binding:"required"`
	Fee       float64   `gorm:"size:255;not null" json:"fee" binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
