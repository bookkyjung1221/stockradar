package utils

import "database/sql"

func ConnectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5433 user=stockradar dbname=stockradar password=stockradar sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
