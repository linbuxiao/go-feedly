package feedly

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type FeedsService service

type Feed struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Topics      *[]string `json:"topics,omitempty"`
	Language    *string   `json:"language,omitempty"`
	Website     *string   `json:"website,omitempty"`
	Velocity    *float64  `json:"velocity,omitempty"`
	Featured    bool      `json:"featured"`
	Sponsored   bool      `json:"sponsored"`
	Curated     bool      `json:"curated"`
	Subscribers int       `json:"subscribers"`
	State       *string   `json:"state,omitempty"`
	Description *string   `json:"description,omitempty"`
}

type FeedsGetResponse Feed

// Get the metadata about a specific feed
func (s *FeedsService) Get(ctx context.Context, feedId string) (*FeedsGetResponse, *resty.Response, error) {
	var result FeedsGetResponse
	url := fmt.Sprintf("/feeds/%s", feedId)
	resp, err := s.client.Do(ctx, resty.MethodGet, url, nil, &result)
	return &result, resp, err
}

type FeedsGetListRequest []string
type FeedsGetListResponse []Feed

// GetList Get the metadata for a list of feeds
func (s *FeedsService) GetList(ctx context.Context, params *FeedsGetListRequest) (*FeedsGetListResponse, *resty.Response, error) {
	var result FeedsGetListResponse
	resp, err := s.client.Do(ctx, resty.MethodPost, "/feeds/.mget", params, &result)
	return &result, resp, err
}
