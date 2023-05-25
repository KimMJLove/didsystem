package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func ApplicationDelete(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        result, err := db.Exec("DELETE FROM applications WHERE id=$1", id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        affectedRows, err := result.RowsAffected()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        if affectedRows == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "application not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "application deleted successfully"})
    }
}