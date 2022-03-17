package client

type optionApplyFunc func(client *Client) error

type Option interface {
	applyOption(client *Client) error
}

func (f optionApplyFunc) applyOption(c *Client) error {
	return f(c)
}

func SetIMXSDKServiceParams(params SDKServiceParams) Option {
	return optionApplyFunc(func(client *Client) error {
		client.sdkServiceParams = params
		return nil
	})
}
