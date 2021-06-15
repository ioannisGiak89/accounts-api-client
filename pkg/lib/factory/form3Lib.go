package factory

import (
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/client"
	"github.com/ioannisGiak89/accounts-api-client/pkg/lib/resources/accounts"
	"net/http"
	"net/url"
)

// StandardFactory abstracts the creation of instances.
type StandardFactory interface {
	BuildAccountsService(client.Form3ResourcesClient) accounts.Form3AccountsService
	BuildForm3Client(baseUrl url.URL) client.Form3ResourcesClient
}

// Form3LibFactory builds instances
type Form3LibFactory struct{}

// NewForm3LibFactory creates a Form3LibFactory
func NewForm3LibFactory() *Form3LibFactory {
	return &Form3LibFactory{}
}

// BuildAccountsService builds a NewForm3AccountsService
func (f *Form3LibFactory) BuildAccountsService(cl client.Form3ResourcesClient) accounts.Form3Accounts {
	return accounts.NewForm3AccountsService(cl, "v1/organisation/accounts/")
}

// BuildForm3Client build a NewForm3RestClient
func (f *Form3LibFactory) BuildForm3Client(baseUrl *url.URL) client.Form3ResourcesClient {
	return client.NewForm3RestClient(baseUrl, &http.Client{})
}
