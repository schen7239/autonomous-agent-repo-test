package trading

import (
	"context"
	"testing"
	"time"
)

func TestCreateBuy(t *testing.T) {
	system := NewTradingSystem()
	id, err := system.CreateBuy(context.Background())
	if err != nil {
		t.Fatalf("CreateBuy failed: %v", err)
	}
	
	if id == "" {
		t.Fatalf("CreateBuy returned an empty ID")
	}
}

func TestCreateSell(t *testing.T) {
	system := NewTradingSystem()
	id, err := system.CreateSell(context.Background())
	if err != nil {
		t.Fatalf("CreateSell failed: %v", err)
	}
	
	if id == "" {
		t.Fatalf("CreateSell returned an empty ID")
	}
}

func TestCommitTrade(t *testing.T) {
	system := NewTradingSystem()
	id, err := system.CreateBuy(context.Background())
	if err != nil {
		t.Fatalf("CreateBuy failed: %v", err)
	}
	
	if err := system.CommitTrade(context.Background(), id); err != nil {
		t.Fatalf("CommitTrade failed: %v", err)
	}
	
	if _, exists := system.trades[id]; exists {
		t.Fatalf("Trade was not committed")
	}
}
