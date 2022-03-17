package client

import (
	"github.com/google/go-querystring/query"
	"github.com/vanclief/ez"
)

type OrderData struct {
	TokenID      string      `json:"token_id"`
	ID           string      `json:"id"`
	TokenAddress string      `json:"token_address"`
	Quantity     string      `json:"quantity"`
	Properties   interface{} `json:"properties"`
}

type OrderDetail struct {
	Type string    `json:"type"`
	Data OrderData `json:"data"`
}

type Order struct {
	OrderID             int         `json:"order_id"`
	Status              string      `json:"status"`
	User                string      `json:"user"`
	Sell                OrderDetail `json:"sell"`
	Buy                 OrderDetail `json:"buy"`
	AmountSold          string      `json:"amount_sold"`
	ExpirationTimestamp string      `json:"expiration_timestamp"`
	Timestamp           string      `json:"timestamp"`
	UpdatedTimestamp    string      `json:"updated_timestamp"`
}

type ListOrdersRequest struct {
	PageSize            int    `url:"page_size,omitempty"`
	Cursor              string `url:"cursor,omitempty"`
	OrderBy             string `url:"order_by,omitempty"`
	Direction           string `url:"direction,omitempty"`
	User                string `url:"user,omitempty"`
	Status              string `url:"status,omitempty"`
	MinTimestamp        string `url:"min_timestamp,omitempty"`
	MaxTimestamp        string `url:"max_timestamp,omitempty"`
	UpdatedMinTimestamp string `url:"updated_min_timestamp,omitempty"`
	UpdatedMaxTimestamp string `url:"updated_max_timestamp,omitempty"`
	BuyTokenType        string `url:"buy_token_type,omitempty"`
	BuyTokenID          string `url:"buy_token_id,omitempty"`
	BuyAssetID          string `url:"buy_asset_id,omitempty"`
	BuyTokenAddress     string `url:"buy_token_address,omitempty"`
	BuyTokenName        string `url:"buy_token_name,omitempty"`
	BuyMinQuantity      string `url:"buy_min_quantity,omitempty"`
	BuyMaxQuantity      string `url:"buy_max_quantity,omitempty"`
	BuyMetadata         string `url:"buy_metadata,omitempty"`
	SellTokenType       string `url:"sell_token_type,omitempty"`
	SellTokenID         string `url:"sell_token_id,omitempty"`
	SellAssetID         string `url:"sell_asset_id,omitempty"`
	SellTokenAddress    string `url:"sell_token_address,omitempty"`
	SellMinQuantity     string `url:"sell_min_quantity,omitempty"`
	SellMaxQuantity     string `url:"sell_max_quantity,omitempty"`
	SellMetadata        string `url:"sell_metadata,omitempty"`
}

type ListOrdersResponse struct {
	Cursor    string  `json:"cursor"`
	Remaining int     `json:"remaining"`
	Result    []Order `json:"result"`
}

func (c *Client) ListOrders(request *ListOrdersRequest) (*ListOrdersResponse, error) {
	const op = "Client.ListOrders"

	response := &ListOrdersResponse{}

	data, err := query.Values(request)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	err = c.RestRequest("GET", OrdersURL, data, nil, response)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil
}
