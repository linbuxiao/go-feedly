package feedly

import (
	"context"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

const (
	MOCK_TOKEN      = "mock-token"
	MOCK_BASE_URL   = "mock-url"
	MOCK_USER_AGENT = "mock-user-agent"
)

type FeedlyTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (f *FeedlyTestSuite) SetupTest() {
	f.ctx = context.Background()
}

func TestFeedlyTestSuite(t *testing.T) {
	suite.Run(t, new(FeedlyTestSuite))
}

func (f *FeedlyTestSuite) TestNewClient() {
	f.Run("with token", func() {
		c := NewClient(WithToken(MOCK_TOKEN))
		f.Equal(MOCK_TOKEN, c.client.Token)
	})

	f.Run("with no token", func() {
		c := NewClient()
		f.Equal("", c.client.Token)
	})

	f.Run("with base url", func() {
		c := NewClient(WithBaseUrl(MOCK_BASE_URL))
		f.Equal(MOCK_BASE_URL, c.client.BaseURL)
	})

	f.Run("with default base url", func() {
		c := NewClient()
		f.Equal(DefaultBaseURL, c.client.BaseURL)
	})

	f.Run("with user agent", func() {
		c := NewClient(WithUserAgent(MOCK_USER_AGENT))
		f.Equal(MOCK_USER_AGENT, c.client.Header.Get("user-agent"))
	})

	f.Run("with default user agent", func() {
		c := NewClient()
		f.Equal(DefaultUserAgent, c.client.Header.Get("user-agent"))
	})

	f.Run("with client", func() {
		const timeout = time.Duration(10000)
		mockClient := &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       timeout,
		}

		c := NewClient(WithClient(mockClient))
		f.Equal(timeout, c.client.GetClient().Timeout)
	})
}
