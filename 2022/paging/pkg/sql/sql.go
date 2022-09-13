package sql

import (
	"context"
	"github.com/jackc/pgx/v4"
)

var (
	db *pgx.Conn
)

func connectToPostgres(ctx context.Context) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Init(ctx context.Context) error {
	var err error
	db, err = connectToPostgres(ctx)
	return err
}

func Close(ctx context.Context) error {
	return db.Close(ctx)
}