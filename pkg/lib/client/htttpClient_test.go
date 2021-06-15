package client_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/client"
	"github.com/ioannisGiak89/accounts-api-client/testUtils"
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
	accountID := testUtils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	require.NoError(t, err)

	t.Run("should return an error if the request fails", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("network request failed")
				},
			},
		)

		responseBody, err := form3Client.Get(accountID)

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("network request failed"), err)
	})

	t.Run("should return an error if status code wasn't 200", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						// Return a
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("not found"))),
						StatusCode: http.StatusNotFound,
					}, nil
				},
			},
		)

		responseBody, err := form3Client.Get(accountID)

		assert.Equal(t, errors.New("not found"), err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return the responseBody", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("A valid account"))),
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		)

		responseBody, err := form3Client.Get(accountID)

		assert.Equal(t, "A valid account", string(responseBody))
		assert.Nil(t, err)
	})
}

func TestHttpClient_Post(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	require.NoError(t, err)
	accountToCreate := testUtils.GetAccountCreateRequest()
	bodyRequest, err := json.Marshal(accountToCreate)
	require.NoError(t, err)

	t.Run("should return an error if the request fails", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("network request failed")
				},
			},
		)

		responseBody, err := form3Client.Post(bodyRequest)

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("network request failed"), err)
	})

	t.Run("should return an error if status code wasn't 201", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						// Return a
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("conflict"))),
						StatusCode: http.StatusConflict,
					}, nil
				},
			},
		)

		responseBody, err := form3Client.Post(bodyRequest)

		assert.Equal(t, errors.New("conflict"), err)
		assert.Nil(t, responseBody)
	})

	t.Run("should return the responseBody", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("Account Created"))),
						StatusCode: http.StatusCreated,
					}, nil
				},
			},
		)

		responseBody, err := form3Client.Post(bodyRequest)

		assert.Equal(t, "Account Created", string(responseBody))
		assert.Nil(t, err)
	})
}

func TestHttpClient_Delete(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	accountID := testUtils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	require.NoError(t, err)

	t.Run("should return an error if the request fails", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("network request failed")
				},
			},
		)

		err := form3Client.Delete(accountID, 0)

		assert.Equal(t, errors.New("network request failed"), err)
	})

	t.Run("should return an error if status code wasn't 204", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("not found"))),
						StatusCode: http.StatusNotFound,
					}, nil
				},
			},
		)

		err := form3Client.Delete(accountID, 0)

		assert.Equal(t, errors.New("not found"), err)
	})

	t.Run("should nil if there is no error", func(t *testing.T) {
		form3Client := client.NewAccountsRestClient(
			baseURL,
			&mockedHttpClient{
				MockDo: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNoContent,
					}, nil
				},
			},
		)

		err := form3Client.Delete(accountID, 0)

		assert.Nil(t, err)
	})
}
