package dtos

type CreateUser struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
}

type VerifyEmail struct {
	Email      string `json:"email" binding:"required"`
	VerifyCode string `json:"verify_code" binding:"required"`
}

type Login struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
