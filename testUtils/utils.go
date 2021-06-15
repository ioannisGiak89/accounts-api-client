package testUtils

import (
	"github.com/google/uuid"
	"github.com/ioannisGiak89/accounts-api-client/pkg/model"
	"log"
)

// Returns a Form3 Accounts API response
func GetAccountApiResponse() *model.AccountApiResponse {
	return &model.AccountApiResponse{
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
			ID:             ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb"),
			OrganisationID: ParseUuid("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"),
			Version:        0,
			Type:           "accounts",
			CreatedOn:      "2021-06-12T13:30:28.831Z",
			ModifiedOn:     "2021-06-12T13:30:28.831Z",
		},
		Links: model.Links{Self: "/v1/organisation/accounts/9ea9bb7c-b5ec-4b00-bd82-af0067c4febb"},
	}
}

// Returns a Form3 Accounts API create request
func GetAccountCreateRequest() *model.AccountCreateRequest {
	return &model.AccountCreateRequest{
		Data: model.Account{
			Attributes: model.AccountAttributes{
				AlternativeNames: nil,
				BankID:           "400300",
				BankIDCode:       "GBDSC",
				BaseCurrency:     "GBP",
				Bic:              "NWBKGB22",
				Country:          "GB",
				Name:             []string{"Samantha Holder2"},
			},
			ID:             ParseUuid("9ea9bb7c-b5ec-4b00-bd82-af0067c4febb"),
			OrganisationID: ParseUuid("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"),
			Version:        0,
			Type:           "accounts",
		},
	}
}

func ParseUuid(id string) uuid.UUID {
	uID, err := uuid.Parse(id)

	if err != nil {
		log.Fatal(err)
	}

	return uID
}
