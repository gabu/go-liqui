# go-liqui

An unofficial [Liqui public and trade API](https://liqui.io/api) client for Go.

## Supports

### Liqui public API

- [x] info
- [x] ticker
- [x] depth
- [x] trades

### Liqui trade API (needs an authentication)

- [x] getInfo
- [ ] Trade
- [x] ActiveOrders
- [x] OrderInfo
- [ ] CancelOrder
- [x] TradeHistory
- [ ] WithdrawCoin Method (not active now, will be available soon)

## Usage

### Ticker

```go
package main

import (
	"context"
	"fmt"

	"github.com/gabu/go-liqui"
)

func main() {
	liqui := liqui.NewClient()
	ctx := context.Background()
	ticker, err := liqui.Ticker(ctx, "ltc_btc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ticker)
}
```

### Authentication

```go
func main() {
	liqui := liqui.NewClient().Auth("YOUR API KEY", "YOUR API SECRET")
}
```
