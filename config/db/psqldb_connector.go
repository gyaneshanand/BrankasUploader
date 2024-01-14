package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// InitDB initializes the database connection pool
func InitDB() error {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=require",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return err
	}

	poolConfig.MaxConns = 10 // We can set this in env
	poolConfig.MinConns = 2  // We can set this in env
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	DB, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return err
	}

	// Check if the database connection is working by pinging the database
	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	slog.Info("Connected to the DB successfully. ")

	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	slog.Info("Closing the DB connection.")
	DB.Close()
}

// AcquireDBConnection acquires a connection from the pool
func AcquireDBConnection() (*pgxpool.Conn, error) {
	conn, err := DB.Acquire(context.Background())
	if err != nil {
		slog.Error("Error acquiring connection:", err)
		return nil, err
	}
	slog.Info("DB connection pool is acquired now.")
	return conn, nil
}

// ReleaseDBConnection releases the acquired connection
func ReleaseDBConnection(conn *pgxpool.Conn) {
	slog.Info("Releasing DB connection.")
	conn.Release()
}
