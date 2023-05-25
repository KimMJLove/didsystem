package models

type Identity struct {
    ID       int    `json:"id"`
    Hubname string `json:"hubname"`
    Description string `json:"description"`
    HubPublicKey []HubPublicKey `json:"hubpublicKey"`
}

type HubPublicKey struct {
    Type           string `json:"type"`                       //"type":"Secp256k1VerificationKey2018"
    Controller     string `json:"controller"`                 //"controller":"group"\
    PublicKeyBase58 string `json:"publicKeyBase58"`           //"publicKeyBase58": "5u5SQEoKYzphH4A4j4tMvM2zE4tVijy1aSFRexjvqPky"
}