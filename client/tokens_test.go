package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {

	client, err := NewClient(Ropsten, "", "")
	assert.Nil(t, err)

	request := &GetTokenRequest{
		TokenAddress: "0x73f99ca65b1a0aef2d4591b1b543d789860851bf",
	}

	response, err := client.GetToken(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
