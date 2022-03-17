package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOrders(t *testing.T) {

	client, err := NewClient(Ropsten, "", "")
	assert.Nil(t, err)

	request := &ListOrdersRequest{
		Status:           "active",
		SellTokenAddress: "0xa4ddc0932b4e97523f8198eda7a28dac2327d365",
	}

	response, err := client.ListOrders(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Cursor)
	assert.NotNil(t, response.Remaining)
	// assert.NotEmpty(t, response.Cursor, "")
	// assert.Equal(t, response.Remaining, 1)

	for _, order := range response.Result {
		assert.NotNil(t, order.OrderID)
		assert.NotNil(t, order.Sell.Data.TokenID)
	}
}
