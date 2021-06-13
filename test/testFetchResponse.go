package test

import (
	"github.com/ioannisGiak89/accounts-api-client/pkg/model"
	"github.com/ioannisGiak89/accounts-api-client/utils"
)

// Returns a Fetch Account response
func GetAccountFetchResponse() *model.AccountFetchResponse {
	id := utils.ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb")
	oID := utils.ParseUuid("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")

	return &model.AccountFetchResponse{
		Data: model.Account{
			Attributes: model.AccountAttributes{
				AlternativeNames: []string{"Test", "Alt", "Names"},
				BankID:           "400300",
				BankIDCode:       "GBDSC",
				BaseCurrency:     "GBP",
				Bic:              "NWBKGB22",
				Country:          "GB",
				Name:             []string{"Samantha Holder"},
			},
			ID:             id,
			OrganisationID: oID,
			Version:        0,
			Type:           "accounts",
			CreatedOn:      "2021-06-12T13:30:28.831Z",
			ModifiedOn:     "2021-06-12T13:30:28.831Z",
		},
		Links: model.Links{Self: "/v1/organisation/accounts/9ea9bb7c-b5ec-4b00-bd82-af0067c4febb"},
	}
}
