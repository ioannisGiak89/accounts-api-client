package form3

import (
	"github.com/google/uuid"
	"github.com/ioannisGiak89/accounts-api-client/testUtils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestFrom3_New(t *testing.T) {
	t.Run("should do all three basic operations", func(t *testing.T) {
		baseURL, err := url.Parse("http://localhost:8080/")
		require.NoError(t, err)

		accountID := uuid.New()
		accountToCreate := testUtils.GetAccountCreateRequest(accountID)

		f3 := New(baseURL)

		accountApiResponse, err := f3.Accounts.Create(accountToCreate)
		assert.Nil(t, err)
		assert.Equal(t, accountToCreate.Data.ID, accountApiResponse.Data.ID)
		assert.Equal(t, accountToCreate.Data.Attributes.Country, accountApiResponse.Data.Attributes.Country)

		fetchResponse, err := f3.Accounts.Fetch(accountID)
		assert.Nil(t, err)
		assert.Equal(t, accountToCreate.Data.ID, fetchResponse.Data.ID)
		assert.Equal(t, accountToCreate.Data.Attributes.Country, fetchResponse.Data.Attributes.Country)

		err = f3.Accounts.Delete(accountID, 0)
		assert.Nil(t, err)
	})

	t.Run("should return errors", func(t *testing.T) {
		baseURL, err := url.Parse("http://localhost:8080/")
		require.NoError(t, err)

		accountID := uuid.New()
		accountToCreate := testUtils.GetAccountCreateRequest(accountID)
		accountToCreate.Data.Attributes.Country = ""
		f3 := New(baseURL)

		_, err = f3.Accounts.Create(accountToCreate)
		assert.NotNil(t, err)

		_, err = f3.Accounts.Fetch(accountID)
		assert.NotNil(t, err)

		err = f3.Accounts.Delete(accountID, 0)
		assert.NotNil(t, err)
	})
}
