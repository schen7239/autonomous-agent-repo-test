package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Trade represents a trade
 type Trade struct {
	ID       string  `json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

var (
	tradesMutex sync.Mutex
	trades      = make(map[string]Trade)
)

// CreateBuy creates a new buy trade
func CreateBuy(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tradesMutex.Lock()
	defer tradesMutex.Unlock()
	trade.ID = "12345" // Temporary ID
	trades[trade.ID] = trade
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trade)
}

// CreateSell creates a new sell trade
func CreateSell(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tradesMutex.Lock()
	defer tradesMutex.Unlock()
	trade.ID = "67890" // Temporary ID
	trades[trade.ID] = trade
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trade)
}

// CommitTrade commits a trade
func CommitTrade(w http.ResponseWriter, r *http.Request) {
	var tradeID string
	if err := json.NewDecoder(r.Body).Decode(&tradeID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tradesMutex.Lock()
	defer tradesMutex.Unlock()
	if _, ok := trades[tradeID]; !ok {
		http.Error(w, "Trade not found", http.StatusNotFound)
		return
	}
	// Commit logic here
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "committed"})
}

func main() {
	http.HandleFunc("/createbuy", CreateBuy)
	http.HandleFunc("/createsell", CreateSell)
	http.HandleFunc("/committrade", CommitTrade)
	log.Fatal(http.ListenAndServe(":8080", nil))
}