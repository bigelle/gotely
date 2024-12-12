package methods

import (
	"encoding/json"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type AnswerInlineQuery struct {
	InlineQueryId string
	Results       []types.InlineQueryResult
	CacheTime     *int
	IsPersonal    *bool
	NextOffset    *string
	Button        *types.InlineQueryResultsButton
}

func (a AnswerInlineQuery) Validate() error {
	if strings.TrimSpace(a.InlineQueryId) == "" {
		return types.ErrInvalidParam("inline_query_id parameter can't be empty")
	}
	for _, res := range a.Results {
		if err := res.Validate(); err != nil {
			return err
		}
	}
	if a.NextOffset != nil {
		if len([]byte(*a.NextOffset)) > 64 {
			return types.ErrInvalidParam("next_offset parameter can't exceed 64 bytes")
		}
	}
	if a.Button != nil {
		if err := a.Button.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (a AnswerInlineQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerInlineQuery) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "answerInlineQuery", a)
}

type AnswerWebAppQuery struct {
	WebAppQueryId string
	Result        types.InlineQueryResult
}

func (a AnswerWebAppQuery) Validate() error {
	if strings.TrimSpace(a.WebAppQueryId) == "" {
		return types.ErrInvalidParam("web_app_query_id parameter can't be empty")
	}
	if err := a.Result.Validate(); err != nil {
		return err
	}
	return nil
}

func (a AnswerWebAppQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerWebAppQuery) Execute() (*types.SentWebAppMessage, error) {
	return internal.MakePostRequest[types.SentWebAppMessage](telego.GetToken(), "answerWebAppQuery", a)
}

type SavePreparedInlineMessage struct {
	UserId            int
	Result            types.InlineQueryResult
	AllowUserChats    *bool
	AllowBotChats     *bool
	AllowGroupChats   *bool
	AllowChannelChats *bool
}

func (s SavePreparedInlineMessage) Validate() error {
	if s.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if err := s.Result.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SavePreparedInlineMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SavePreparedInlineMessage) Execute() (*types.PreparedInlineMessage, error) {
	return internal.MakePostRequest[types.PreparedInlineMessage](telego.GetToken(), "savePreparedInlineMessage", s)
}
