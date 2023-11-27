package models

type Client struct {
	Id         string `json:"id" binding:"uuid4_rfc4122"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	VatId      string `json:"vatId"`
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	Locality   string `json:"locality"`
	Country    string `json:"country" binding:"iso3166_1_alpha2"`
}
