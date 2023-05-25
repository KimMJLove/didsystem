package models

// DID represents a Decentralized Identifier
type DID struct {
	ID      string   `json:"did"`
	PrivKey []string `json:"privKey"`
	Status  string   `json:"status"`
	Auth	string   `json:"pwj"`
}

