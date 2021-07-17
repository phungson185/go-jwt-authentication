package dtos

type CreateItem struct {
	Name        string `json:"name" binding:"required`
	Description string `json:"description"`
	Price       int64  `json:"price,string" binding:"required`
	Currency    string `json:"currency" binding:"required`
	Owner       string `json:"owner" binding:"required`
	Creator     string `json:"creator" binding:"required`
	Metadata    string `json:"metadata" binding:"required`
}

type UpdateItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price,string"`
	Currency    string `json:"currency"`
}
