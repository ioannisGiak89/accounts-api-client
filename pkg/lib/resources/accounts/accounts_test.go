package accounts_test

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"github.com/ioannisGiak89/accounts-api-client/test"
	"github.com/ioannisGiak89/accounts-api-client/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

type mockedHttpClient struct {
	BaseUrl    *url.URL
	MockGet    func(accountID uuid.UUID) ([]byte, error)
	MockDelete func(accountID uuid.UUID, version int) error
}

func (cl *mockedHttpClient) Delete(accountID uuid.UUID, version int) error {
	return cl.MockDelete(accountID, version)
}

func (cl *mockedHttpClient) Get(accountID uuid.UUID) ([]byte, error) {
	return cl.MockGet(accountID)
}

func TestForm3AccountsService_Get(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	accountID := utils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	require.NoError(t, err)

	t.Run("should return an AccountFetchResponse", func(t *testing.T) {
		expectedResponse := test.GetAccountFetchResponse()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: baseURL,
			MockGet: func(accountID uuid.UUID) ([]byte, error) {
				return jsonResponse, nil
			},
		})

		response, err := accountsService.Get(accountID)

		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: baseURL,
			MockGet: func(accountID uuid.UUID) ([]byte, error) {
				return nil, errors.New("there was an HTTP error")
			},
		})

		responseBody, err := accountsService.Get(accountID)

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("there was an HTTP error"), err)
	})

	t.Run("should return an error if the unmarshal fails", func(t *testing.T) {
		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: baseURL,
			MockGet: func(accountID uuid.UUID) ([]byte, error) {
				return []byte{12, 12}, nil
			},
		})

		response, err := accountsService.Get(accountID)

		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestForm3AccountsService_Delete(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	accountID := utils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	require.NoError(t, err)

	t.Run("should to a delete", func(t *testing.T) {
		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: baseURL,
			MockDelete: func(accountID uuid.UUID, version int) error {
				return nil
			},
		})

		err := accountsService.Delete(accountID, 0)

		assert.Nil(t, err)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: baseURL,
			MockDelete: func(accountID uuid.UUID, version int) error {
				return errors.New("there was an HTTP error")
			},
		})

		err := accountsService.Delete(accountID, 0)

		assert.Equal(t, errors.New("there was an HTTP error"), err)
	})
}
