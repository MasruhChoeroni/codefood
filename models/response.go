package models

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
}
