package utils

import (
	"database/sql"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"didproject/api/mydb"
)

type TableCount struct {
	DidsCount        int `json:"dids_count"`
	ContractsCount   int `json:"contracts_count"`
	ApplicationsCount int `json:"applications_count"`
	IssuersCount     int `json:"issuers_count"`
}

// ShowAccount godoc
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [get]
func GetTableCounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var counts TableCount

		// Get count for dids table
		row := db.QueryRow("SELECT COUNT(*) FROM dids")
		err := row.Scan(&counts.DidsCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Get count for contracts table
		row = db.QueryRow("SELECT COUNT(*) FROM contracts")
		err = row.Scan(&counts.ContractsCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Get count for applications table
		row = db.QueryRow("SELECT COUNT(*) FROM applications")
		err = row.Scan(&counts.ApplicationsCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Get count for issuers table
		row = db.QueryRow("SELECT COUNT(*) FROM issuers")
		err = row.Scan(&counts.IssuersCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 将 counts 转换为字符串
		countsJSON, err := json.Marshal(counts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to marshal counts to JSON",
			})
			return
		}

		// 将 counts 保存到 Redis
		err = mydb.SetData("table_counts", string(countsJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save table counts in Redis",
			})
			return
		}

		c.JSON(http.StatusOK, counts)
	}
}