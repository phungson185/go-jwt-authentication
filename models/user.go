package models

import "time"

type User struct {
	Id          uint32    `json:"id" gorm:"primary_key;auto_increment"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone" gorm:"size:255;not null"`
	UserAddress string    `json:"user_address" gorm:"size:255;not null"`
	Status      bool      `json:"status"`
	VerifyCode  string    `json:"verify_code"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}


