package client

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AccountsApi interface {
	Get(accountID uuid.UUID) ([]byte, error)
	Delete(accountID uuid.UUID, version int) error
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// AccountsRestClient
type AccountsRestClient struct {
	BaseUrl *url.URL
	Client  HTTPClient
}

// Creates a new HTTP client
func NewAccountsRestClient(baseUrl *url.URL) *AccountsRestClient {
	return &AccountsRestClient{
		BaseUrl: baseUrl,
		Client:  &http.Client{},
	}
}

// Get does a get request to an endpoint
func (cl *AccountsRestClient) Get(accountID uuid.UUID) ([]byte, error) {
	path := fmt.Sprintf(
		"%s%s%s",
		cl.BaseUrl.String(),
		"v1/organisation/accounts/",
		accountID.String(),
	)

	req, err := http.NewRequest(http.MethodGet, path, nil)

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

func (cl *AccountsRestClient) Delete(accountID uuid.UUID, version int) error {
	path := fmt.Sprintf(
		"%s%s%s?version=%b",
		cl.BaseUrl.String(),
		"v1/organisation/accounts/",
		accountID.String(),
		version,
	)
	req, err := http.NewRequest(http.MethodDelete, path, nil)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := cl.Client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
