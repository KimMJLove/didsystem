package handlers

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"didproject/api/models"

)

func ContractUpdate(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req models.Contract
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // 连接数据库
        db, err := sql.Open("postgres", "user=did password=diddev dbname=didtest host=192.168.88.77 port=5432 sslmode=disable")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer db.Close()

        // 更新文章
        _, err = db.Exec(`
            UPDATE contracts
            SET title = $1, type = $2, content = $3, author = $4, description = $5
            WHERE id = $6
        `, req.Title, req.Type, req.Content, req.Author, req.Description, req.ID)

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Contract updated successfully"})
    }
}

