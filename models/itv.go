package models

import "time"

type TechnicalFile struct {
	Id          string                   `json:"id"`
	Code        string                   `json:"code"`
	Description string                   `json:"description"`
	Client      Client                   `json:"client"`
	ClientRef   string                   `json:"clientRef"`
	Family      TechnicalFileFamily      `json:"family"`
	Type        TechnicalFileType        `json:"type"`
	RawMaterial TechnicalFileRawMaterial `json:"rawMaterial"`
	Date        time.Time                `json:"date"`
}

// Theres a new type for some properties
// for data consistency with the DB, this
// allows the frontend to link the current
// value with the list of possible values

type TechnicalFileFamily struct {
	// Family of the technical file
	// T-shirt, polo, sweatshirt, etc.
	Id     string `json:"id"`
	Family string `json:"family"`
}

type TechnicalFileType struct {
	// Type of the technical file
	// male, female, unisex, child, baby, etc.
	Id   string `json:"id"`
	Type string `json:"type"`
}

type TechnicalFileRawMaterial struct {
	// Raw material of the technical file
	// cotton, polyester, etc.
	Id          string `json:"id"`
	RawMaterial string `json:"rawMaterial"`
}
