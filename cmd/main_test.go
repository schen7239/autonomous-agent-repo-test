package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBuy(t *testing.T) {
	req, err := http.NewRequest("POST", "/createBuy", bytes.NewBuffer([]byte(`{"amount": 100.0, "symbol": "AAPL"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	CreateBuy(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestCreateSell(t *testing.T) {
	req, err := http.NewRequest("POST", "/createSell", bytes.NewBuffer([]byte(`{"amount": 100.0, "symbol": "AAPL"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	CreateSell(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestCommitTrade(t *testing.T) {
	req, err := http.NewRequest("POST", "/commitTrade", bytes.NewBuffer([]byte(`{"id": "1"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	CommitTrade(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
