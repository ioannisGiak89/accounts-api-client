package accounts

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/client"
	"github.com/ioannisGiak89/accounts-api-client/pkg/model"
)

// Defines the Accounts interface
type Form3Accounts interface {
	Get(uuid uuid.UUID) (*model.AccountFetchResponse, error)
	//Delete(uuid uuid.UUID) (*model.AccountFetchResponse, error)
}

// Form3AccountsService provides the Accounts API
type Form3AccountsService struct {
	Client client.AccountsApi
}

func New(cl client.AccountsApi) *Form3AccountsService {
	return &Form3AccountsService{
		Client: cl,
	}
}

func (f3a *Form3AccountsService) Get(accountID uuid.UUID) (*model.AccountFetchResponse, error) {

	responseBody, err := f3a.Client.Get("v1/organisation/accounts/" + accountID.String())

	if err != nil {
		return nil, err
	}

	var accountsResponse model.AccountFetchResponse
	err = json.Unmarshal(responseBody, &accountsResponse)

	if err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}
