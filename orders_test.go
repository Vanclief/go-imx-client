package imx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOrders(t *testing.T) {

	client, err := NewClient("mainnet")
	assert.Nil(t, err)

	request := &ListOrdersRequest{
		Status:           "active",
		SellTokenAddress: "0xa4ddc0932b4e97523f8198eda7a28dac2327d365",
	}

	response, err := client.ListOrders(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	for _, order := range response.Result {
		assert.NotNil(t, order.OrderID)
		assert.NotNil(t, order.Sell.Data.TokenID)
	}
}
