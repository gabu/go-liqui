package liqui

import (
	"context"
	"net/url"
)

type ActiveOrders map[string]ActiveOrder

type ActiveOrder struct {
	Pair             string
	Type             string
	Amount           float64
	Rate             float64
	TimestampCreated int `json:"timestamp_created"`
}

// ActiveOrders returns the list of your active orders.
// To use this method you need a privilege of the info key.
// If the order disappears from the list, it was either executed or canceled.
func (c *Client) ActiveOrders(ctx context.Context, pair string) (ActiveOrders, error) {
	v := url.Values{}
	v.Set("pair", pair)

	req, err := c.newAuthenticatedRequest(ctx, "ActiveOrders", v)
	if err != nil {
		return ActiveOrders{}, err
	}

	var ret = &ActiveOrders{}
	_, err = c.do(req, ret)
	if err != nil {
		switch e := err.(type) {
		case *ErrorResponse:
			if e.Response.Response.StatusCode == 200 {
				return ActiveOrders{}, nil
			}
		default:
			return ActiveOrders{}, err
		}

	}
	return *ret, nil
}
