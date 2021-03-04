package db

import (
	"database/sql"
	"fmt"
	"time"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectDB ...
func ConnectDB() {
	connect()
}

// DB ...
func DB() *sql.DB {
	return db
}

// Connect connect mysql database
func connect() {
	if db == nil {
		host := os.Getenv("MYSQL_HOST")
		username := os.Getenv("MSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		dbName := os.Getenv("MYSQL_DBNAME")
		uri := username + ":" + password + "@" + host + "/" + dbName
		fmt.Println(host)
		fmt.Printf("Connect to database: %v/%v\n", host, dbName)
		conn, err := sql.Open("mysql", uri)
		if err != nil {
			panic(err)
		}
		db = conn
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		fmt.Printf("Database connected.\n")
	}
}

// ExistTable return is table exited
func ExistTable(name string) bool {
	schema := os.Getenv("MYSQL_DBNAME")
	row, err := db.Query("SELECT COUNT(*) FROM information_schema.TABLES WHERE (TABLE_SCHEMA = '%v') AND (TABLE_NAME = '%v')", schema, name)
	if err != nil {
		return false
	}
	return row.Next()
}
