package sql

import (
	"context"
	"github.com/crnkofe/blog/2022/paging/pkg/paging"
	"github.com/jmoiron/sqlx"

)

type computer struct {
	db sqlx.ExtContext
}

func Computer() computer {
	return computer{
		db: db,
	}
}

func (c *computer) GetPaged(ctx context.Context, lastID int, limit int) ([]paging.Computer, error) {
	computers := []paging.Computer{}
	err := sqlx.SelectContext(ctx, c.db, &computers,
		"SELECT * FROM computer WHERE id > ? ORDER BY id LIMIT ?", lastID, limit)
	if err != nil {
		return []paging.Computer{}, err
	}
	return computers, err
}