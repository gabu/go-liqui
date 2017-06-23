package liqui

import (
	"context"
	"net/url"
	"strconv"
)

type TradeHistory map[string]PastTrade

type PastTrade struct {
	Pair        string
	Type        string
	Amount      float64
	Rate        float64
	OrderID     int  `json:"order_id"`
	IsYourOrder bool `json:"is_your_order"`
	Timestamp   int
}

// TradeHistory returns trade history.
// To use this method you need a privilege of the info key.
func (c *Client) TradeHistory(ctx context.Context, count int, from_id int, pair string) (TradeHistory, error) {
	if count == 0 {
		count = 1000
	}
	v := url.Values{}
	v.Set("count", strconv.Itoa(count))
	v.Set("from_id", strconv.Itoa(from_id))
	v.Set("pair", pair)

	req, err := c.newAuthenticatedRequest(ctx, "TradeHistory", v)
	if err != nil {
		return TradeHistory{}, err
	}

	var ret = &TradeHistory{}
	_, err = c.do(req, ret)
	if err != nil {
		switch e := err.(type) {
		case *ErrorResponse:
			if e.Response.Response.StatusCode == 200 {
				return TradeHistory{}, nil
			}
		default:
			return TradeHistory{}, err
		}

	}
	return *ret, nil
}
