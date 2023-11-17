package models

type TechnicalFile struct {
	Id        string `json:"id"`
	Client    Client `json:"client"`
	ClientRef string `json:"clientRef"`
}
