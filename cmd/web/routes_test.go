package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_routes(t *testing.T) {

	tests := []struct {
		name   string
		method string
		path   string
		status int
	}{
		{
			name:   "test home route",
			method: "GET",
			path:   "/",
		},
	}

	r := routes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(tt.method, tt.path, nil)

			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v but want %v", status, tt.status)
			}

		})
	}

}
