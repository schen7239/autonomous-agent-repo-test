package trading

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"strconv"
)

// Trade represents a trade
 type Trade struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

var (
	trades    = make(map[string]Trade)
	tradesMu  sync.Mutex
)

// CreateBuy creates a new buy trade
func CreateBuy(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		log.Println("Error decoding request body", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	trade.Type = "buy"
	trade.ID = generateID()
	tradesMu.Lock()
	trades[trade.ID] = trade
	tradesMu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trade)
}

// CreateSell creates a new sell trade
func CreateSell(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		log.Println("Error decoding request body", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	trade.Type = "sell"
	trade.ID = generateID()
	tradesMu.Lock()
	trades[trade.ID] = trade
	tradesMu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trade)
}

// CommitTrade commits a trade
func CommitTrade(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		log.Println("Error decoding request body", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	tradesMu.Lock()
	if _, ok := trades[trade.ID]; !ok {
		tradesMu.Unlock()
		http.Error(w, "Trade not found", http.StatusNotFound)
		return
	}
	delete(trades, trade.ID)
	tradesMu.Unlock()
	w.WriteHeader(http.StatusOK)
}

// generateID generates a unique ID
func generateID() string {
	return "id-" + strconv.Itoa(len(trades))
}
