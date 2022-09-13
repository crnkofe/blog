package sql

import (
	"context"
	"github.com/crnkofe/blog/2022/paging/pkg/paging"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type computer struct {
	db *pgx.Conn
}

func Computer() computer {
	return computer{
		db: db,
	}
}

func (c *computer) GetPaged(ctx context.Context, lastID int, limit int) ([]paging.Computer, error) {
	computers := []paging.Computer{}
	err := pgxscan.Select(ctx, c.db, &computers, "SELECT id, name FROM computer WHERE id > $1 ORDER BY id LIMIT $2", lastID, limit)
	return computers, err
}