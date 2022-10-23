package utils

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
