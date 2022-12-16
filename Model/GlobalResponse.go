package Model

import "mini-wallet-exercise/Constant/APIResponse"

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewErrorResponse(error interface{}) *Response {
	return &Response{Status: APIResponse.FailResponse, Data: map[string]interface{}{"error": error}}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{Status: APIResponse.SuccessResponse, Data: data}
}
