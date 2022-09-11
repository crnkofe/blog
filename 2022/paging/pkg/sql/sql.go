package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func connectToMySQL() (*sqlx.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/paging")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return sqlx.NewDb(db, "sql-o"), nil
}

func Init() error {
	var err error
	db, err = connectToMySQL()
	return err
}

func Close() error {
	return db.Close()
}