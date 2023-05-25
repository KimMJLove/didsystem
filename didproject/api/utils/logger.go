package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/elastic/go-elasticsearch/v7"
)

func Logger() gin.HandlerFunc {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.88.77:9200"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 定义允许捕获的接口URL前缀
	allowedURLPrefixes := []string{
		"/system/register",
		"/system/login",
		"/issuer/select/:id",
		"/did/create",
		"/application/create",
		"/application/add",
		"/application/update/:id",
		"/application/delete/:id",
		"/application/select/:id",
	

		"/contract/delete/",
		"/contract/select/:id",
		"/contract/create",


	}

	return func(c *gin.Context) {
		// 检查请求的URL是否在允许捕获的列表中
		url := c.Request.URL.Path
		isAllowedURL := false
		for _, allowedURLPrefix := range allowedURLPrefixes {
			if strings.HasPrefix(url, allowedURLPrefix) {
				isAllowedURL = true
				break
			}
		}
		if !isAllowedURL {
			c.Next()
			return
		}

		start := time.Now()

		// 处理请求
		c.Next()

		// Log request
		duration := time.Since(start)
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		logData := map[string]interface{}{
			"timestamp":   time.Now(),
			"duration":    duration.Seconds(),
			"status":      c.Writer.Status(),
			"request":     map[string]interface{}{
				"method":  c.Request.Method,
				"url":     c.Request.URL.String(),
				"headers": c.Request.Header,
				"body":    string(reqBody),
				"ip":      c.ClientIP(),
			},
			"response":   c.Writer.Size(),
			"client_ip":  c.ClientIP(),
			"server_ip":  c.Request.Host,
			"server_port": cfg.Addresses[0],
		}
		indexName := fmt.Sprintf("log-%s", time.Now().Format("2006.01.02"))
		res, err := client.Indices.Exists([]string{indexName})
		if err != nil {
			log.Fatalf("Error checking if index %s exists: %s", indexName, err)
		}
		if res.StatusCode >= 400 && res.StatusCode <= 599 {
			// 处理错误
		}

		logDataBytes, err := json.Marshal(logData)
		if err != nil {
			log.Fatalf("Error marshalling logData: %s", err)
		}

		_, err = client.Index(
			indexName,
			strings.NewReader(string(logDataBytes)),
			client.Index.WithContext(context.Background()),
		)
		if err != nil {
			log.Fatalf("Error indexing document: %s", err)
		}
	}
}



const createIndexTemplate = `
{
    "mappings": {
        "_doc": {
            "dynamic": false,
            "properties": {
                "timestamp": {
                    "type": "date"
                },
                "duration": {
                    "type": "float"
                },
                "status": {
                    "type": "integer"
                },
                "request": {
                    "properties": {
                        "method": {
                            "type": "keyword"
                        },
                        "url": {
                            "type": "text"
                        },
                        "headers": {
                            "type": "object",
                            "enabled": false
                        },
                        "body": {
                            "type": "text"
                        },
                        "ip": {
                            "type": "text"
                        }
                    }
                },
                "response": {
                    "properties": {
                        "status": {
                            "type": "integer"
                        },
                        "headers": {
                            "type": "object",
                            "enabled": false
                        },
                        "body": {
                            "type": "text"
                        }
                    }
                },
                "client_ip": {
                    "type": "text"
                },
                "server_ip": {
                    "type": "text"
                }
            }
        }
    }
}`