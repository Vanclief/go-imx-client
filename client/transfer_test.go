package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/ez"
)

func TestTransferERC20(t *testing.T) {

	// Throwaway PK, don't waste your time checking if it has funds
	client, err := NewClient(Ropsten, "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e", "DvukuyBzEK-JyP6zp1NVeNVYLJCrzjp_")
	assert.Nil(t, err)

	request := &TransferERC20Request{}

	response, err := client.TransferERC20(request)
	ez.ErrorStacktrace(err)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestTransferERC721(t *testing.T) {

	// Throwaway PK, don't waste your time checking if it has funds
	client, err := NewClient(Ropsten, "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e", "DvukuyBzEK-JyP6zp1NVeNVYLJCrzjp_")
	assert.Nil(t, err)

	request := &TransferERC721Request{
		TokenID:      "1064",
		TokenAddress: "0x3d25036695dafab7eee3465ff146f6d6c6d0045b",
		ToAddress:    "0x969aF39505BAC751b0A70f6AFf8Ef38a5011eb1e",
	}

	response, err := client.TransferERC721(request)
	ez.ErrorStacktrace(err)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
