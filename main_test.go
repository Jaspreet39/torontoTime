package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	// create a req
	req, err := http.NewRequest("Get", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTime)
	handler.ServeHTTP(rr, req)

	// check the status code is that we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong stauts code: got %v want %v", status, http.StatusOK)
	}

	var response TimeResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// "2006-1-2 15:4:5"
	_, err = time.Parse("2006-1-2 15:4:5", response.CurrentTime)
	if err != nil {
		t.Errorf("Handler returned unexpected body: got %v", response.CurrentTime)
	}
}
