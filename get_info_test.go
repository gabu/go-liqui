package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestGetInfo(t *testing.T) {
	liqui := newAuthClient()
	ctx := context.Background()
	ret, err := liqui.GetInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
