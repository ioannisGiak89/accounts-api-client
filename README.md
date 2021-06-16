
# Account API Client Library

A client library in Go to access Form3's fake API



## Usage/Examples

```go
baseURL, err := url.Parse("http://localhost:8080/")

if err != nil {
    log.Fatal(err)
}

f3 := form3.New(baseURL)
accountID := uuid.New()

accountToCreate := &model.AccountCreateRequest{
    Data: model.Account{
        Attributes:     model.AccountAttributes{
            AlternativeNames: nil,
            BankID:           "400300",
            BankIDCode:       "GBDSC",
            BaseCurrency:     "GBP",
            Bic:              "NWBKGB22",
            Country:          "GB",
            Name:             []string{"Samantha Holder2"},
        },
        ID:             accountID,
        OrganisationID: uuid.New(),
        Version:        0,
        Type:           "accounts",
    },
}

accountApiResponse, err := f3.Accounts.Create(accountToCreate)

if err != nil {
    log.Fatal(err)
}

account, err := f3.Accounts.Fetch(accountID)

if err != nil {
    log.Fatal(err)
}

err = f3.Accounts.Delete(accountID, 0)

if err != nil {
    log.Fatal(err)
}
```

  
## API Reference

### Accounts

#### `Fetch(accountID uuid.UUID) (*model.AccountApiResponse, error)`

Takes an account ID and returns the Form3's API fetch response or an error.

#### `Delete(accountID uuid.UUID) error`

Takes an account ID and deletes an account. Returns an error if something foes wrong.

#### `Create(account *model.AccountCreateRequest) (*model.AccountApiResponse, error)`

Takes an AccountCreateRequest and creates an account. Returns the Form3's API response or an error.


  
## Run Locally

Clone the project

```bash
  git clone git@github.com:ioannisGiak89/accounts-api-client.git
```

Go to the project directory

```bash
  cd accounts-api-client
```

Spin up docker containers

**Note** This will also run the tests on the start up. Integration tests are making calls to the fake API. To run the tests outside the container please see the [Running Tests](#Running-Tests) section bellow

```bash
  docker-compose up
```


  
## Running Tests

To run tests, run the following command

```bash
  go test ./...
```

This will run both unit and integration tests. 

By default, integrations tests have been configured to run from within the lib container and make calls to the fake API.

To run the tests from your host machine, change the var baseUrl to http://localhost:8080/ in form3Integration_test.go file.
## Future Improvments

* Suport configuration as an object.
* Suport configuration as env variables.
* Cache API responses to avoid multiple calls to the API within sort period of time.
* Add support for other resources rather than accounts.
