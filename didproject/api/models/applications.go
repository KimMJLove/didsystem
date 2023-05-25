package models

type Application struct {
    Name      string       `json:"name"`
    ID        string       `json:id`
    DID       string       `json:"did"`
    Type      string       `json:"type"`
	URLs	  string 	   `json:"URLs"`
    PublicKey []AppPublicKey  `json:"publicKey"`
    Group     []Group      `json:"group"`
}

type AppPublicKey struct {
    Type           string `json:"type"`
    PublicKeyHex string `json:"PublicKeyHex"`
}

type Group struct {
    ID     string `json:"id"`
	DID    string `json:issDID`
}
