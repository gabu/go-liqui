package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestTradeHistory(t *testing.T) {
	liqui := newAuthClient()
	ctx := context.Background()
	ret, err := liqui.TradeHistory(ctx, 0, 0, "ptoy_btc")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
