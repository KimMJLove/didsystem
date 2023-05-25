package handlers

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"didproject/api/models"

)

func ContractCreate(db *sql.DB) gin.HandlerFunc {
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

		// 插入新的文章
		var id int64
		err = db.QueryRow(`
			INSERT INTO contracts (title, type, content, author, description)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`, req.Title, req.Type, req.Content, req.Author, req.Description).Scan(&id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}