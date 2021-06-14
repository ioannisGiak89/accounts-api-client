package factory

import (
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/client"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"net/url"
)

type StandardFactory interface {
	BuildAccountsService(client.AccountsApi) accounts.Form3AccountsService
	BuildForm3Client(baseUrl url.URL) client.AccountsApi
}

type Form3LibFactory struct{}

func NewForm3Lib() *Form3LibFactory {
	return &Form3LibFactory{}
}

func (f *Form3LibFactory) BuildAccountsService(cl client.AccountsApi) accounts.Form3Accounts {
	return accounts.New(cl)
}

func (f *Form3LibFactory) BuildForm3Client(baseUrl *url.URL) client.AccountsApi {
	return client.NewAccountsRestClient(baseUrl)
}
