package client

import (
	"encoding/json"
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

	token := TokenToMint{
		Type: "ERC721",
		Data: data,
	}

	mint := Mint{
		User:          "0x",
		Tokens:        []TokenToMint{token},
		Nonce:         498388787,
		AuthSignature: "",
	}

	// Signature
	b, err := json.Marshal(mint)
	assert.Nil(t, err)

	// Throwaway PK, don't waste your time checking if it has funds
	hexPrivateKey := "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e"

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := Sign(b, privateKey)
	assert.Equal(t, "0x7803e1e4da85e8d0fb3657471304696e1d01cc994851303d659da27c9671a10053d9fa95f16f4dfaa79f31df0c43134bd1889a9ed93e90f2976333618d96e22301", signature)
	assert.Nil(t, err)
}

func TestMintSignature(t *testing.T) {

	// Data
	data := &MintData{
		ID:           "1",
		Blueprint:    "metadata",
		TokenAddress: "0x",
	}

	token := TokenToMint{
		Type: "ERC721",
		Data: data,
	}

	mint := Mint{
		User:          "0x",
		Tokens:        []TokenToMint{token},
		Nonce:         498388787,
		AuthSignature: "",
	}

	hexPrivateKey := "38b7ca1ab75167a94477ce6682fc687a36ce9cdff3e6a6d582177c0c3b407d2e"

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := MintSignature(mint, privateKey)
	assert.Equal(t, "0x6108d3ccabb8bb369dde0676549185981e4fa2de4ad96d3d8f22478e140173d90f3dc017fb53c432a3c9633f49a8c60281ed12d90d7fa18c70822297c904468e01", signature)
	assert.Nil(t, err)
}
