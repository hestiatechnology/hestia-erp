package models

type Company struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	VatId      string `json:"vat_id"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
	Timezone   string `json:"timezone"`
}
