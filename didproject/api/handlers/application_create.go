package handlers

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	"didproject/api/models"
)

func appgenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}

func ApplicationCreate(c *gin.Context) {
	// Generate random DID ID and Ethereum key pair
	id, err := appgenerateRandomBytes(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate random bytes"})
		return
	}

	// Create Ethereum key pair
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Ethereum key pair"})
		return
	}

	// Convert private key to hex string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast Ethereum public key to ECDSA"})
		return
	}

	// Store DID ID and public key in PostgreSQL database
	pool, err := pgxpool.Connect(context.Background(), "postgres://did:diddev@192.168.88.77:5432/didtest")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to PostgreSQL database"})
		return
	}
	defer pool.Close()

	var publicKeyID int
	err = pool.QueryRow(context.Background(), "INSERT INTO public_keys (public_key) VALUES ($1) RETURNING id", hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))).Scan(&publicKeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert public key into database"})
		return
	}

	var didID int
	err = pool.QueryRow(context.Background(), "INSERT INTO dids (did_id, public_key_id, status, auth) VALUES ($1, $2, 'creating','app') RETURNING id", fmt.Sprintf("did:app:%s", hex.EncodeToString(id)), publicKeyID).Scan(&didID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert DID into database"})
		return
	}

	// Generate DID JSON
	publicKeyDID := models.PublicKey{
		ID:         fmt.Sprintf("did:app:%s#keys-1", hex.EncodeToString(id)),
		Type:       "EcdsaSecp256k1VerificationKey2019",
		Controller: fmt.Sprintf("did:app:%s", hex.EncodeToString(id)),
		PublicKey:  hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA)),
	}

	did := models.DID{
		ID:      fmt.Sprintf("did:app:%s", hex.EncodeToString(id)),
		PrivKey: []string{privateKeyHex, hex.EncodeToString(privateKey.D.Bytes())},
		Status:  "creating",
		Auth:	  "app",
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"did": did, "publicKey": publicKeyDID}})
}