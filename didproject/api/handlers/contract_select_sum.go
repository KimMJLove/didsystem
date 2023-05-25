package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)



func ContractSum(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM contracts").Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve contract count"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}
