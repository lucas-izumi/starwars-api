package main

import (
	"testing"
    "net/http/httptest"
    "net/http"
    "strings"
)

func TestSearchPlanetNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/planets/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "123")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchPlanetEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
	expected := `{}`
	received := strings.TrimSuffix(rr.Body.String(), "\n")

	if received != expected {
		t.Errorf("handler returned unexpected body: \ngot  %v \nwant %v",
			received, expected)
	}
}