package accounts

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/client"
	"github.com/ioannisGiak89/accounts-api-client/pkg/model"
)

// Defines the Accounts interface
type Form3Accounts interface {
	Fetch(uuid uuid.UUID) (*model.AccountApiResponse, error)
	Delete(accountID uuid.UUID, version int) error
	Create(account *model.AccountCreateRequest) (*model.AccountApiResponse, error)
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

func (f3a *Form3AccountsService) Fetch(accountID uuid.UUID) (*model.AccountApiResponse, error) {

	responseBody, err := f3a.Client.Get(accountID)

	if err != nil {
		return nil, err
	}

	var accountsResponse model.AccountApiResponse
	err = json.Unmarshal(responseBody, &accountsResponse)

	if err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}

func (f3a *Form3AccountsService) Delete(accountID uuid.UUID, version int) error {
	err := f3a.Client.Delete(accountID, version)
	return err
}

func (f3a *Form3AccountsService) Create(account *model.AccountCreateRequest) (*model.AccountApiResponse, error) {

	jsonBody, err := json.Marshal(account)

	if err != nil {
		return nil, err
	}

	responseBody, err := f3a.Client.Post(jsonBody)

	if err != nil {
		return nil, err
	}

	var accountsResponse model.AccountApiResponse
	err = json.Unmarshal(responseBody, &accountsResponse)

	if err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}
