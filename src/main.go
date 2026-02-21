package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createBuy", createBuy)
	r.HandleFunc("/createSell", createSell)
	r.HandleFunc("/commitTrade", commitTrade)
	http.ListenAndServe(":8080", r)
}

func createBuy(w http.ResponseWriter, r *http.Request) {
	tradeMutex.Lock()
	defer tradeMutex.Unlock()
	id := "buy-" + strconv.Itoa(len(trades))
	trades[id] = "buy"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func createSell(w http.ResponseWriter, r *http.Request) {
	tradeMutex.Lock()
	defer tradeMutex.Unlock()
	id := "sell-" + strconv.Itoa(len(trades))
	trades[id] = "sell"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func commitTrade(w http.ResponseWriter, r *http.Request) {
	tradeMutex.Lock()
	defer tradeMutex.Unlock()
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing trade ID", http.StatusBadRequest)
		return
	}
	if _, ok := trades[id]; !ok {
		http.Error(w, "Trade ID not found", http.StatusNotFound)
		return
	}
	delete(trades, id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trade committed"))
}
