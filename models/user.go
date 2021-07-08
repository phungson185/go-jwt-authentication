package models

import "time"

type User struct {
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	UserAddress string    `json:"userAddress"`
	Status      bool      `json:"status"`
	VerifyCode  string    `json:"verifyCode"`
	CreatedAt   time.Time `json:"createdAt" time_format:"2006-01-02" time_utc:"7"`
	UpdatedAt   time.Time `json:"updatedAt" time_format:"2006-01-02" time_utc:"7"`
}
