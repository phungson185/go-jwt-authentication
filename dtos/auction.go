package dtos

type CreateAuction struct {
	InitialPrice float64 `json:"initial_price,string" binding:"required"`
	FinalPrice   float64 `json:"final_price,string" binding:"required"`
	EndAt        int64   `json:"end_at,string" binding:"required"`
}

type UpdateAuction struct {
	InitialPrice float64 `json:"initial_price,string" binding:"required"`
	FinalPrice   float64 `json:"final_price,string" binding:"required"`
	EndAt        int64   `json:"end_at,string" binding:"required"`
}
