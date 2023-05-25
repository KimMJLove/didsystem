package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"didproject/api/models"
)

func InformationSelect(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")

		var user models.Information
		err := db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Gender,
			&user.Email,
			&user.Age,
			&user.Created,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
