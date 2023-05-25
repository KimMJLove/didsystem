package handlers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"

    "didproject/api/models"
)

func ApplicationAdd(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Parse request body to get name, URLs, and issdid
        var reqBody struct {
            Name  string `json:"name"`
            URLs  string `json:"URLs"`
            Type string  `json:Type`
            IssDID string `json:"issdid"`
        }
        if err := c.ShouldBindJSON(&reqBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // 获取查询参数group的did和公钥
        appdid := c.Query("did_id")
        // 查询 dids 表
        row := db.QueryRow("SELECT public_key_id FROM dids WHERE did_id=$1 and auth = 'app' ", appdid)
        // 将查询结果填充到 diddoc 结构体中
        var publicKeyid string
        err := row.Scan(&publicKeyid)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "failed to query did"})
            return
        }
        rowp := db.QueryRow("SELECT public_key FROM public_keys WHERE id=$1", publicKeyid)
        var apppublicKey string
        errs := rowp.Scan(&apppublicKey)
        if errs != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "failed to query did"})
            return
        }

        // issdid := c.Query("did_id")
        issdid := reqBody.IssDID
        type1 := reqBody.Type
        // var uurlls string
        // var name string
        application := models.Application{
            Name:       reqBody.Name,
            DID:        appdid,
            Type:       type1,
            URLs:       reqBody.URLs,
            PublicKey: []models.AppPublicKey{
                {
                    Type:         "Secp256k1",
                    PublicKeyHex:  apppublicKey,
                },
            },
            Group: []models.Group{
                {
                    ID:  "1",
                    DID: issdid,
                },
            },
        }

        // Insert application into applications table
        stmt, err := db.Prepare("INSERT INTO applications(name, appdid, urls, type, issdid) VALUES($1, $2, $3, $4, $5) RETURNING id")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer stmt.Close()

        var appID int
        err = stmt.QueryRow(application.Name, application.DID, application.URLs, application.Type, application.Group[0].DID).Scan(&appID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"data": gin.H{"application": application}})
    }
}
