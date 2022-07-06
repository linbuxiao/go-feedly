package feedly

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/linbuxiao/go-feedly/util"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BoardsServiceTestSuite struct {
	suite.Suite
	client *Client
	ctx    context.Context
}

func (s *BoardsServiceTestSuite) SetupTest() {
	s.client = NewClient()
	httpmock.ActivateNonDefault(s.client.Client().GetClient())
	s.ctx = context.Background()
}

func (s *BoardsServiceTestSuite) TearDownTest() {
	httpmock.Deactivate()
}

func TestBoardsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BoardsServiceTestSuite))
}

func (s *BoardsServiceTestSuite) TestList() {
	str := `[
  {
    "id": "user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/tag/global.saved"
  },
  {
    "id": "user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/tag/tech",
    "label": "tech",
    "description": "Best of tech articles"
  },
  {
    "id": "user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/tag/inspiration",
    "label": "inspiration"
  }
]`
	responder := util.NewJsonResponderFromString(200, str)
	httpmock.RegisterResponder(resty.MethodGet, DefaultBaseURL+"/boards", responder)
	result, _, err := s.client.Boards.List(s.ctx)
	s.NoError(err)
	s.NotNil(result)
	s.Len(result, 3)
}

func (s *BoardsServiceTestSuite) TestUpdateOne() {
	var result BoardsUpdateOneRequest
	input := `{
  "id": "user/af190c49-0ac8-4f08-9f83-805f1a3bc142/tag/marketing",
  "label": "new label",
  "description": "new description",
  "isPublic": true,
  "showNotes": true,
  "showHighlights": "false"
}`
	err := json.Unmarshal([]byte(input), &result)
	s.NoError(err)
	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder(resty.MethodPost, DefaultBaseURL+"/boards", responder)
	_, err = s.client.Boards.UpdateOne(s.ctx, &result)
	s.NoError(err)
}
