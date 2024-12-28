package static_test

import (
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/internal/static"
)

func Test_GetStaticDir(t *testing.T) {
	staticDir := static.GetStaticDir()

	_, err := staticDir.ReadDir("dist")

	if err != nil {
		t.Fatal("expected to be able to read dist directory, but got error: ", err.Error())
	}

}
