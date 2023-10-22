package pq

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitTestDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_TEST_HOST"), os.Getenv("DB_TEST_PORT"), os.Getenv("DB_TEST_USER"), os.Getenv("DB_TEST_PASSWORD"), os.Getenv("DB_TEST_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgresql")
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database is not healthy")
	}

	return db, nil
}

func SetUp(db *sql.DB, fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read sql file")
	}

	_, err = db.Exec(string(file))
	if err != nil {
		log.Fatal("failed to exec sql file")
	}
}
