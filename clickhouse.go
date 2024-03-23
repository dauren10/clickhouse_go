package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	// Connection parameters
	host := "localhost"
	nativePort := "19000"
	database := "helloworld"
	connectParams := fmt.Sprintf("tcp://%s:%s?database=%s&read_timeout=20&write_timeout=20", host, nativePort, database)

	// Open a connection to ClickHouse
	db, err := sql.Open("clickhouse", connectParams)
	if err != nil {
		fmt.Println("Error connecting to ClickHouse:", err)
		return
	}
	defer db.Close()

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error beginning transaction:", err)
		return
	}
	defer tx.Rollback() // Rollback the transaction if it's not committed

	// Prepare the INSERT statement
	insertQuery := `
		INSERT INTO helloworld.my_first_table (user_id, message, timestamp, metric) VALUES (?, ?, ?, ?)
	`
	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	// Execute INSERT query for each row
	rows := [][]interface{}{
		{101, "Hello, ClickHouse!", time.Now(), -1.0},
		{102, "Insert a lot of rows per batch", time.Now(), 1.41421},
		{102, "Sort your data based on your commonly-used queries", time.Now(), 2.718},
		{101, "Granules are the smallest chunks of data read", time.Now().Add(5 * time.Second), 3.14159},
	}

	for _, row := range rows {
		_, err := stmt.Exec(row[0], row[1], row[2], row[3])
		if err != nil {
			fmt.Println("Error inserting data into ClickHouse:", err)
			return
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		fmt.Println("Error committing transaction:", err)
		return
	}

	fmt.Println("Data inserted successfully.")
}
