package utils_test

import (
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

func Test_NewUtils(t *testing.T) {
	u := utils.NewUtil()

	if u == nil {
		t.Fatalf("Expected u to not be nil but got nil")
	}
}
