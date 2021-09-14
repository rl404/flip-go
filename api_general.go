package flip

import (
	"context"
	"fmt"
	"net/http"
)

// Balance is response model from get balance.
type Balance struct {
	Balance int `json:"balance"`
}

// GetBalance to get account balance.
func (c *Client) GetBalance() (*Balance, error) {
	return c.GetBalanceWithContext(context.Background())
}

// GetBalanceWithContext to get account balance with context.
func (c *Client) GetBalanceWithContext(ctx context.Context) (*Balance, error) {
	var response Balance
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/general/balance", c.baseURL),
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Bank is response model from get banks.
type Bank struct {
	BankCode BankCode   `json:"bank_code"`
	Name     string     `json:"name"`
	Fee      float64    `json:"fee"`
	Queue    int        `json:"queue"`
	Status   BankStatus `json:"status"`
}

// GetBanks to get list of bank info.
func (c *Client) GetBanks(bankCode ...string) ([]Bank, error) {
	return c.GetBanksWithContext(context.Background(), bankCode...)
}

// GetBanksWithContext to get list of bank info with context.
func (c *Client) GetBanksWithContext(ctx context.Context, bankCode ...string) ([]Bank, error) {
	var queryString string
	if len(bankCode) > 0 {
		queryString = fmt.Sprintf("?code=%s", bankCode[0])
	}

	var response []Bank
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/general/banks%s", c.baseURL, queryString),
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Maintenance is response model from is maintenance.
type Maintenance struct {
	Maintenance bool `json:"maintenance"`
}

// IsMaintenance to get flip maintenance status.
func (c *Client) IsMaintenance() (*Maintenance, error) {
	return c.IsMaintenanceWithContext(context.Background())
}

// IsMaintenanceWithContext to get flip maintenance status with context.
func (c *Client) IsMaintenanceWithContext(ctx context.Context) (*Maintenance, error) {
	var response Maintenance
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/general/maintenance", c.baseURL),
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// BankAccount is response model from inquiry bank account.
type BankAccount struct {
	BankCode      BankCode      `json:"bank_code"`
	AccountNumber string        `json:"account_number"`
	AccountHolder string        `json:"account_holder"`
	Status        AccountStatus `json:"status"`
}

// InquiryBankAccount to inquiry bank account.
func (c *Client) InquiryBankAccount(request InquiryBankAccountRequest) (*BankAccount, error) {
	return c.InquiryBankAccountWithContext(context.Background(), request)
}

// InquiryBankAccountWithContext to inquiry bank account with context.
func (c *Client) InquiryBankAccountWithContext(ctx context.Context, request InquiryBankAccountRequest) (*BankAccount, error) {
	if err := validate(&request); err != nil {
		return nil, err
	}

	var response BankAccount
	err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/disbursement/bank-account-inquiry", c.baseURL),
		c.secretKey,
		nil,
		request.encode(),
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// City is response model for city.
type City map[int]string

// GetCities to get city list
func (c *Client) GetCities() (*City, error) {
	return c.GetCitiesWithContext(context.Background())
}

// GetCitiesWithContext to get city list with context.
func (c *Client) GetCitiesWithContext(ctx context.Context) (*City, error) {
	var response City
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursement/city-list", c.baseURL),
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Country is response model for country.
type Country map[int]string

// GetCountries to get country list
func (c *Client) GetCountries() (*Country, error) {
	return c.GetCountriesWithContext(context.Background())
}

// GetCountriesWithContext to get country list with context.
func (c *Client) GetCountriesWithContext(ctx context.Context) (*Country, error) {
	var response Country
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursement/country-list", c.baseURL),
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CityCountry is response model for city and country.
type CityCountry map[int]string

// GetCitiesCountries to get city and country list
func (c *Client) GetCitiesCountries() (*CityCountry, error) {
	return c.GetCitiesCountriesWithContext(context.Background())
}

// GetCitiesCountriesWithContext to get city and country list with context.
func (c *Client) GetCitiesCountriesWithContext(ctx context.Context) (*CityCountry, error) {
	var response CityCountry
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursement/city-country-list", c.baseURL),
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
