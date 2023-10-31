package pqtest

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	dbUser   = "postgres"
	dbPassWd = "postgres"
	dbHost   = "localhost"
	dbName   = "test"
	dbPort   = "5432"
)

// TODO: need to change to use env value instead
func InitTestDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassWd, dbName)

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

	stmts := strings.Split(string(file), ";\n")
	for _, stmt := range stmts {
		if len(stmt) == 0 {
			continue
		}
		_, err = db.Exec(stmt)
		if err != nil {
			log.Fatal("failed to exec sql file")
		}
	}
}
