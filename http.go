package flip

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Requester is http request interface.
type Requester interface {
	Call(ctx context.Context, method, url, secretKey string, header http.Header, request []byte, response interface{}) (statusCode int, err error)
}

type requester struct {
	client *http.Client
	logger Logger
}

func defaultRequester(client *http.Client, logger Logger) *requester {
	return &requester{
		client: client,
		logger: logger,
	}
}

// Call to prepare request and execute.
func (r *requester) Call(ctx context.Context, method, url, secretKey string, header http.Header, request []byte, response interface{}) (int, error) {
	now := time.Now()

	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(string(request)))
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	if header != nil {
		req.Header = header
	}

	req.SetBasicAuth(secretKey, "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	r.logger.Debug("%s %s", method, url)
	r.logRequestHeader(req.Header)
	r.logRequestBody(request)
	defer func() { r.logger.Info("%s %s [%s]", method, url, time.Since(now)) }()

	return r.doRequest(req, response)
}

type errGeneralResponse struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type errResponse struct {
	Code   string              `json:"code"`
	Errors []errResponseDetail `json:"errors"`
}

type errResponseDetail struct {
	Attribute string `json:"attribute"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func (r *requester) doRequest(req *http.Request, response interface{}) (int, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	// Flip error response is a bit weird and not consistent.
	r.logResponseBody(resp.StatusCode, respBody)

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusIMUsed {
		// Try parse to general error response first.
		var errGenResp errGeneralResponse
		if json.Unmarshal(respBody, &errGenResp) == nil {
			if errGenResp.Status != 0 {
				return resp.StatusCode, errors.New(errGenResp.Message)
			}
		}

		// Parse to error response with inner code.
		var errResp errResponse
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			r.logger.Error(err.Error())
			return http.StatusInternalServerError, ErrInternal
		}
		return resp.StatusCode, errors.New(errResp.Errors[0].Message)
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	return resp.StatusCode, nil
}

func (r *requester) logRequestHeader(header http.Header) {
	if header == nil || len(header) == 0 {
		return
	}

	for k, h := range header {
		for _, v := range h {
			r.logger.Debug("header: %s: %s", k, v)
		}
	}
}

func (r *requester) logRequestBody(request []byte) {
	if request == nil {
		return
	}
	r.logger.Debug("request: %s", string(request))
}

func (r *requester) logResponseBody(code int, response []byte) {
	if response == nil {
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, response, "", "  "); err != nil {
		r.logger.Error(err.Error())
		r.logger.Debug("response: %d %s", code, string(response))
		return
	}

	r.logger.Debug("response: %d %s", code, out.String())
}
