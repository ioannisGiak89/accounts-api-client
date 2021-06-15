package client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Form3ResourcesClient defines the Form3 Resources client
type Form3ResourcesClient interface {
	Get(path string) ([]byte, error)
	Delete(path string) error
	Post(path string, body []byte) ([]byte, error)
}

// HTTPClient interface. This interface is implemented by http.Client and is used for mocking
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Form3RestClient implements the Form3ResourcesClient interface
type Form3RestClient struct {
	baseUrl *url.URL
	client  HTTPClient
}

// Creates a new Form3 rest client
func NewForm3RestClient(baseUrl *url.URL, httpClient HTTPClient) *Form3RestClient {
	return &Form3RestClient{
		baseUrl: baseUrl,
		client:  httpClient,
	}
}

// Get does a get request to an endpoint
func (cl *Form3RestClient) Get(path string) ([]byte, error) {
	res, err := cl.createAndDoRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	resBody, err := cl.readResponseBody(res)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(resBody))
	}

	return resBody, nil
}

// Post does a post request to an endpoint
func (cl *Form3RestClient) Post(path string, body []byte) ([]byte, error) {
	res, err := cl.createAndDoRequest(http.MethodPost, path, body)

	if err != nil {
		return nil, err
	}

	resBody, err := cl.readResponseBody(res)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(string(resBody))
	}

	return resBody, nil
}

// Delete does a delete request to an endpoint
func (cl *Form3RestClient) Delete(path string) error {
	res, err := cl.createAndDoRequest(http.MethodDelete, path, nil)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		resBody, err := cl.readResponseBody(res)

		if err != nil {
			return err
		}

		return errors.New(string(resBody))
	}

	return nil
}

// Private method that creates and does the request. Used to avoid code duplication
func (cl *Form3RestClient) createAndDoRequest(httpMethod string, path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, fmt.Sprintf(
		"%s%s",
		cl.baseUrl.String(),
		path,
	), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return cl.client.Do(req)
}

// Private method that reads the response body. Used to avoid code duplication
func (cl *Form3RestClient) readResponseBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
