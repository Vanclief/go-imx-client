package imx

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestMint(t *testing.T) {

	client := New("")

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
		User:          "0x",
		Tokens:        []Token{token},
		Nonce:         498388787,
		AuthSignature: "",
	}

	hexPrivateKey := "XXX"

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	assert.Nil(t, err)

	signature, err := MintSignature(mint, privateKey)
	assert.Nil(t, err)

	mint.AuthSignature = signature

	request := &MintRequest{Mints: []Mint{mint}}

	response, err := client.Mint(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
