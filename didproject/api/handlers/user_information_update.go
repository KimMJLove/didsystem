package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"didproject/api/models"
)

// type User struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Gender   string `json:"gender"`
// 	Email    string `json:"email"`
// 	Age      int    `json:"age"`
// 	Created  string `json:"created"`
// }

func InformationUpdate(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取要更新的用户信息
		var updatedUser models.Information
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// 检查要更新的用户是否存在
		username := c.Query("username")
		var existingUser models.Information
		err := db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(
			&existingUser.ID,
			&existingUser.Name,
			&existingUser.Username,
			&existingUser.Password,
			&existingUser.Gender,
			&existingUser.Email,
			&existingUser.Age,
			&existingUser.Created,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
			return
		}

		// 更新用户信息
		_, err = db.Exec("UPDATE users SET name = $1, password = $2, gender = $3, email = $4, age = $5 WHERE username = $6",
			updatedUser.Name,
			updatedUser.Password,
			updatedUser.Gender,
			updatedUser.Email,
			updatedUser.Age,
			username,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user information"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User information updated successfully"})
	}
}
