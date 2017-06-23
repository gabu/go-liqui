package liqui

import (
	"context"
)

type Info struct {
	ServerTime int `json:"server_time"`
	Pairs      map[string]PairInfo
}

type PairInfo struct {
	DecimalPlaces int     `json:"decimal_places"`
	MinPrice      float64 `json:"min_price"`
	MaxPrice      float64 `json:"max_price"`
	MinAmount     float64 `json:"min_amount"`
	Hidden        int
	Fee           float64
}

// Info returns all the information about currently active pairs,
// such as the maximum number of digits after the decimal point,
// the minimum price, the maximum price, the minimum transaction size,
// whether the pair is hidden, the commission for each pair.
func (c *Client) Info(ctx context.Context) (Info, error) {
	req, err := c.newRequest(ctx, "GET", "info", nil)
	if err != nil {
		return Info{}, err
	}

	var ret = &Info{}
	_, err = c.do(req, ret)
	if err != nil {
		return Info{}, err
	}
	return *ret, nil
}
