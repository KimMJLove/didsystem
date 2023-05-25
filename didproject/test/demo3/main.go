package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// 构造数据库连接字符串
	dbUri := "postgres://did:diddev@192.168.88.77:5432/didtest?sslmode=disable"

	// 连接数据库
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully")
}
