package client

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vanclief/ez"
)

type SigMint struct {
	EtherKey      string  `json:"ether_key"`
	Tokens        []Token `json:"tokens"`
	Nonce         int64   `json:"nonce"`
	AuthSignature string  `json:"auth_signature"`
}

func MintSignature(mint Mint, privateKey *ecdsa.PrivateKey) (string, error) {
	const op = "imx.Mint"

	signMint := SigMint{
		EtherKey:      mint.User,
		Tokens:        mint.Tokens,
		Nonce:         mint.Nonce,
		AuthSignature: mint.AuthSignature,
	}

	b, err := json.Marshal(signMint)
	if err != nil {
		return "", ez.Wrap(op, err)
	}

	return Sign(b, privateKey)
}

func Sign(msg []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	dataHash := crypto.Keccak256Hash(msg)
	hashedMessage := hashMessage(dataHash.Bytes())
	return signMessage(hashedMessage, privateKey)
}

func hashMessage(data []byte) common.Hash {
	hexData := hexutil.Encode(data)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hexData), hexData)

	return crypto.Keccak256Hash([]byte(msg))
}

func signMessage(msg common.Hash, privateKey *ecdsa.PrivateKey) (string, error) {
	const op = "imx.signMessage"

	signatureBytes, err := crypto.Sign(msg.Bytes(), privateKey)
	if err != nil {
		return "", ez.Wrap(op, err)
	}

	return hexutil.Encode(signatureBytes), nil
}
