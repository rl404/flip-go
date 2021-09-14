package flip

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// InquiryBankAccountRequest is request model for inquiry bank account.
type InquiryBankAccountRequest struct {
	AccountNumber string   `validate:"required" mod:"no_space"`
	BankCode      BankCode `validate:"required" mod:"no_space"`
}

func (i *InquiryBankAccountRequest) encode() io.Reader {
	data := url.Values{}
	data.Set("account_number", i.AccountNumber)
	data.Set("bank_code", string(i.BankCode))
	return strings.NewReader(data.Encode())
}

// CreateDisbursementRequest is request model to create disbursement.
type CreateDisbursementRequest struct {
	IdempotencyKey string   `validate:"required"`
	AccountNumber  string   `validate:"required"`
	BankCode       BankCode `validate:"required"`
	Amount         float64  `validate:"required,gt=0"`
	Remark         string   `validate:"max=18"`
	RecipientCity  int      // from GetCities()
}

func (c *CreateDisbursementRequest) encode() io.Reader {
	data := url.Values{}
	data.Set("account_number", c.AccountNumber)
	data.Set("bank_code", string(c.BankCode))
	data.Set("amount", fmt.Sprintf("%g", c.Amount))

	if c.Remark != "" {
		data.Set("remark", c.Remark)
	}

	if c.RecipientCity > 0 {
		data.Set("recipient_city", strconv.Itoa(c.RecipientCity))
	}

	return strings.NewReader(data.Encode())
}

type getDisbursementRequest struct {
	ID int `validate:"required,gt=0"`
}

// GetDisbursementsRequest is request model to get disbursement list.
type GetDisbursementsRequest struct {
	Pagination int `validate:"gt=0" mod:"default=20"`
	Page       int `validate:"gt=0" mod:"default=1"`
	Sort       Sort

	// Additional attributes.
	ID            int               // exact
	Amount        float64           // exact
	Status        TransactionStatus // exact
	Timestamp     string            // like (yyyy-mm-dd hh:mm:ss)
	BankCode      BankCode          // like
	AccountNumber string            // like
	RecipientName string            // like
	Remark        string            // like
	TimeServed    string            // like (yyyy-mm-dd hh:mm:ss)
	CreatedFrom   Source            // like
	Direction     Direction         // exact
}

func (g *GetDisbursementsRequest) encode() string {
	if g.Pagination <= 0 {
		g.Pagination = 20
	}
	if g.Page <= 0 {
		g.Page = 1
	}

	query := &url.Values{}
	query.Add("pagination", strconv.Itoa(g.Pagination))
	query.Add("page", strconv.Itoa(g.Page))

	if g.Sort != "" {
		query.Add("sort", string(g.Sort))
	}
	if g.ID != 0 {
		query.Add("id", strconv.Itoa(g.ID))
	}
	if g.Amount > 0 {
		query.Add("amount", fmt.Sprintf("%g", g.Amount))
	}
	if g.Status != "" {
		query.Add("status", string(g.Status))
	}
	if g.Timestamp != "" {
		query.Add("timestamp", g.Timestamp)
	}
	if g.BankCode != "" {
		query.Add("bank_code", string(g.BankCode))
	}
	if g.AccountNumber != "" {
		query.Add("account_number", g.AccountNumber)
	}
	if g.RecipientName != "" {
		query.Add("recipient_name", g.RecipientName)
	}
	if g.Remark != "" {
		query.Add("remark", g.Remark)
	}
	if g.TimeServed != "" {
		query.Add("time_served", g.TimeServed)
	}
	if g.CreatedFrom != "" {
		query.Add("created_from", string(g.CreatedFrom))
	}
	if g.Direction != "" {
		query.Add("direction", string(g.Direction))
	}

	return query.Encode()
}

// CreateSpecialDisbursementRequest is request model to create special disbursement.
type CreateSpecialDisbursementRequest struct {
	IdempotencyKey       string       `validate:"required"`
	AccountNumber        string       `validate:"required"`
	BankCode             BankCode     `validate:"required"`
	Amount               float64      `validate:"required,gt=0"`
	Remark               string       ``
	RecipientCity        int          ``                    // from GetCities()
	SenderCountry        int          `validate:"required"` // from GetCountries()
	SenderPlaceOfBirth   int          ``                    // from GetCities()
	SenderDateOfBirth    string       ``                    // yyyy-mm-dd
	SenderIdentityType   IdentityType ``
	SenderName           string       `validate:"required"`
	SenderAddress        string       `validate:"required"`
	SenderIdentityNumber string       ``
	SenderJob            JobType      `validate:"required"`
	Direction            Direction    `validate:"required"`
}

func (c *CreateSpecialDisbursementRequest) encode() io.Reader {
	data := url.Values{}
	data.Set("account_number", c.AccountNumber)
	data.Set("bank_code", string(c.BankCode))
	data.Set("amount", fmt.Sprintf("%g", c.Amount))
	data.Set("sender_country", strconv.Itoa(c.SenderCountry))
	data.Set("sender_name", c.SenderName)
	data.Set("sender_address", c.SenderAddress)
	data.Set("sender_job", string(c.SenderJob))
	data.Set("direction", string(c.Direction))

	if c.Remark != "" {
		data.Set("remark", c.Remark)
	}
	if c.RecipientCity > 0 {
		data.Set("recipient_city", strconv.Itoa(c.RecipientCity))
	}
	if c.SenderPlaceOfBirth > 0 {
		data.Set("sender_place_of_birth", strconv.Itoa(c.SenderPlaceOfBirth))
	}
	if c.SenderDateOfBirth != "" {
		data.Set("sender_date_of_birth", c.SenderDateOfBirth)
	}
	if c.SenderIdentityType != "" {
		data.Set("sender_identity_type", string(c.SenderIdentityType))
	}
	if c.SenderIdentityNumber != "" {
		data.Set("sender_identity_number", c.SenderIdentityNumber)
	}

	return strings.NewReader(data.Encode())
}
