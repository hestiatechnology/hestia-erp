package models

type Client struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	VatId      string `json:"vatId"`
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
}
