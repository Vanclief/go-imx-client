package client

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/vanclief/ez"
	"github.com/vanclief/go-imx-client/types"
)

type ListTransfersRequest struct {
	PageSize            int    `url:"page_size,omitempty"`
	Cursor              string `url:"cursor,omitempty"`
	OrderBy             string `url:"order_by,omitempty"`
	Direction           string `url:"direction,omitempty"`
	User                string `url:"user,omitempty"`
	Receiver            string `url:"receiver,omitempty"`
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

type ListTransfersResponse struct {
	Cursor    string           `json:"cursor"`
	Remaining int              `json:"remaining"`
	Result    []types.Transfer `json:"result"`
}

func (c *Client) ListTransfers(request ListTransfersRequest) (response ListTransfersResponse, err error) {
	const op = "Client.ListTransfers"

	data, err := query.Values(request)
	if err != nil {
		return response, ez.Wrap(op, err)
	}

	err = c.RestRequest("GET", TransfersURL, data, nil, &response)
	if err != nil {
		return response, ez.Wrap(op, err)
	}

	return response, nil
}

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
