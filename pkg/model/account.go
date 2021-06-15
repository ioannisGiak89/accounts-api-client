package model

import (
	"github.com/google/uuid"
)

// AccountApiResponse struct represents the response from Form3 Accounts API
type AccountApiResponse struct {
	Data  Account
	Links Links
}

// AccountCreateRequest struct represents the request send to Form3 Accounts API to create an account
type AccountCreateRequest struct {
	Data Account
}

// Account struct represents a Form3 Account
type Account struct {
	Attributes     AccountAttributes
	ID             uuid.UUID
	OrganisationID uuid.UUID `json:"organisation_id"`
	Version        int
	Type           string
	CreatedOn      string `json:"created_on"`
	ModifiedOn     string `json:"modified_on"`
}

// AccountAttributes struct represents the attributes of a Form3 Account
type AccountAttributes struct {
	AlternativeNames []string `json:"alternative_names"`
	BankID           string   `json:"bank_id"`
	BankIDCode       string   `json:"bank_id_code"`
	BaseCurrency     string   `json:"base_currency"`
	Bic              string
	Country          string
	Name             []string
}

// Links struct represents the links included in a Form3 Accounts API response
type Links struct {
	Self string
}
