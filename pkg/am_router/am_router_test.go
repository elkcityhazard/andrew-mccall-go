package amrouter

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_AddRoute(t *testing.T) {
	rtr := NewRouter()

	err := rtr.AddRoute("GET", "/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	}))

	if err != nil {
		t.Error("expected no error, but got an error")
	}
}
