package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func InitTestDB() (*sql.DB, error) {
	dbConfig := mysql.Config{
		User:      os.Getenv("DB_TEST_USER"),
		Passwd:    os.Getenv("DB_TEST_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_TEST_ADDR"),
		DBName:    os.Getenv("DB_TEST_NAME"),
		ParseTime: true,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql")
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Second * 10)

	pingErr := db.Ping()
	if pingErr != nil {
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
