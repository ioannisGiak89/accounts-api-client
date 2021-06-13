package client_test

import (
	"bytes"
	"errors"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type mockedHttpClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

func (cl *mockedHttpClient) Do(req *http.Request) (*http.Response, error) {
	return cl.MockDo(req)
}

func TestHttpClient_Get(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	require.NoError(t, err)

	t.Run("should return an error if the request fails", func(t *testing.T) {
		form3Client := client.AccountsRestClient{
			BaseUrl: *baseURL,
			Client: &mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("network request failed")
				},
			},
		}

		responseBody, err := form3Client.Get("a/path/to/accounts")

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("network request failed"), err)
	})

	t.Run("should return an error if status code wasn't 200", func(t *testing.T) {
		form3Client := client.AccountsRestClient{
			BaseUrl: *baseURL,
			Client: &mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("not found"))),
						StatusCode: 404,
					}, nil
				},
			},
		}

		responseBody, err := form3Client.Get("a/path/to/accounts")

		assert.Equal(t, errors.New("not found"), err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return the responseBody", func(t *testing.T) {
		form3Client := client.AccountsRestClient{
			BaseUrl: *baseURL,
			Client: &mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("A valid account"))),
						StatusCode: 200,
					}, nil
				},
			},
		}

		responseBody, err := form3Client.Get("a/path/to/accounts")

		assert.Equal(t, "A valid account", string(responseBody))
		assert.Nil(t, err)
	})
}
