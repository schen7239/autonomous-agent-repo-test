package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Trade represents a trade operation
type Trade struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Amount  float64 `json:"amount"`
	Symbol  string `json:"symbol"`
}

var (
	trades    = make(map[string]Trade)
	tradesMu  sync.Mutex
	tradeID   int
)

// createTrade creates a new trade and returns its ID
func createTrade(t Trade) string {
	tradesMu.Lock()
	defer tradesMu.Unlock()
	tradeID++
	t.ID = strconv.Itoa(tradeID)
	trades[t.ID] = t
	return t.ID
}

// CreateBuy handles the creation of a buy trade
func CreateBuy(w http.ResponseWriter, r *http.Request) {
	var t Trade
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.Type = "buy"
	id := createTrade(t)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

// CreateSell handles the creation of a sell trade
func CreateSell(w http.ResponseWriter, r *http.Request) {
	var t Trade
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.Type = "sell"
	id := createTrade(t)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

// CommitTrade commits a trade by its ID
func CommitTrade(w http.ResponseWriter, r *http.Request) {
	var t Trade
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tradesMu.Lock()
	defer tradesMu.Unlock()
	if _, ok := trades[t.ID]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/createBuy", CreateBuy)
	http.HandleFunc("/createSell", CreateSell)
	http.HandleFunc("/commitTrade", CommitTrade)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
