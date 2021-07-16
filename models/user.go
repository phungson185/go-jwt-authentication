package models

import "time"

type User struct {
	Id          uint32    `json:"id" gorm:"primary_key;auto_increment"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone" gorm:"size:255;not null"`
	UserAddress string    `json:"userAddress" gorm:"size:255;not null"`
	Status      bool      `json:"status"`
	VerifyCode  string    `json:"verifyCode"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}


