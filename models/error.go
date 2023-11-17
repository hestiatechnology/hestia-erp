package models

type ErrorMessage struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"` // User friendly message
}
