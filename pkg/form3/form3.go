package form3

import (
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/factory"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"net/url"
)

// FormResources is a struct with all the available resources of the lib
type FormResources struct {
	Accounts accounts.Form3Accounts
}

// New creates and initialises a new Form3 client lib
func New(bu *url.URL) *FormResources {
	libFactory := factory.NewForm3LibFactory()
	httpClient := libFactory.BuildForm3Client(bu)
	accountsService := libFactory.BuildAccountsService(httpClient)

	return &FormResources{
		Accounts: accountsService,
	}
}
