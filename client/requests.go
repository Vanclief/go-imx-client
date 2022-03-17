package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/vanclief/ez"
)

type IMXResponse struct {
	Payload interface{}
}

type IMXError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) RestRequest(method, path string, data url.Values, body, responseType interface{}) error {
	URL := c.getBaseURL() + path

	return c.httpRequest(method, URL, data, body, responseType)
}

func (c *Client) SDKRequest(method, path string, data url.Values, body, responseType interface{}) error {

	if path != SDKInitURL {
		c.initSDKClient()
	}

	URL := fmt.Sprintf("%s:%d/%s",
		c.sdkServiceParams.Host,
		c.sdkServiceParams.Port,
		path,
	)

	return c.httpRequest(method, URL, data, body, responseType)
}

func (c *Client) httpRequest(method, URL string, data url.Values, body, responseType interface{}) error {
	op := "Client.httpRequest"

	if data == nil {
		data = url.Values{}
	}

	var jsonBody []byte
	var err error

	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return ez.Wrap(op, err)
		}
	}

	request, err := http.NewRequest(method, URL+"?"+data.Encode(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return ez.Wrap(op, err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err := c.http.Do(request)
	if err != nil {
		return ez.Wrap(op, err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ez.Wrap(op, err)
	}

	if response.StatusCode != 200 {
		errorType := ez.HTTPStatusToError(response.StatusCode)

		apiError := &IMXError{}
		err = json.Unmarshal(responseBody, apiError)
		if err != nil {
			return ez.Wrap(op, err)
		}

		return ez.New(op, errorType, apiError.Message, nil)
	}

	apiResponse := &IMXResponse{}
	if responseType != nil {
		apiResponse.Payload = responseType
	}

	err = json.Unmarshal(responseBody, apiResponse.Payload)
	if err != nil {
		return ez.New(op, ez.EINVALID, err.Error(), err)
	}

	defer response.Body.Close()

	return nil
}
