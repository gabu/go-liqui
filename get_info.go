package liqui

import (
	"context"
)

type GetInfoResult struct {
	Funds      map[string]float64
	Rights     map[string]bool
	OpenOrders int `json:"open_orders"`
	ServerTime int `json:"server_time"`
}

// GetInfo returns information about the user’s current balance,
// API-key privileges, the number of open orders and Server Time.
//  To use this method you need a privilege of the key info.
func (c *Client) GetInfo(ctx context.Context) (GetInfoResult, error) {
	req, err := c.newAuthenticatedRequest(ctx, "getInfo", nil)
	if err != nil {
		return GetInfoResult{}, err
	}

	var ret = &GetInfoResult{}
	_, err = c.do(req, ret)
	if err != nil {
		return GetInfoResult{}, err
	}
	return *ret, nil
}
