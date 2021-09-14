package flip

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Requester is http request interface.
type Requester interface {
	Call(ctx context.Context, method, url, secretKey string, header http.Header, request io.Reader, response interface{}) error
}

type requester struct {
	client *http.Client
	logger Logger
}

// DefaultRequester to create new http request client.
func DefaultRequester(client *http.Client, logger Logger) *requester {
	return &requester{
		client: client,
		logger: logger,
	}
}

// Call to prepare request and execute.
func (r *requester) Call(ctx context.Context, method, url, secretKey string, header http.Header, request io.Reader, response interface{}) (err error) {
	now := time.Now()

	req, err := http.NewRequestWithContext(ctx, method, url, request)
	if err != nil {
		r.logger.Error(err.Error())
		return ErrInternal
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

type errResponse struct {
	Code   string              `json:"code"`
	Errors []errResponseDetail `json:"errors"`
}

type errResponseDetail struct {
	Attribute string `json:"attribute"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func (r *requester) doRequest(req *http.Request, response interface{}) error {
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Error(err.Error())
		return ErrInternal
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err.Error())
		return ErrInternal
	}

	r.logResponseBody(resp.StatusCode, respBody)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errResp errResponse
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			r.logger.Error(err.Error())
			return ErrInternal
		}
		return errors.New(errResp.Code)
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		r.logger.Error(err.Error())
		return ErrInternal
	}

	return nil
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

func (r *requester) logRequestBody(request io.Reader) {
	if request == nil {
		return
	}

	var sb strings.Builder
	if _, err := io.Copy(&sb, request); err != nil {
		return
	}

	r.logger.Debug("request: %s", sb.String())
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

	r.logger.Debug("response: %d %s", code, string(out.Bytes()))
}
