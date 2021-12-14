package imx

import (
	"net/http"

	"github.com/vanclief/ez"
)

type Client struct {
	http    *http.Client
	BaseURL string
}
type Network string

var (
	Mainnet Network = "mainnet"
	Ropsten Network = "ropsten"
)

func NewClient(network Network) (*Client, error) {
	const op = "imx.NewClient"

	var baseURL string

	switch network {
	case Mainnet:
		baseURL = "https://api.x.immutable.com/v1/"

	case Ropsten:
		baseURL = "https://api.ropsten.x.immutable.com/v1/"

	default:
		return nil, ez.New(op, ez.EINVALID, "Network not supported", nil)
	}

	client := &Client{
		http:    &http.Client{},
		BaseURL: baseURL,
	}

	return client, nil
}
