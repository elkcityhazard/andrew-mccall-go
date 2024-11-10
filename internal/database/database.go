package database

import (
	"database/sql"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDriver(app *config.AppConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", app.DSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}
