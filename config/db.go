package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err := sql.Open("pgx",dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err:= db.Ping(); err != nil {
		log.Fatalf("Could not make ping to the database: %v", err)
	}

	log.Println("Successfully connected to database")

	return db
}
