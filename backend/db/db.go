package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)


type Database struct {
	Context *pgxpool.Pool
}

func Initialize() (*Database, error) {
    dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
    }

	return &Database{
		Context: dbpool,
	}, nil
}