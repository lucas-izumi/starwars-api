package main

import (
	"testing"
    "net/http/httptest"
    "net/http"
    "bytes"
    "strings"
)

func TestAddPlanet(t *testing.T) {

	var jsonStr = []byte(`{"name":"Tatooine","climate":"arid","terrain":"desert","films":"5"}`)

	req, err := http.NewRequest("POST", "/planets/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddPlanetEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: \ngot  %v \nwant %v",
			status, http.StatusOK)
	}
	expected := `[{"id":"1","name":"Tatooine","climate":"arid","terrain":"desert","films":"5"}]`
	received := strings.TrimSuffix(rr.Body.String(), "\n")

	if received != expected {
		t.Errorf("handler returned unexpected body: \ngot  %v \nwant %v",
			received, expected)
	}
}