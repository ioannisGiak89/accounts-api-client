package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AccountsApi interface {
	Get(path string) ([]byte, error)
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// AccountsRestClient
type AccountsRestClient struct {
	BaseUrl url.URL
	Client  HTTPClient
}

// Creates a new HTTP client
func NewAccountsRestClient(baseUrl url.URL) *AccountsRestClient {
	return &AccountsRestClient{
		BaseUrl: baseUrl,
		Client:  &http.Client{},
	}
}

// Get does a get request to an endpoint
func (cl *AccountsRestClient) Get(path string) ([]byte, error) {

	cl.BaseUrl.Path = path
	req, err := http.NewRequest(http.MethodGet, cl.BaseUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := cl.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	return body, nil
}
