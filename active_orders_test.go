package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestActiveOrders(t *testing.T) {
	liqui := newAuthClient()
	ctx := context.Background()
	ret, err := liqui.ActiveOrders(ctx, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
