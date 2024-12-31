package main

import (
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
	"github.com/elkcityhazard/andrew-mccall-go/internal/handlers"
	"github.com/elkcityhazard/andrew-mccall-go/internal/repository/sqldbconn"
)

func TestMain(m *testing.M) {

	// mock repository.DBServicer
	handlers.SetHandlerRepo(handlers.NewHandlerRepo(&config.AppConfig{}, sqldbconn.NewMockDBRepo()))

}
