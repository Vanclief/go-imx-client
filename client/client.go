package client

import (
	"net/http"

	"github.com/vanclief/ez"
)

type Client struct {
	http             *http.Client
	network          Network
	privateKey       string
	providerKey      string
	sdkEnabled       bool
	sdkServiceParams SDKServiceParams
}

type Network string

var (
	Mainnet Network = "mainnet"
	Ropsten Network = "ropsten"
)

func NewClient(network Network, privateKey, providerKey string, opts ...Option) (*Client, error) {
	const op = "imx.NewClient"

	sdkServiceParams := SDKServiceParams{
		Host: "http://localhost",
		Port: 4000,
	}

	if network != Mainnet && network != Ropsten {
		return nil, ez.New(op, ez.EINVALID, "Network not supported", nil)
	}

	c := &Client{
		http:             &http.Client{},
		network:          network,
		privateKey:       privateKey,
		providerKey:      providerKey,
		sdkServiceParams: sdkServiceParams,
	}

	for _, opt := range opts {
		if err := opt.applyOption(c); err != nil {
			return nil, ez.Wrap(op, err)
		}
	}

	return c, nil
}

func (c *Client) getBaseURL() string {
	if c.network == Mainnet {
		return MainnetURL
	}

	return RopstenURL
}

type SDKServiceParams struct {
	Host string
	Port int
}

type InitSDKRequest struct {
	Network    Network `json:"network"`
	PrivateKey string  `json:"private_key"`
	AlchemyKey string  `json:"alchemy_key"`
}

type InitSDKResponse struct {
	PublicApiUrl                string `json:"publicApiUrl"`
	ContractAddress             string `json:"contractAddress"`
	RegistrationContractAddress string `json:"registrationContractAddress"`
}

func (c *Client) initSDKClient() error {
	const op = "Client.initSDKClient"

	if c.sdkEnabled {
		return nil
	}

	request := InitSDKRequest{
		Network:    c.network,
		PrivateKey: "2772b0cdf316c9874a4e36873b08f46b05789f294286b2f9e13726612352c022",
		AlchemyKey: "DvukuyBzEK-JyP6zp1NVeNVYLJCrzjp_",
	}

	var response InitSDKResponse

	err := c.SDKRequest("POST", SDKInitURL, nil, request, &response)
	if err != nil {
		return ez.Wrap(op, err)
	}

	c.sdkEnabled = true

	return nil
}
