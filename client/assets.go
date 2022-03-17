package client

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/vanclief/ez"
)

type Collection struct {
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

type Asset struct {
	ID           string      `json:"id"`
	TokenAddress string      `json:"token_address"`
	TokenID      string      `json:"token_id"`
	User         string      `json:"user"`
	Status       string      `json:"eth"`
	URI          string      `json:"uri"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	ImageURL     string      `json:"image_url"`
	Metadata     interface{} `json:"metadata"`
	Collection   Collection  `json:"collection"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
}

type ListAssetsRequest struct {
	PageSize            int    `url:"page_size,omitempty"`
	Cursor              string `url:"cursor,omitempty"`
	OrderBy             string `url:"order_by,omitempty"`
	Direction           string `url:"direction,omitempty"`
	User                string `url:"user,omitempty"`
	Status              string `url:"status,omitempty"`
	Name                string `url:"name,omitempty"`
	Metadata            string `url:"metadata,omitempty"`
	SellOrders          bool   `url:"sell_orders,omitempty"`
	BuyOrders           bool   `url:"buy_orders,omitempty"`
	IncludeFees         bool   `url:"include_fees,omitempty"`
	Collection          string `url:"collection,omitempty"`
	UpdatedMinTimestamp string `url:"updated_min_timestamp,omitempty"`
	UpdatedMaxTimestamp string `url:"updated_max_timestamp,omitempty"`
}

type ListAssetsResponse struct {
	Cursor    string  `json:"cursor"`
	Remaining int     `json:"remaining"`
	Result    []Asset `json:"result"`
}

func (c *Client) ListAssets(request *ListAssetsRequest) (*ListAssetsResponse, error) {
	const op = "Client.ListAssets"

	response := &ListAssetsResponse{}

	data, err := query.Values(request)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	err = c.RestRequest("GET", AssetsURL, data, nil, response)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil
}

type GetAssetRequest struct {
	TokenAddress string
	TokenID      string
}

func (c *Client) GetAsset(request *GetAssetRequest) (*Asset, error) {
	const op = "Client.GetAssets"

	response := &Asset{}

	endpoint := fmt.Sprintf("%s/%s/%s", AssetsURL, request.TokenAddress, request.TokenID)

	err := c.RestRequest("GET", endpoint, nil, nil, response)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil

}
