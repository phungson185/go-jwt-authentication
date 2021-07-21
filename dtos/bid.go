package dtos

type CreateBid struct {
	Price float64 `json:"price,string" binding:"required"`
}
