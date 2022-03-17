package client

import (
	"fmt"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-imx-client/types"
)

type GetTransferRequest struct {
	TransferID string
}

func (c *Client) GetTransfer(request GetTransferRequest) (response types.Transfer, err error) {
	const op = "Client.GetTransfer"

	URL := fmt.Sprintf("%s/%s", TransfersURL, request.TransferID)

	err = c.RestRequest("GET", URL, nil, nil, &response)
	if err != nil {
		return response, ez.Wrap(op, err)
	}

	return response, nil

}

type TransferERC20Request struct {
	Type         string `json:"type,omitempty"`
	TokenAddress string `json:"token_address,omitempty"`
	ToAddress    string `json:"to_address,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
}

type TransferERC20Response struct {
	TransferIDs []int `json:"transfer_ids"`
}

func (c *Client) TransferERC20(request *TransferERC20Request) (response TransferERC20Response, err error) {
	const op = "Client.Transfer"

	request.Type = "ERC20"

	err = c.SDKRequest("POST", SDKTransferERC20URL, nil, request, &response)
	if err != nil {
		return response, ez.Wrap(op, err)
	}

	return response, nil
}

type TransferERC721Request struct {
	Type         string `json:"type,omitempty"`
	TokenID      string `json:"token_id,omitempty"`
	TokenAddress string `json:"token_address,omitempty"`
	ToAddress    string `json:"to_address,omitempty"`
}

type TransferERC721Response struct {
	TransferIDs []int `json:"transfer_ids"`
}

func (c *Client) TransferERC721(request *TransferERC721Request) (response TransferERC721Response, err error) {
	const op = "Client.Transfer"

	request.Type = "ERC721"

	err = c.SDKRequest("POST", SDKTransferERC721URL, nil, request, &response)
	if err != nil {
		return response, ez.Wrap(op, err)
	}

	return response, nil
}