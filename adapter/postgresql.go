package adapter

import (
	"context"
	"fmt"
	"net/url"

	// Initialize "pgx".
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewPostgreSQL instantiates the PostgreSQL database
func NewPostgreSQL() (*pgxpool.Pool, error) {

	// Just hardcoded, we will improve using envvar
	databaseHost := "localhost"
	databasePort := "5432"
	databaseUsername := "root"
	databasePassword := "root"
	databaseName := "commerce"
	databaseSSLMode := "disable"

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
