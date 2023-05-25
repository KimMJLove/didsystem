package mydb

import (
	"log"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// 初始化 Redis 连接
func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.88.77:6379",
		Password: "diddev",
		DB:       0,
	})

	// 测试连接
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// 关闭 Redis 连接
func CloseRedis() {
	err := redisClient.Close()
	if err != nil {
		log.Printf("Failed to close Redis connection: %v", err)
	}
}

// 在 db 包中提供使用 Redis 的函数
func SetData(key string, value string) error {
	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetData(key string) (string, error) {
	value, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return value, nil
}

// 在其他函数中调用 Redis 函数
func SomeFunction() {
	err := SetData("key", "value")
	if err != nil {
		log.Printf("Failed to set data in Redis: %v", err)
	}

	value, err := GetData("key")
	if err != nil {
		log.Printf("Failed to get data from Redis: %v", err)
	}
	log.Println(value)
}
