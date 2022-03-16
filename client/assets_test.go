package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAssets(t *testing.T) {

	client, err := NewClient(Ropsten)
	assert.Nil(t, err)

	request := &ListAssetsRequest{
		User:       "0xc0324Dca5073Df1aaf26730471718c500d31cA01",
		Collection: "0x3d25036695dafab7eee3465ff146f6d6c6d0045b",
	}

	response, err := client.ListAssets(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	for _, asset := range response.Result {
		assert.NotNil(t, asset.TokenID)
		assert.NotNil(t, asset.ID)
	}
}

func TestGetAsset(t *testing.T) {

	client, err := NewClient("ropsten")
	assert.Nil(t, err)

	request := &GetAssetRequest{
		TokenAddress: "0x3d25036695dafab7eee3465ff146f6d6c6d0045b",
		TokenID:      "1471",
	}

	response, err := client.GetAsset(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
