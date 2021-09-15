package flip

import (
	"context"
	"fmt"
	"net/http"
)

// CreateSpecialDisbursement to create special disbursement.
func (c *Client) CreateSpecialDisbursement(request CreateSpecialDisbursementRequest) (*Disbursement, error) {
	return c.CreateSpecialDisbursementWithContext(context.Background(), request)
}

// CreateSpecialDisbursementWithContext to create special disbursement with context.
func (c *Client) CreateSpecialDisbursementWithContext(ctx context.Context, request CreateSpecialDisbursementRequest) (*Disbursement, error) {
	if err := validate(&request); err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Add("idempotency-key", request.IdempotencyKey)

	var response Disbursement
	err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/special-disbursement", c.baseURL),
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
