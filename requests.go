package imx

import (
	"bytes"
	"encoding/json"
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

func (c *Client) httpRequest(method, path string, data url.Values, body, responseType interface{}) error {
	op := "imx.Client.httpRequest"
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

	URL := c.BaseURL + path

	request, err := http.NewRequest(method, URL+"?"+data.Encode(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return ez.Wrap(op, err)
	}

	// request.Header.Add("auth-token", c.Token)
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
