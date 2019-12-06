package bca

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client struct
//
// BcaGetTokenURL, BcaAccStatementsURL can be empty as it will have default value when you call the related API
type Client struct {
	BcaBaseURL          string
	BcaClientID         string
	BcaClientSecret     string
	BcaApiKey           string
	BcaApiSecret        string
	BcaCompanyID        string
	BcaGetTokenURL      string
	BcaAccStatementsURL string
	Origin              string
}

// NewClient : this function will always be called when the library is in use
func NewClient() Client {
	return Client{}
}

// ===================== HTTP CLIENT ================================================
var defHTTPTimeout = 15 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, errors.New("sangu-bca.client.NewRequest: " + err.Error())
	}

	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}

	// if token request, set basic auth header
	if strings.Contains(fullPath, c.BcaGetTokenURL) {
		req.SetBasicAuth(c.BcaClientID, c.BcaClientSecret)
	}

	return req, nil
}

// ExecuteRequest : execute request
func (c *Client) ExecuteRequest(req *http.Request, v interface{}, x interface{}) error {
	res, err := httpClient.Do(req)
	if err != nil {
		return errors.New("sangu-bca.client.ExecuteRequest.Do: " + err.Error())
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("sangu-bca.client.ExecuteRequest.Read: " + err.Error())
	}

	if v != nil && res.StatusCode == 200 {
		if err = json.Unmarshal(resBody, v); err != nil {
			return errors.New("sangu-bca.client.ExecuteRequest.UnmarshalOK: " + err.Error())
		}
	}

	if x != nil && res.StatusCode != 200 {
		if err = json.Unmarshal(resBody, x); err != nil {
			return errors.New("sangu-bca.client.ExecuteRequest.UnmarshalNotOK: " + err.Error())
		}
	}

	return nil
}

// Call the BCA API at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occurred, the return of this function is the error
// itself, otherwise nil.
func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v interface{}, x interface{}) error {
	req, err := c.NewRequest(method, path, header, body)
	if err != nil {
		return err
	}

	return c.ExecuteRequest(req, v, x)
}

// ===================== END HTTP CLIENT ================================================
