package feedly

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/linbuxiao/go-feedly/util"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AlertsServiceTestSuite struct {
	suite.Suite
	client *Client
	ctx    context.Context
}

func (s *AlertsServiceTestSuite) SetupTest() {
	s.client = NewClient()
	httpmock.ActivateNonDefault(s.client.Client().GetClient())
	s.ctx = context.Background()
}

func (s *AlertsServiceTestSuite) TearDownTest() {
	httpmock.Deactivate()
}

func (s *AlertsServiceTestSuite) Context() context.Context {
	return s.ctx
}

func TestAlertsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AlertsServiceTestSuite))
}

func (s *AlertsServiceTestSuite) TestList() {
	tests := []struct {
		result           string
		status           int
		userAlertsLen    int
		enterpriseAlerts int
	}{
		{
			result: `{
  "userAlerts": [
    {
      "label": "Valve amplifier",
      "feedId": "feed/https://feedly.com/f/alert/11093cad-b6ec-43ba-9d5c-0c42c1c96c61",
      "velocity": 19.2,
      "layers": [
        {
          "type": "matches",
          "parts": [
            {
              "id": "nlp/f/entity/wd:786212",
              "label": "Valve amplifier"
            }
          ],
          "salience": "mention",
          "complexFilter": false
        }
      ],
      "created": 1624663223394
    }
  ],
  "enterpriseAlerts": [
    {
      "label": "Differential privacy",
      "feedId": "feed/https://feedly.com/f/alert/2e1996c8-148d-4e2c-beda-7eaadbcdb4d9",
      "velocity": 1.8,
      "layers": [
        {
          "type": "matches",
          "parts": [
            {
              "id": "nlp/f/entity/wd:5275358",
              "label": "Differential privacy"
            }
          ],
          "salience": "about",
          "complexFilter": false
        }
      ],
      "created": 1619750627429
    }
  ]
}`,
			status:           200,
			userAlertsLen:    1,
			enterpriseAlerts: 1,
		},
	}

	for _, tt := range tests {
		responder := util.NewJsonResponderFromString(tt.status, tt.result)
		httpmock.RegisterResponder(resty.MethodGet, DefaultBaseURL+"/alerts",
			responder)
		result, _, err := s.client.Alerts.List(s.Context())
		s.NoError(err)
		s.NotNil(result)
		s.Len(result.UserAlerts, tt.userAlertsLen)
		s.Len(result.EnterpriseAlerts, tt.enterpriseAlerts)
	}
}
