package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestTrades(t *testing.T) {
	liqui := NewClient()
	ctx := context.Background()
	ret, err := liqui.Trades(ctx, "ltc_btc-eth_btc", 5)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
