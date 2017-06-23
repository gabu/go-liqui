package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestOrderInfo(t *testing.T) {
	liqui := newAuthClient()
	ctx := context.Background()
	ret, err := liqui.OrderInfo(ctx, "15081866")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
