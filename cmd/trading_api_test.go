package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBuy(t *testing.T) {
	req, err := http.NewRequest("POST", "/createBuy", bytes.NewBuffer([]byte(`{"amount": 100, "symbol": "AAPL"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	http.HandlerFunc(CreateBuy).ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rec.Code)
	}
	id := rec.Body.String()
	if id == "" {
		t.Error("Expected trade ID")
	}
}

func TestCreateSell(t *testing.T) {
	req, err := http.NewRequest("POST", "/createSell", bytes.NewBuffer([]byte(`{"amount": 100, "symbol": "AAPL"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	http.HandlerFunc(CreateSell).ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rec.Code)
	}
	id := rec.Body.String()
	if id == "" {
		t.Error("Expected trade ID")
	}
}

func TestCommitTrade(t *testing.T) {
	req, err := http.NewRequest("POST", "/commitTrade", bytes.NewBuffer([]byte(`{"id": "1"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	http.HandlerFunc(CommitTrade).ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
