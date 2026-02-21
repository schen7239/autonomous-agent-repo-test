package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Trade struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Amount  float64 `json:"amount"`
	Symbol  string `json:"symbol"`
}

var trades = make(map[string]Trade)
var tradesMutex sync.Mutex

func createBuy(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	trade.Type = "buy"
	tradesMutex.Lock()
	trades[trade.ID] = trade
	tradesMutex.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func createSell(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	trade.Type = "sell"
	tradesMutex.Lock()
	trades[trade.ID] = trade
	tradesMutex.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func commitTrade(w http.ResponseWriter, r *http.Request) {
	var trade Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tradesMutex.Lock()
	if _, ok := trades[trade.ID]; !ok {
		w.WriteHeader(http.StatusNotFound)
		tradesMutex.Unlock()
		return
	}
	delete(trades, trade.ID)
	tradesMutex.Unlock()
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/buy", createBuy)
	http.HandleFunc("/api/sell", createSell)
	http.HandleFunc("/api/commit", commitTrade)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
