package models

import (
    "time"
)

type DIDdoc struct {
    Context        string `json:"@context"`
    ID             string `json:"id"`
    Version        string `json:"version"`
    Created        time.Time `json:"created"`
    PublicKey      DocPublicKey `json:"publicKey"`
    Authentication string `json:"authentication"`
    Recovery       string `json:"recovery"`
    Service        Service `json:"service"`
    Proof          Proof `json:"proof"`
}

type DocPublicKey struct {
    ID           string `json:"id"`
    Type         string `json:"type"`
    PublicKeyHex string `json:"publicKeyHex"`
}

type Service struct {
    ID              string `json:"id"`
    Type            string `json:"type"`
    ServiceEndpoint string `json:"serviceEndpoint"`
}

type Proof struct {
    Type           string `json:"type"`
    IssDID         string `json:"issdid`
}
