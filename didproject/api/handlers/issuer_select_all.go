package handlers

import (
    "database/sql"
    "net/http"
    "encoding/json"
    
    "github.com/gin-gonic/gin"
)

type Issuer struct {
    ID                int           `json:"id"`
    UUID              string        `json:"uuid"`
    DID               string        `json:"did"`
    Website           string        `json:"website"`
    Endpoint          string        `json:"endpoint"`
    ShortDescription  string        `json:"short_description"`
    LongDescription   string        `json:"long_description"`
    ServiceType       string        `json:"service_type"`
    RequestData       RequestData   `json:"request_data"`
    Deleted           bool          `json:"deleted"`
    CreatedAt         string        `json:"created_at"`
    UpdatedAt         string        `json:"updated_at"`
}

type RequestData struct {
    Types []string `json:"types"`
}

func IssuerSelectAll(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        rows, err := db.Query("SELECT id, uuid, did, website, endpoint, short_description, long_description, service_type, request_data, deleted, created_at, updated_at FROM issuers")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        issuers := []Issuer{}
        for rows.Next() {
            var issuer Issuer
            var requestData []byte
            err := rows.Scan(
                &issuer.ID,
                &issuer.UUID,
                &issuer.DID,
                &issuer.Website,
                &issuer.Endpoint,
                &issuer.ShortDescription,
                &issuer.LongDescription,
                &issuer.ServiceType,
                &requestData,
                &issuer.Deleted,
                &issuer.CreatedAt,
                &issuer.UpdatedAt,
            )
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            err = json.Unmarshal(requestData, &issuer.RequestData)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            issuers = append(issuers, issuer)
        }

        c.JSON(http.StatusOK, gin.H{"issuers": issuers})
    }
}
