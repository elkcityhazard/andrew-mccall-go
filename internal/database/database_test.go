package database

import (
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
)

func Test_NewDriver(t *testing.T) {

	_, err := NewDriver(&config.AppConfig{
		DSN: "There is none",
	})

	if err == nil {
		t.Error("should  be an error")
	}

	_, err = NewDriver(&config.AppConfig{})

	if err != nil {
		t.Error("There should not be an error")
	}

}
