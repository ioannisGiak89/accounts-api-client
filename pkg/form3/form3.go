package form3

import (
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/factory"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"net/url"
)

type FormResources struct {
	Accounts accounts.Form3Accounts
}

func New(bu url.URL) *FormResources {
	libFactory := factory.NewForm3Lib()
	httpClient := libFactory.BuildForm3Client(bu)
	accountsService := libFactory.BuildAccountsService(httpClient)

	return &FormResources{
		Accounts: accountsService,
	}
}
