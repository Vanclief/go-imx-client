package client

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestMint(t *testing.T) {

	client, err := NewClient("ropsten")
	assert.NotNil(t, client)
	assert.Nil(t, err)

	// Data
	data := &MintData{
		ID:           "9999",
		Blueprint:    "metadata",
		TokenAddress: "0x",
	}

	token := Token{
		Type: "ERC721",
		Data: data,
	}

	mint := Mint{
		User:          "0x1ea42bdf38c50dB0fdF396612fA8C59E09EE023e",
		Tokens:        []Token{token},
		Nonce:         498388787,
		AuthSignature: "",
	}

	hexPrivateKey := "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e"

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	assert.Nil(t, err)

	signature, err := MintSignature(mint, privateKey)
	assert.Nil(t, err)

	mint.AuthSignature = signature

	assert.NotNil(t, mint.AuthSignature)

	// request := &MintRequest{Mints: []Mint{mint}}

	// Can't do with a real contract
	// response, err := client.Mint(request)
	// assert.Nil(t, err)
	// assert.NotNil(t, response)
}
