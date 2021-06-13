package model

import (
	"github.com/google/uuid"
)

type AccountFetchResponse struct {
	Data  Account
	Links Links
}

type Account struct {
	Attributes     AccountAttributes
	ID             uuid.UUID
	OrganisationID uuid.UUID `json:"organisation_id"`
	Version        int
	Type           string
	CreatedOn      string `json:"created_on"`
	ModifiedOn     string `json:"modified_on"`
}

type AccountAttributes struct {
	AlternativeNames []string `json:"alternative_names"`
	BankID           string   `json:"bank_id"`
	BankIDCode       string   `json:"bank_id_code"`
	BaseCurrency     string   `json:"base_currency"`
	Bic              string
	Country          string
	Name             []string
}

type Links struct {
	Self string
}
