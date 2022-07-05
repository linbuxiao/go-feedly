package feedly

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type AlertsService service

type AlertsListResponse struct {
	UserAlerts       []Alert `json:"userAlerts"`
	EnterpriseAlerts []Alert `json:"enterpriseAlerts"`
}

type Alert struct {
	Label    string  `json:"label"`
	FeedId   string  `json:"feedId"`
	Velocity float64 `json:"velocity"`
	Layers   []Layer `json:"layers"`
	Created  int64   `json:"created"`
}

type Layer struct {
	Type          string      `json:"type"`
	Parts         []LayerPart `json:"parts"`
	Salience      string      `json:"salience"`
	ComplexFilter bool        `json:"complexFilter"`
}

type LayerPart struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

// List Get the list of web alerts (Pro+ and Enterprise only)
func (s *AlertsService) List(ctx context.Context) (*AlertsListResponse, *resty.Response, error) {
	var result AlertsListResponse
	resp, err := s.client.Client().R().SetResult(&result).SetContext(ctx).Get("/alerts")
	return &result, resp, err
}
