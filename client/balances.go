package client

import (
	"fmt"

	"github.com/vanclief/ez"
)

type Balance struct {
	Symbol              string `json:"symbol"`
	Balance             string `json:"balance"`
	PreparingWithdrawal string `json:"preparing_withdrawal"`
	Withdrawable        string `json:"withdrawable"`
}

type ListBalancesRequest struct {
	Owner string `url:"owner,omitempty"`
}

type ListBalancesResponse struct {
	Cursor    string    `json:"cursor"`
	Remaining int       `json:"remaining"`
	Result    []Balance `json:"result"`
}

func (c *Client) ListBalances(request ListBalancesRequest) (response ListBalancesResponse, err error) {
	const op = "Client.ListBalances"

	URL := fmt.Sprintf("%s/%s", BalancesURL, request.Owner)

	err = c.RestRequest("GET", URL, nil, nil, &response)
	if err != nil {
		return response, ez.Wrap(op, err)
	}

	return response, nil
}
