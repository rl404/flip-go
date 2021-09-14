package flip

import (
	"context"
	"fmt"
	"net/http"
)

// Disbursement is disbursement transaction response model.
type Disbursement struct {
	ID            int               `json:"id"`
	UserID        int               `json:"user_id"`
	Amount        float64           `json:"amount"`
	Status        TransactionStatus `json:"status"`
	Timestamp     string            `json:"timestamp"` // yyyy-mm-dd hh:mm:ss GMT+7
	BankCode      BankCode          `json:"bank_code"`
	AccountNumber string            `json:"account_number"`
	RecipientName string            `json:"recipient_name"`
	SenderBank    string            `json:"sender_bank"`
	Remark        string            `json:"remark"`
	Receipt       string            `json:"receipt"`
	TimeServed    string            `json:"time_served"` // yyyy-mm-dd hh:mm:ss GMT+7
	BundleID      int               `json:"bundle_id"`
	CompanyID     int               `json:"company_id"`
	RecipientCity int               `json:"recipient_city"` // from GetCities()
	CreatedFrom   Source            `json:"created_from"`
	Direction     Direction         `json:"direction"`
	Sender        *Sender           `json:"sender"`
	Fee           float64           `json:"fee"`
}

// Sender is disbursement sender model.
type Sender struct {
	SenderName         string       `json:"sender_name"`
	PlaceOfBirth       int          `json:"place_of_birth"` // from GetCities()
	DateOfBirth        string       `json:"date_of_birth"`  // yyyy-mm-dd
	Address            string       `json:"address"`
	SenderIdentityType IdentityType `json:"sender_identity_type"`
	SenderCountry      int          `json:"sender_country"` // from GetCountries()
	Job                JobType      `json:"job"`
}

// CreateDisbursement to create disbursement.
func (c *Client) CreateDisbursement(request CreateDisbursementRequest) (*Disbursement, error) {
	return c.CreateDisbursementWithContext(context.Background(), request)
}

// CreateDisbursementWithContext to create disbursement with context.
func (c *Client) CreateDisbursementWithContext(ctx context.Context, request CreateDisbursementRequest) (*Disbursement, error) {
	if err := validate(&request); err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Add("idempotency-key", request.IdempotencyKey)

	var response Disbursement
	err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/disbursement", c.baseURL),
		c.secretKey,
		header,
		request.encode(),
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetDisbursement to get disbursement by id.
func (c *Client) GetDisbursement(id int) (*Disbursement, error) {
	return c.GetDisbursementWithContext(context.Background(), id)
}

// GetDisbursementWithContext to get disbursement by id with context.
func (c *Client) GetDisbursementWithContext(ctx context.Context, id int) (*Disbursement, error) {
	if err := validate(&getDisbursementRequest{ID: id}); err != nil {
		return nil, err
	}

	var response Disbursement
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursement/%d", c.baseURL, id),
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

// Disbursements is disbursement pagination response model.
type Disbursements struct {
	TotalData   int            `json:"total_data"`
	DataPerPage int            `json:"data_per_page"`
	TotalPage   int            `json:"total_page"`
	Page        int            `json:"page"`
	Data        []Disbursement `json:"data"`
}

// GetDisbursements to get disbursement list.
func (c *Client) GetDisbursements(request GetDisbursementsRequest) (*Disbursements, error) {
	return c.GetDisbursementsWithContext(context.Background(), request)
}

// GetDisbursementsWithContext to get disbursement list with context.
func (c *Client) GetDisbursementsWithContext(ctx context.Context, request GetDisbursementsRequest) (*Disbursements, error) {
	if err := validate(&request); err != nil {
		return nil, err
	}

	var response Disbursements
	err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursement?%s", c.baseURL, request.encode()),
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
