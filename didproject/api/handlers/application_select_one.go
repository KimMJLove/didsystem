package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func ApplicationSelectOne(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        row := db.QueryRow("SELECT urls FROM applications WHERE id=$1", id)
        var urls string
        err = row.Scan(&urls)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusNotFound, gin.H{"error": "Urls not found"})
                return
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{"urls": urls})
    }
}