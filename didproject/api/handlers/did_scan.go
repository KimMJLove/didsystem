package handlers

import (
    "database/sql"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"

    "didproject/api/models"
)

func DIDScan(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取查询参数
        didID := c.Query("did_id")

        // 查询 dids 表
        row := db.QueryRow("SELECT public_key_id, created, updated FROM dids WHERE did_id=$1", didID)
        // 将查询结果填充到 diddoc 结构体中
        var publicKeyid string
        var createdAt, updatedAt time.Time

        err := row.Scan(&publicKeyid, &createdAt, &updatedAt)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "failed to query did"})
            return
        }

		rowp := db.QueryRow("SELECT public_key FROM public_keys WHERE id=$1", publicKeyid)
		var publicKey string
        errs := rowp.Scan(&publicKey)
        if errs != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "failed to query did"})
            return
        }

        diddoc := models.DIDdoc{
            Context: "https://w3id.org/did/v1",
            ID:      didID,
            Version: "1",
            Created: createdAt,
            PublicKey: models.DocPublicKey{
                    ID:           didID + "#keys-1",
                    Type:         "Secp256k1",
                    PublicKeyHex: publicKey,
            },
            Authentication: didID + "#key-1",
            Recovery: didID + "#key-2",
            Service: models.Service{
                    ID:              didID + "#resolver",
                    Type:            "DIDResolve",
                    ServiceEndpoint: "https://did.baidu.com",
            },
            Proof: models.Proof{
                Type:           "Secp256k1",
                IssDID:         "did:iss:4455718552da01afccf74dadd86d475f",
            },
        }

        c.JSON(http.StatusOK, gin.H{"data": gin.H{"diddoc": diddoc}})
    }
}
