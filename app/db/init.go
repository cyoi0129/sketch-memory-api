package db

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	godotenv.Load(".env")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	sslmode := os.Getenv("SSL_MODE")
	dbPort := 5432
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB接続")
	return db
}
