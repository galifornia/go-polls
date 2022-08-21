package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer dbpool.Close()

	return dbpool, nil
}

const TRIES = 10

func ConnectToDB() *pgxpool.Pool {
	dsn := os.Getenv("DSN")
	count := 0

	for count < TRIES {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			count++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
