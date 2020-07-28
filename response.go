package helper

import (
	"encoding/json"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ApiReturn(code int, msg string, data interface{}) string {
	var ar = &ApiResponse{}
	ar.Code = code
	ar.Message = msg
	ar.Data = data
	arj, _ := json.Marshal(ar)
	return string(arj)
}

func EchoReturn(code int, msg string, data interface{}) *ApiResponse {
	var ar = &ApiResponse{}
	ar.Code = code
	ar.Message = msg
	ar.Data = data
	return ar
}