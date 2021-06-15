package accounts_test

import (
	"encoding/json"
	"errors"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"github.com/ioannisGiak89/accounts-api-client/testUtils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

// Implements Form3ResourcesClient interface. This struct is used to mock the Form3RestClient
type mockedHttpClient struct {
	BaseUrl    *url.URL
	MockGet    func(path string) ([]byte, error)
	MockDelete func(path string) error
	MockPost   func(path string, body []byte) ([]byte, error)
}

func (cl *mockedHttpClient) Post(path string, body []byte) ([]byte, error) {
	return cl.MockPost(path, body)
}

func (cl *mockedHttpClient) Delete(path string) error {
	return cl.MockDelete(path)
}

func (cl *mockedHttpClient) Get(path string) ([]byte, error) {
	return cl.MockGet(path)
}

func TestForm3AccountsService_Fetch(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	accountID := testUtils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	require.NoError(t, err)

	t.Run("should return an AccountApiResponse", func(t *testing.T) {
		expectedResponse := testUtils.GetAccountApiResponse()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockGet: func(path string) ([]byte, error) {
				return jsonResponse, nil
			},
		}, "path/to/accounts/endpoint")

		response, err := accountsService.Fetch(accountID)

		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockGet: func(path string) ([]byte, error) {
				return nil, errors.New("there was an HTTP error")
			},
		}, "path/to/accounts/endpoint")

		responseBody, err := accountsService.Fetch(accountID)

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("there was an HTTP error"), err)
	})

	t.Run("should return an error if the unmarshal fails", func(t *testing.T) {
		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockGet: func(path string) ([]byte, error) {
				return []byte{12, 12}, nil
			},
		}, "path/to/accounts/endpoint")

		response, err := accountsService.Fetch(accountID)

		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestForm3AccountsService_Delete(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	accountID := testUtils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	require.NoError(t, err)

	t.Run("should to a delete", func(t *testing.T) {
		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockDelete: func(path string) error {
				return nil
			},
		}, "path/to/accounts/endpoint")

		err := accountsService.Delete(accountID, 0)

		assert.Nil(t, err)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockDelete: func(path string) error {
				return errors.New("there was an HTTP error")
			},
		}, "path/to/accounts/endpoint")

		err := accountsService.Delete(accountID, 0)

		assert.Equal(t, errors.New("there was an HTTP error"), err)
	})
}

func TestForm3AccountsService_Create(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	require.NoError(t, err)
	accountToCreate := testUtils.GetAccountCreateRequest()

	t.Run("should create an account and return an AccountApiResponse", func(t *testing.T) {
		expectedResponse := testUtils.GetAccountApiResponse()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockPost: func(path string, body []byte) ([]byte, error) {
				return jsonResponse, nil
			},
		}, "path/to/accounts/endpoint")

		response, err := accountsService.Create(accountToCreate)

		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
		assert.Equal(t, accountToCreate.Data.ID, expectedResponse.Data.ID)
		assert.Equal(t, accountToCreate.Data.Attributes.Country, expectedResponse.Data.Attributes.Country)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockPost: func(path string, body []byte) ([]byte, error) {
				return nil, errors.New("there was an HTTP error")
			},
		}, "path/to/accounts/endpoint")

		responseBody, err := accountsService.Create(accountToCreate)

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("there was an HTTP error"), err)
	})

	t.Run("should return an error if the unmarshal fails", func(t *testing.T) {
		accountsService := accounts.NewForm3AccountsService(&mockedHttpClient{
			BaseUrl: baseURL,
			MockPost: func(path string, body []byte) ([]byte, error) {
				return []byte{12, 12}, nil
			},
		}, "path/to/accounts/endpoint")

		response, err := accountsService.Create(accountToCreate)

		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}
