package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func ContractSelectOne(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        row := db.QueryRow("SELECT content FROM contracts WHERE id=$1", id)
        var content string
        err = row.Scan(&content)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
                return
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{"content": content})
    }
}