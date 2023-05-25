package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

type Contract struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ContractSelect(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取查询参数
        limit := c.DefaultQuery("limit", "10")
        offset := c.DefaultQuery("offset", "0")

        // 查询合同列表
        rows, err := db.Query(`SELECT * FROM contracts ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        // 构造合同列表
		var contracts []Contract
		for rows.Next() {
			var contract Contract
			err := rows.Scan(&contract.ID, &contract.Title, &contract.Type, &contract.Content, &contract.Author, &contract.Description, &contract.CreatedAt, &contract.UpdatedAt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to retrieve data from the database",
				})
				return
			}
			contracts = append(contracts, contract)
		}

        c.JSON(http.StatusOK, contracts)
    }
}