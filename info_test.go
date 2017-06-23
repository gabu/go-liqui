package liqui

import (
	"context"
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	liqui := NewClient()
	ctx := context.Background()
	ret, err := liqui.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
