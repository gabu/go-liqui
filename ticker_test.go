package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestTicker(t *testing.T) {
	liqui := NewClient()
	ctx := context.Background()
	ret, err := liqui.Ticker(ctx, "ltc_btc")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
