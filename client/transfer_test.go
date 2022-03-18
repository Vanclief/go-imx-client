package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransfer(t *testing.T) {

	client, err := NewClient(Ropsten, "", "")
	assert.Nil(t, err)

	request := GetTransferRequest{
		TransferID: "3490228",
	}

	response, err := client.GetTransfer(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestTransferERC20(t *testing.T) {

	// Throwaway PK, don't waste your time checking if it has funds
	// client, err := NewClient(Ropsten, "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e", "DvukuyBzEK-JyP6zp1NVeNVYLJCrzjp_")
	// assert.Nil(t, err)

	// request := &TransferERC20Request{
	// 	TokenID:      "6805",
	// 	Symbol:       "FCT",
	// 	TokenAddress: "0x73f99ca65b1a0aef2d4591b1b543d789860851bf",
	//  ToAddress:    "0xc0324Dca5073Df1aaf26730471718c500d31cA01",
	// 	Amount:       "1",
	// }

	// response, err := client.TransferERC20(request)
	// assert.Nil(t, err)
	// assert.NotNil(t, response)
}

func TestTransferERC721(t *testing.T) {

	// Throwaway PK, don't waste your time checking if it has funds
	// client, err := NewClient(Ropsten, "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e", "DvukuyBzEK-JyP6zp1NVeNVYLJCrzjp_")
	// assert.Nil(t, err)

	// request := &TransferERC721Request{
	// 	TokenID:      "1500",
	// 	TokenAddress: "0x3d25036695dafab7eee3465ff146f6d6c6d0045b",
	// 	ToAddress:    "0xc0324Dca5073Df1aaf26730471718c500d31cA01",
	// }

	// response, err := client.TransferERC721(request)
	// assert.Nil(t, err)
	// assert.NotNil(t, response)
}
