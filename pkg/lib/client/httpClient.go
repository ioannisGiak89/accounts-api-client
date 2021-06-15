package client

import (
	"bytes"
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
	Post(body []byte) ([]byte, error)
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// AccountsRestClient
type AccountsRestClient struct {
	baseUrl *url.URL
	client  HTTPClient
}

// Creates a new HTTP client
func NewAccountsRestClient(baseUrl *url.URL, httpClient HTTPClient) *AccountsRestClient {
	return &AccountsRestClient{
		baseUrl: baseUrl,
		client:  httpClient,
	}
}

// Fetch does a get request to an endpoint
func (cl *AccountsRestClient) Get(accountID uuid.UUID) ([]byte, error) {
	path := fmt.Sprintf(
		"%s%s%s",
		cl.baseUrl.String(),
		"v1/organisation/accounts/",
		accountID.String(),
	)

	req, err := http.NewRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := cl.client.Do(req)

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

// Fetch does a get request to an endpoint
func (cl *AccountsRestClient) Post(body []byte) ([]byte, error) {
	path := fmt.Sprintf(
		"%s%s",
		cl.baseUrl.String(),
		"v1/organisation/accounts/",
	)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(string(resBody))
	}

	return resBody, nil
}

func (cl *AccountsRestClient) Delete(accountID uuid.UUID, version int) error {
	path := fmt.Sprintf(
		"%s%s%s?version=%b",
		cl.baseUrl.String(),
		"v1/organisation/accounts/",
		accountID.String(),
		version,
	)
	req, err := http.NewRequest(http.MethodDelete, path, nil)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := cl.client.Do(req)

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
