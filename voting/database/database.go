package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func openDB(dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())
	return conn, nil
}

const TRIES = 10

func ConnectToDB() *pgx.Conn {
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
