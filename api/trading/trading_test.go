package trading

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBuy(t *testing.T) {
	req, err := http.NewRequest("POST", "/buy", bytes.NewBuffer([]byte(`{"quantity": 10}`)))
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
	req, err := http.NewRequest("POST", "/sell", bytes.NewBuffer([]byte(`{"quantity": 10}`)))
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
	req, err := http.NewRequest("POST", "/commit", bytes.NewBuffer([]byte(`{"id": "id-0"}`)))
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
