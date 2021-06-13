package accounts_test

import (
	"encoding/json"
	"errors"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"github.com/ioannisGiak89/accounts-api-client/test"
	"github.com/ioannisGiak89/accounts-api-client/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

type mockedHttpClient struct {
	BaseUrl url.URL
	MockGet func(path string) ([]byte, error)
}

func (cl *mockedHttpClient) Get(path string) ([]byte, error) {
	return cl.MockGet(path)
}

func TestForm3AccountsService_Get(t *testing.T) {

	baseURL, err := url.Parse("http://localhost:8080/")
	require.NoError(t, err)

	t.Run("should return an AccountFetchResponse", func(t *testing.T) {
		accountID := utils.ParseUuid("7b524f63-fd75-443f-9c27-6c560920ea75")
		expectedResponse := test.GetAccountFetchResponse()
		jsonResponse, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: *baseURL,
			MockGet: func(path string) ([]byte, error) {
				return jsonResponse, nil
			},
		})

		response, err := accountsService.Get(accountID)

		assert.Nil(t, err)
		assert.Equal(t, response, expectedResponse)
	})

	t.Run("should return an error if the client fails", func(t *testing.T) {
		accountID := utils.ParseUuid("7b524f63-fd75-443f-9c27-6c560920ea73")
		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: *baseURL,
			MockGet: func(path string) ([]byte, error) {
				return nil, errors.New("there was an HTTP error")
			},
		})

		responseBody, err := accountsService.Get(accountID)

		assert.Nil(t, responseBody)
		assert.Equal(t, errors.New("there was an HTTP error"), err)
	})

	t.Run("should return an error if the unmarshal fails", func(t *testing.T) {
		accountID := utils.ParseUuid("7b524f63-fd75-443f-9c27-6c560920ea74")
		accountsService := accounts.New(&mockedHttpClient{
			BaseUrl: *baseURL,
			MockGet: func(path string) ([]byte, error) {
				return []byte{12, 12}, nil
			},
		})

		response, err := accountsService.Get(accountID)

		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}
