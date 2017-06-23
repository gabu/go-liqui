package liqui

import (
	"context"
)

type Tickers map[string]Ticker

type Ticker struct {
	High    float64
	Low     float64
	Avg     float64
	Vol     float64
	VolCur  float64 `json:"vol_cur"`
	Last    float64
	Buy     float64
	Sell    float64
	Updated int
}

// Ticker returns all the information about currently active pairs, such as: the maximum price,
// the minimum price, average price, trade volume, trade volume in currency, the last trade,
// Buy and Sell price. All information is provided over the past 24 hours.
//
// Specify multiple pairs join by "-" to pair, returns multiple tickers. e.g. eth_btc-ltc_btc
func (c *Client) Ticker(ctx context.Context, pair string) (Tickers, error) {
	req, err := c.newRequest(ctx, "GET", "ticker/"+pair, nil)
	if err != nil {
		return Tickers{}, err
	}

	var ret = &Tickers{}
	_, err = c.do(req, ret)
	if err != nil {
		return Tickers{}, err
	}
	return *ret, nil
}
