package feedly

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

const (
	defaultBaseURL = "https://cloud.feedly.com/v3/"
	userAgent      = "go-feedly"
)

type Client struct {
	clientMu  sync.Mutex
	client    *resty.Client
	BaseUrl   string
	Token     string
	UserAgent string
	common    service
	Alerts    *AlertsService
}

type service struct {
	client *Client
}

type ClientOption func(client *Client) error

func (c *Client) Client() *resty.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

func NewClient(options ...ClientOption) *Client {
	c := &Client{
		client: resty.New(),
	}
	for _, opt := range options {
		opt(c)
	}
	if c.Token != "" {
		c.client = c.client.SetAuthToken(c.Token)
	}
	if c.UserAgent == "" {
		c.UserAgent = userAgent
	}
	if c.BaseUrl == "" {
		c.BaseUrl = defaultBaseURL
	}
	c.client.SetBaseURL(c.BaseUrl)
	if c.Token != "" {
		c.client.SetAuthToken(c.Token)
	}
	c.common.client = c
	c.Alerts = (*AlertsService)(&c.common)
	return c
}

func WithToken(token string) ClientOption {
	return func(client *Client) error {
		client.Token = token
		return nil
	}
}

func WithBaseUrl(baseUrl string) ClientOption {
	return func(client *Client) error {
		client.BaseUrl = baseUrl
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(client *Client) error {
		client.UserAgent = userAgent
		return nil
	}
}
