package mydb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var db *sql.DB

// InitDB 连接数据库并初始化全局数据库实例
func InitDB() {
	dbConfig := getDBConfig()
	connString := getConnectionString(dbConfig)
	conn, err := sql.Open(dbConfig.Type, connString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}

	fmt.Println("Connected to the database")

	db = conn
}

// GetDB 返回全局数据库实例
func GetDB() *sql.DB {
	return db
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}

// 获取数据库配置
func getDBConfig() DBConfig {
	var config DBConfig
	err := viper.UnmarshalKey("postgresql", &config)
	if err != nil {
		log.Fatalf("Error reading database config from config file: %s", err)
	}
	return config
}

// 获取连接字符串
func getConnectionString(config DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
}

// 数据库配置结构体
type DBConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
