package flip

import (
	"net/http"
	"time"
)

// Client is flip client.
type Client struct {
	secretKey string
	baseURL   string
	requester Requester
	logger    Logger
}

// Option is config for flip client.
type Option struct {
	SecretKey string
	BaseURL   string
	Requester Requester
	Logger    Logger
}

// New to create new flip client with config.
func New(option Option) *Client {
	if option.Logger == nil {
		option.Logger = defaultLogger(LogError)
	}

	if option.Requester == nil {
		option.Requester = defaultRequester(&http.Client{
			Timeout: 10 * time.Second,
		}, option.Logger)
	}

	return &Client{
		secretKey: option.SecretKey,
		baseURL:   option.BaseURL,
		requester: option.Requester,
		logger:    option.Logger,
	}
}

// NewDefault to create new flip client with default config.
func NewDefault(secretKey string, env EnvironmentType) *Client {
	return New(Option{
		SecretKey: secretKey,
		BaseURL:   envURL[env],
		Requester: defaultRequester(&http.Client{
			Timeout: 10 * time.Second,
		}, defaultLogger(envLog[env])),
		Logger: defaultLogger(envLog[env]),
	})
}
