package models

type Issuer struct {
    ID               int64  `json:"id"`
    UUID             string `json:"uuid"`
    DID              string `json:"did"`
    Website          string `json:"website"`
    Endpoint         string `json:"endpoint"`
    ShortDescription string `json:"short_description"`
    LongDescription  string `json:"long_description"`
    ServiceType      string `json:"service_type"`
    RequestData      string `json:"request_data"`
    Deleted          bool   `json:"deleted"`
    CreatedAt        string `json:"created_at"`
    UpdatedAt        string `json:"updated_at"`
}
