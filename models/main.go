package models

type LimitOffset struct {
	Limit  int `json:"limit" binding:"required,gte=0,lte=100"`
	Offset int `json:"offset" binding:"gte=0"`
}
