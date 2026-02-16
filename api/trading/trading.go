package trading

import (
	"context"
	"sync"
	"time"
)

// Trade represents a trade operation.
type Trade struct {
	ID        string
	Type      string
	Timestamp time.Time
}

// TradingSystem represents the trading system.
type TradingSystem struct {
	trades map[string]Trade
	mu     sync.Mutex
}

// NewTradingSystem creates a new trading system.
func NewTradingSystem() *TradingSystem {
	return &TradingSystem{
		trades: make(map[string]Trade),
	}
}

// CreateBuy creates a new buy trade and returns the item ID.
func (ts *TradingSystem) CreateBuy(ctx context.Context) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	id := time.Now().Format("20060102150405")
	ts.trades[id] = Trade{
		ID:        id,
		Type:      "buy",
		Timestamp: time.Now(),
	}
	return id, nil
}

// CreateSell creates a new sell trade and returns the item ID.
func (ts *TradingSystem) CreateSell(ctx context.Context) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	id := time.Now().Format("20060102150405")
	ts.trades[id] = Trade{
		ID:        id,
		Type:      "sell",
		Timestamp: time.Now(),
	}
	return id, nil
}

// CommitTrade commits the trade using the item ID.
func (ts *TradingSystem) CommitTrade(ctx context.Context, id string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	trade, exists := ts.trades[id]
	if !exists {
		return fmt.Errorf("trade with id %s does not exist", id)
	}
	
	// Perform trade commitment logic here
	delete(ts.trades, id)
	return nil
}
