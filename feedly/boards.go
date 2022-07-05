package feedly

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type BoardsService service

type Board struct {
	Id             string  `json:"id"`
	Created        int64   `json:"created"`
	Customizable   bool    `json:"customizable"`
	Enterprise     bool    `json:"enterprise"`
	Label          string  `json:"label"`
	Description    *string `json:"description,omitempty"`
	IsPublic       *bool   `json:"isPublic,omitempty"`
	ShowNotes      *bool   `json:"showNotes,omitempty"`
	ShowHighlights *string `json:"showHighlights,omitempty"`
	HtmlUrl        *string `json:"htmlUrl,omitempty"`
	StreamId       *string `json:"streamId,omitempty"`
}

// List Get the list of boards
func (s *BoardsService) List(ctx context.Context) ([]Board, *resty.Response, error) {
	var result []Board
	resp, err := s.client.Client().R().SetResult(&result).SetContext(ctx).Get("/boards")
	return result, resp, err
}

type BoardsUpdateOneRequest struct {
	Id             string  `json:"id"`
	Label          *string `json:"label,omitempty"`
	Description    *string `json:"description,omitempty"`
	IsPublic       *bool   `json:"isPublic,omitempty"`
	ShowNotes      *bool   `json:"showNotes,omitempty"`
	ShowHighlights *string `json:"showHighlights,omitempty"`
}

// UpdateOne Update a board
func (s *BoardsService) UpdateOne(ctx context.Context, params *BoardsUpdateOneRequest) (*resty.Response, error) {
	return s.client.Client().R().SetContext(ctx).SetBody(params).Post("/boards")
}
