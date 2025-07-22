package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB() (*pgxpool.Pool, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://aksabumilangit:sea@localhost:5432/klinik?sslmode=disable"
	}
	return pgxpool.New(context.Background(), dsn)
}
