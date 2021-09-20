package flip

import (
	"context"
	"fmt"
	"net/http"
)

// CreateSpecialDisbursement to create special disbursement.
func (c *Client) CreateSpecialDisbursement(request CreateSpecialDisbursementRequest) (*Disbursement, int, error) {
	return c.CreateSpecialDisbursementWithContext(context.Background(), request)
}

// CreateSpecialDisbursementWithContext to create special disbursement with context.
func (c *Client) CreateSpecialDisbursementWithContext(ctx context.Context, request CreateSpecialDisbursementRequest) (*Disbursement, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	header := make(http.Header)
	header.Add("idempotency-key", request.IdempotencyKey)

	var response Disbursement
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/special-disbursement", c.baseURL),
		c.secretKey,
		header,
		[]byte(request.encode()),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response, code, nil
}
