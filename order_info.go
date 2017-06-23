package liqui

import (
	"context"
	"net/url"
)

type OrderInfos map[string]OrderInfo

type OrderInfo struct {
	Pair             string
	Type             string
	StartAmount      float64 `json:"start_amount"`
	Amount           float64
	Rate             float64
	TimestampCreated int `json:"timestamp_created"`
	Status           int
}

// OrderInfo returns trade history.
// To use this method you need a privilege of the info key.
func (c *Client) OrderInfo(ctx context.Context, orderID string) (OrderInfos, error) {
	v := url.Values{}
	v.Set("order_id", orderID)

	req, err := c.newAuthenticatedRequest(ctx, "OrderInfo", v)
	if err != nil {
		return OrderInfos{}, err
	}

	var ret = &OrderInfos{}
	_, err = c.do(req, ret)
	if err != nil {
		switch e := err.(type) {
		case *ErrorResponse:
			if e.Response.Response.StatusCode == 200 {
				return OrderInfos{}, nil
			}
		default:
			return OrderInfos{}, err
		}

	}
	return *ret, nil
}
