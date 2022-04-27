package client

import (
	"fmt"

	"github.com/vanclief/ez"
)

type Token struct {
	Name         string `json:"name"`
	ImageURL     string `json:"image_url"`
	TokenAddress string `json:"token_address"`
	Symbol       string `json:"symbol"`
	Decimals     string `json:"decimals"`
	Quantum      string `json:"quantum"`
}

type GetTokenRequest struct {
	TokenAddress string
}

func (c *Client) GetToken(request *GetTokenRequest) (*Token, error) {
	const op = "Client.GetToken"

	response := &Token{}

	endpoint := fmt.Sprintf("%s/%s", TokensURL, request.TokenAddress)

	err := c.RestRequest("GET", endpoint, nil, nil, response)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil

}
