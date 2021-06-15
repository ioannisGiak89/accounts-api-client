package accounts

import (
	"encoding/json"
	"fmt"
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

// Form3AccountsService implements the Accounts interface. This service is meant to be the lib API
// and handle any logic around Accounts
type Form3AccountsService struct {
	Client           client.Form3ResourcesClient
	accountsEndpoint string
}

// NewForm3AccountsService creates a Form3AccountsService
func NewForm3AccountsService(cl client.Form3ResourcesClient, ae string) *Form3AccountsService {
	return &Form3AccountsService{
		Client:           cl,
		accountsEndpoint: ae,
	}
}

// Fetch is used to retrieve Form3 Accounts
func (f3a *Form3AccountsService) Fetch(accountID uuid.UUID) (*model.AccountApiResponse, error) {
	path := fmt.Sprintf(
		"%s%s",
		f3a.accountsEndpoint,
		accountID.String(),
	)
	responseBody, err := f3a.Client.Get(path)

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

// Delete is used to delete Form3 Accounts
func (f3a *Form3AccountsService) Delete(accountID uuid.UUID, version int) error {
	path := fmt.Sprintf(
		"%s%s?version=%b",
		f3a.accountsEndpoint,
		accountID.String(),
		version,
	)
	err := f3a.Client.Delete(path)

	return err
}

// Create is used to create Form3 Accounts
func (f3a *Form3AccountsService) Create(account *model.AccountCreateRequest) (*model.AccountApiResponse, error) {
	jsonBody, err := json.Marshal(account)

	if err != nil {
		return nil, err
	}

	responseBody, err := f3a.Client.Post(f3a.accountsEndpoint, jsonBody)

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
