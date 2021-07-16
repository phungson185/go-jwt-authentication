package models

import "time"

type Item struct {
	ID          uint32    `json:"id" gorm:"primary_key;auto_increment" `
	Name        string    `json:"name" gorm:"size:255;not null" `
	Description string    `json:"description" gorm:"size:255;not null" `
	Price       int64     `json:"price"`
	Currency    string    `json:"currency" gorm:"size:255;not null" `
	Owner       string    `json:"owner" gorm:"size:255;not null" `
	Creator     string    `json:"creator" gorm:"size:255;not null" `
	Metadata    string    `json:"metadata" gorm:"size:255;not null" `
	Status      string    `json:"status" gorm:"size:255;not null;default:Pending" `
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" `
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP" `
}