package imx

import (
	"github.com/vanclief/ez"
)

type MintRequest struct {
	Mints []Mint `json:"mints"`
}

type Mint struct {
	User          string  `json:"user"`
	Tokens        []Token `json:"tokens"`
	Nonce         int64   `json:"nonce"`
	AuthSignature string  `json:"auth_signature"`
}

type Token struct {
	Type string    `json:"type"`
	Data *MintData `json:"data"`
}

type MintData struct {
	ID           string `json:"id"`
	Blueprint    string `json:"blueprint"`
	TokenAddress string `json:"token_address"`
}

type MintResponse struct {
}

func (c *Client) Mint(request *MintRequest) (*MintResponse, error) {
	const op = "imx.Client.Mint"

	URL := "mints"

	response := &MintResponse{}

	err := c.httpRequest("POST", URL, nil, request, response)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil
}
