package models

type PublicKey struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Controller string `json:"controller"`
	PublicKey  string `json:"publicKey"`
}

type TableCount struct {
	DidsCount        int `json:"dids_count"`
	ContractsCount   int `json:"contracts_count"`
	ApplicationsCount int `json:"applications_count"`
	IssuersCount     int `json:"issuers_count"`
}