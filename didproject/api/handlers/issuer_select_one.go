package handlers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"

	"didproject/api/models"
)

func IssuerSelectOne(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get issuer ID from request URL
        issuerID := c.Param("id")

        // Query issuer from issuers table
        row := db.QueryRow("SELECT id, uuid, did, website, endpoint, short_description, long_description, service_type, request_data, deleted, created_at, updated_at FROM issuers WHERE id=$1", issuerID)
        var issuer models.Issuer
        err := row.Scan(&issuer.ID, &issuer.UUID, &issuer.DID, &issuer.Website, &issuer.Endpoint, &issuer.ShortDescription, &issuer.LongDescription, &issuer.ServiceType, &issuer.RequestData, &issuer.Deleted, &issuer.CreatedAt, &issuer.UpdatedAt)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"data": gin.H{"issuer": issuer}})
    }
}