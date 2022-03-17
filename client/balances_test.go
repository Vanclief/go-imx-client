package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBalances(t *testing.T) {

	client, err := NewClient(Ropsten, "", "")
	assert.Nil(t, err)

	request := ListBalancesRequest{
		Owner: "0xc0324Dca5073Df1aaf26730471718c500d31cA01",
	}

	response, err := client.ListBalances(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	for _, balance := range response.Result {
		assert.NotNil(t, balance.Symbol)
		assert.NotNil(t, balance.Balance)
		assert.NotNil(t, balance.PreparingWithdrawal)
		assert.NotNil(t, balance.Withdrawable)
	}
}
