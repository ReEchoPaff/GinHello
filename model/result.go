package model

type Result struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"请求体"`
	Data    interface{} `json:"data"`
}
