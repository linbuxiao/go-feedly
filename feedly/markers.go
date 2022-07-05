package feedly

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

var ErrMarkersLastReadEntryIdAndAsOfParams = errors.New("lastReadEntryId and asOf one of is required")

type MakersService service

type MakersGetUnreadCountsListResponse struct {
	Unreadcounts []struct {
		Id      string `json:"id"`
		Count   int    `json:"count"`
		Updated int64  `json:"updated"`
	} `json:"unreadcounts"`
	Updated int64 `json:"updated"`
}

// GetUnreadCountsList Get the list of unread counts
func (s *MakersService) GetUnreadCountsList(ctx context.Context) (*MakersGetUnreadCountsListResponse, *resty.Response, error) {
	var result MakersGetUnreadCountsListResponse
	resp, err := s.client.Do(ctx, resty.MethodGet, "/markers/counts", nil, &result)
	return &result, resp, err
}

type MakersMarkOneOrMultipleArticlesAsReadRequest struct {
	Action   string   `json:"action"`
	Type     string   `json:"type"`
	EntryIds []string `json:"entryIds"`
}

// MarkOneOrMultipleArticlesAsRead Mark one or multiple articles as read
func (s *MakersService) MarkOneOrMultipleArticlesAsRead(ctx context.Context, params *MakersMarkOneOrMultipleArticlesAsReadRequest) (
	*resty.Response, error,
) {
	return s.client.Do(ctx, resty.MethodPost, "/markers", params, nil)
}

type LastReadEntryIdAndAsOf struct {
	LastReadEntryId *string `json:"lastReadEntryId,omitempty"`
	AsOf            *string `json:"asOf,omitempty"`
}

type MakersMarkOneFeedAsReadRequest struct {
	Action  string   `json:"action"`
	Type    string   `json:"type"`
	FeedIds []string `json:"feedIds"`
	LastReadEntryIdAndAsOf
}

// markAsRead Use for MarkFeedAsRead, MarkCategoryAsRead, MarkTagAsRead
func (s *MakersService) markAsRead(ctx context.Context, params interface{}) (*resty.Response, error) {
	b, _ := json.Marshal(params)
	lastReadEntryId := gjson.Get(string(b), "lastReadEntryId")
	asOf := gjson.Get(string(b), "asOf")
	if lastReadEntryId.Exists() && asOf.Exists() || !lastReadEntryId.Exists() && !asOf.Exists() {
		return nil, ErrMarkersLastReadEntryIdAndAsOfParams
	}
	return s.client.Do(ctx, resty.MethodPost, "/markers", params, nil)
}

// MarkFeedAsRead Mark a feed as read
func (s *MakersService) MarkFeedAsRead(ctx context.Context, params *MakersMarkOneFeedAsReadRequest) (*resty.Response, error) {
	return s.markAsRead(ctx, params)
}

type MakersMarkCategoryAsReadRequest struct {
	Action      string   `json:"action"`
	Type        string   `json:"type"`
	CategoryIds []string `json:"categoryIds"`
	LastReadEntryIdAndAsOf
}

// MarkCategoryAsRead Mark a category as read
func (s *MakersService) MarkCategoryAsRead(ctx context.Context, params *MakersMarkCategoryAsReadRequest) (*resty.Response, error) {
	return s.markAsRead(ctx, params)
}

type MakersMarkTagAsReadRequest struct {
	Action string   `json:"action"`
	Type   string   `json:"type"`
	TagIds []string `json:"tagIds"`
	LastReadEntryIdAndAsOf
}

// MarkTagAsRead Mark a tag as read
func (s *MakersService) MarkTagAsRead(ctx context.Context, params *MakersMarkTagAsReadRequest) (*resty.Response, error) {
	return s.markAsRead(ctx, params)
}
