package liqui

import (
	"context"
	"net/url"
	"strconv"
)

type Depths map[string]Depth

type Depth struct {
	Asks [][]float64
	Bids [][]float64
}

// Depth returns the information about active orders on the pair.
// Additionally it accepts an optional GET-parameter limit,
// which indicates how many orders should be displayed (150 by default).
// Is set to less than 2000.
//
// Specify multiple pairs join by "-" to pair, returns multiple depths. e.g. eth_btc-ltc_btc
func (c *Client) Depth(ctx context.Context, pair string, limit int) (Depths, error) {
	if limit == 0 {
		limit = 150
	}
	v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	req, err := c.newRequest(ctx, "GET", "depth/"+pair, v)
	if err != nil {
		return Depths{}, err
	}

	var ret = &Depths{}
	_, err = c.do(req, ret)
	if err != nil {
		return Depths{}, err
	}
	return *ret, nil
}
