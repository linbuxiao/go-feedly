package feedly

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type FeedCollectionsService service

type FeedCollectionsListResponse []FeedCollection

type FeedCollection struct {
	Id          string  `json:"id"`
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	Cover       *string `json:"cover,omitempty"`
	Created     *int64  `json:"created,omitempty"`
	Feeds       []struct {
		Id                          string  `json:"id"`
		Title                       string  `json:"title,omitempty"`
		Added                       int64   `json:"added,omitempty"`
		Updated                     int64   `json:"updated,omitempty"`
		Website                     string  `json:"website,omitempty"`
		VisualUrl                   string  `json:"visualUrl,omitempty"`
		Velocity                    float64 `json:"velocity,omitempty"`
		NumReadEntriesPastMonth     int     `json:"numReadEntriesPastMonth,omitempty"`
		NumLongReadEntriesPastMonth int     `json:"numLongReadEntriesPastMonth,omitempty"`
		TotalReadingTimePastMonth   int     `json:"totalReadingTimePastMonth,omitempty"`
		NumTaggedEntriesPastMonth   int     `json:"numTaggedEntriesPastMonth,omitempty"`
		FeedId                      string  `json:"feedId,omitempty"`
	} `json:"feeds"`
}

// List Get the list of personal collections
func (s *FeedCollectionsService) List(ctx context.Context) (*FeedCollectionsListResponse, *resty.Response, error) {
	var result FeedCollectionsListResponse
	resp, err := s.client.Do(ctx, resty.MethodGet, "/collections", nil, &result)
	return &result, resp, err
}

// Get details about a personal collection
func (s *FeedCollectionsService) Get(ctx context.Context, id string) (*FeedCollection, *resty.Response, error) {
	var result FeedCollection
	url := fmt.Sprintf("/collections/%s", id)
	resp, err := s.client.Do(ctx, resty.MethodGet, url, nil, &result)
	return &result, resp, err
}

type FeedRo struct {
	Id    string `json:"id"`
	Title string `json:"title,omitempty"`
}

type FeedCollectionsCreateOrUpdateOneRequest struct {
	Id          string    `json:"id"`
	Label       string    `json:"label,omitempty"`
	Description string    `json:"description,omitempty"`
	Feeds       *[]FeedRo `json:"feeds,omitempty"`
}

// CreateOrUpdateOne Create or update a personal collection
func (s *FeedCollectionsService) CreateOrUpdateOne(ctx context.Context,
	params *FeedCollectionsCreateOrUpdateOneRequest,
) (*resty.Response, error) {
	return s.client.Do(ctx, resty.MethodPost, "/collections", params, nil)
}

type FeedCollectionsAddFeedRequest FeedRo

// AddFeed Add a feed to a personal collection
func (s *FeedCollectionsService) AddFeed(ctx context.Context, params *FeedCollectionsAddFeedRequest) (*resty.Response, error) {
	url := fmt.Sprintf("/collections/%s/feeds", params.Id)
	return s.client.Do(ctx, resty.MethodPut, url, params, nil)
}

type FeedCollectionsAddMultipleFeedsRequest []FeedRo

// AddMultipleFeeds Add multiple feeds to a personal collection
func (s *FeedCollectionsService) AddMultipleFeeds(ctx context.Context,
	collectionId string, params *FeedCollectionsAddMultipleFeedsRequest,
) (*resty.Response, error) {
	url := fmt.Sprintf("/collections/%s/feeds/.mput", collectionId)
	return s.client.Do(ctx, resty.MethodPost, url, params, nil)
}

// RemoveFeed Remove multiple feeds from a personal collection
func (s *FeedCollectionsService) RemoveFeed(ctx context.Context,
	collectionId string, feedId string,
) (*resty.Response, error) {
	url := fmt.Sprintf("/collections/%s/feeds/%s", collectionId, feedId)
	return s.client.Do(ctx, resty.MethodDelete, url, nil, nil)
}
