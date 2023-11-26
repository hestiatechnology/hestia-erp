package models

type LimitOffset struct {
	Limit  int `json:"limit" binding:"required"`
	Offset int `json:"offset" binding:"required"`
}
