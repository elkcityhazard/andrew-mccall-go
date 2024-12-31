package sqldbconn

import (
	"database/sql"
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
)

func Test_NewSQLDbConn(t *testing.T) {
	mockApp := &config.AppConfig{}
	mockSQLDb := &sql.DB{}
	mockSQLDbConn := NewSQLDbConn(mockApp, mockSQLDb)

	if mockSQLDbConn == nil {
		t.Fatal("Expected an SQLDbConn but got nil")
	}
}
