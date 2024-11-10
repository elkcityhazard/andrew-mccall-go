package main

import (
	"database/sql"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
)

func newSessionManager(app *config.AppConfig, db *sql.DB) {
	session := scs.New()
	session.Store = mysqlstore.New(db)
	app.SessionManager = session

}
