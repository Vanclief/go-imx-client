package imx

import (
	"net/url"
	"strings"
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
	User         string `json:"user"`
	TokenAddress string `json:"token_address"`
}

type ListAssetsResponse struct {
	Result []Asset `json:"result"`
}

func (c *Client) ListAssets(request *ListAssetsRequest) (*ListAssetsResponse, error) {

	response := &ListAssetsResponse{}

	data := url.Values{
		"user":       {strings.ToLower(request.User)},
		"collection": {strings.ToLower(request.TokenAddress)},
	}

	err := c.httpRequest("GET", "assets", data, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
