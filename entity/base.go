package entity

type BaseResponse struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors any         `json:"error,omitempty"`
}
