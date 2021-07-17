package dtos

type Res struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(status bool, message string, data interface{}) Res {
	res := Res{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}
