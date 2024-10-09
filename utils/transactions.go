package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func isDatabaseSeeded(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddTransactionToDB() {

	file, err := os.Open("example_data.csv")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV file: %v", err)
	}

	connStr := "user=stockradar password=stockradar dbname=stockradar port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	seeded, err := isDatabaseSeeded(db)
	if err != nil {
		log.Fatalf("failed to check if the database is seeded: %v", err)
	}

	if seeded {
		fmt.Println("Database is already seeded. No need to seed again.")
		return
	}

	stmt, err := db.Prepare(`
        INSERT INTO transactions (payout, currency, user_id, sale_amount, datetime, shop_name, shop_offset_hour)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `)
	if err != nil {
		log.Fatalf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}

		payout, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatalf("invalid payout in row %d: %v", i+1, err)
		}

		currency := record[1]
		userID := record[2]

		saleAmount, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatalf("invalid sale_amount in row %d: %v", i+1, err)
		}

		datetime, err := time.Parse("2006-01-02 15:04:05", record[4])
		if err != nil {
			log.Fatalf("invalid datetime in row %d: %v", i+1, err)
		}

		shopName := record[5]

		shopOffsetHour, err := strconv.Atoi(record[6])
		if err != nil {
			log.Fatalf("invalid shop_offset_hour in row %d: %v", i+1, err)
		}

		_, err = stmt.Exec(payout, currency, userID, saleAmount, datetime, shopName, shopOffsetHour)
		if err != nil {
			log.Fatalf("failed to insert record at row %d: %v", i+1, err)
		}

		fmt.Printf("Inserted record for user_id %s\n", userID)
	}

	fmt.Println("CSV data has been successfully inserted into the database.")
}
