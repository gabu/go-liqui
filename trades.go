package liqui

import (
	"context"
	"net/url"
	"strconv"
)

type Trades map[string][]Trade

type Trade struct {
	Type      string
	Price     float64
	Amount    float64
	Tid       int
	Timestamp int
}

// Trades returns the information about the last trades.
// Additionally it accepts an optional GET-parameter limit,
// which indicates how many orders should be displayed (150 by default).
// The maximum allowable value is 2000.
//
// Specify multiple pairs join by "-" to pair, returns multiple trades. e.g. eth_btc-ltc_btc
func (c *Client) Trades(ctx context.Context, pair string, limit int) (Trades, error) {
	if limit == 0 {
		limit = 150
	}
	v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	req, err := c.newRequest(ctx, "GET", "trades/"+pair, v)
	if err != nil {
		return Trades{}, err
	}

	var ret = &Trades{}
	_, err = c.do(req, ret)
	if err != nil {
		return Trades{}, err
	}
	return *ret, nil
}
