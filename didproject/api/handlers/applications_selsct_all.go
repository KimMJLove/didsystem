package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Application struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AppDID    string `json:"appdid"`
	Type      string `json:"type"`
	URLs      string `json:"urls"`
	IssDID    string `json:"issdid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetApplications(c *gin.Context) {
	// 获取查询参数
	limit := c.DefaultQuery("limit", "12")
	offset := c.DefaultQuery("offset", "0")

	// Connect to the database
	db, err := sql.Open("postgres", "postgres://did:diddev@192.168.88.77:5432/didtest?sslmode=disable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// 查询应用列表
	rows, err := db.Query("SELECT * FROM applications ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// 将行转换为Application结构的切片
	applications := make([]Application, 0)
	for rows.Next() {
		var app Application
		err := rows.Scan(&app.ID, &app.Name, &app.AppDID, &app.Type, &app.URLs, &app.IssDID, &app.CreatedAt, &app.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		applications = append(applications, app)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回应用列表作为JSON响应
	c.JSON(http.StatusOK, applications)
}
