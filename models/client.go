package models

import (
	"github.com/google/uuid"
)

type Client struct {
	Id         uuid.UUID `json:"id" binding:"uuid4_rfc4122"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	VatId      string    `json:"vatId"`
	Street     string    `json:"street"`
	PostalCode string    `json:"postalCode"`
	Locality   string    `json:"locality"`
	Country    string    `json:"country" binding:"iso3166_1_alpha2"`
}

type NewClient struct {
	Name       string `json:"name" binding:"required"`
	Code       string `json:"code" binding:"required"`
	VatId      string `json:"vatId" binding:"required"`
	Street     string `json:"street" binding:"required"`
	PostalCode string `json:"postalCode" binding:"required"`
	Locality   string `json:"locality" binding:"required"`
	Country    string `json:"country" binding:"iso3166_1_alpha2,required"`
}
