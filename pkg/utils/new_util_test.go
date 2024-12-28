package utils

import "testing"

func Test_NewUtils(t *testing.T) {
	u := NewUtil()

	if u == nil {
		t.Fatalf("Expected u to not be nil but got nil")
	}
}
