package imx

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {

	// Data
	data := &MintData{
		ID:           "1",
		Blueprint:    "metadata",
		TokenAddress: "0x3d25036695dafab7eee3465ff146f6d6c6d0045b",
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

	// Signature
	b, err := json.Marshal(mint)
	if err != nil {
		fmt.Println(err)
	}

	hexPrivateKey := "XXX"

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := Sign(b, privateKey)
	assert.Equal(t, "0x469b07327fc41a2d85b7e69bcf4a9184098835c47cc7575375e3a306c3718ae35702af84f3a62aafeb8aab6a455d761274263d79e7fc99fbedfeaf759d8dc93601", signature)
	assert.Nil(t, err)
}

func TestMintSignature(t *testing.T) {

	// Data
	data := &MintData{
		ID:           "1",
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
	if err != nil {
		log.Fatal(err)
	}

	signature, err := MintSignature(mint, privateKey)
	assert.Equal(t, "0x469b07327fc41a2d85b7e69bcf4a9184098835c47cc7575375e3a306c3718ae35702af84f3a62aafeb8aab6a455d761274263d79e7fc99fbedfeaf759d8dc93601", signature)
	assert.Nil(t, err)
}
