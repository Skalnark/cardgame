package dto

import "encoding/json"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewError(message string, code int) []byte {
	err := Error{Message: message, Code: code}
	response, _ := json.Marshal(err)
	return response
}
