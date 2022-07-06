package feedly

import (
	"context"
	"github.com/go-resty/resty/v2"
	"net/http"
	"sync"
)

const (
	DefaultBaseURL   = "https://cloud.feedly.com/v3"
	DefaultUserAgent = "go-feedly"
)

type Client struct {
	clientMu  sync.Mutex
	client    *resty.Client
	rawClient *http.Client
	BaseUrl   string
	Token     string
	UserAgent string
	common    service
	Alerts    *AlertsService
	Boards    *BoardsService
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

func (c *Client) Do(ctx context.Context, method string,
	url string, body interface{}, v interface{},
) (*resty.Response, error) {
	r := c.Client().R().SetContext(ctx)
	if body != nil {
		r = r.SetBody(body)
	}
	if v != nil {
		r = r.SetResult(&v)
	}
	return r.Execute(method, url)
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
	if c.rawClient != nil {
		c.client = resty.NewWithClient(c.rawClient)
	}
	if c.UserAgent == "" {
		c.UserAgent = DefaultUserAgent
	}
	if c.BaseUrl == "" {
		c.BaseUrl = DefaultBaseURL
	}
	c.client.SetBaseURL(c.BaseUrl)
	c.client.Header.Set("user-agent", c.UserAgent)
	if c.Token != "" {
		c.client.SetAuthToken(c.Token)
	}
	c.common.client = c
	c.Alerts = (*AlertsService)(&c.common)
	c.Boards = (*BoardsService)(&c.common)
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

func WithClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.rawClient = c
		return nil
	}
}
