package methods

import (
	"encoding/json"
	"strings"

	"github.com/bigelle/gotely/objects"
)

type AnswerInlineQuery struct {
	InlineQueryId string
	Results       []objects.InlineQueryResult
	CacheTime     *int
	IsPersonal    *bool
	NextOffset    *string
	Button        *objects.InlineQueryResultsButton
}

func (a AnswerInlineQuery) Validate() error {
	if strings.TrimSpace(a.InlineQueryId) == "" {
		return objects.ErrInvalidParam("inline_query_id parameter can't be empty")
	}
	for _, res := range a.Results {
		if err := res.Validate(); err != nil {
			return err
		}
	}
	if a.NextOffset != nil {
		if len([]byte(*a.NextOffset)) > 64 {
			return objects.ErrInvalidParam("next_offset parameter can't exceed 64 bytes")
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
	return MakePostRequest[bool]("answerInlineQuery", a)
}

type AnswerWebAppQuery struct {
	WebAppQueryId string
	Result        objects.InlineQueryResult
}

func (a AnswerWebAppQuery) Validate() error {
	if strings.TrimSpace(a.WebAppQueryId) == "" {
		return objects.ErrInvalidParam("web_app_query_id parameter can't be empty")
	}
	if err := a.Result.Validate(); err != nil {
		return err
	}
	return nil
}

func (a AnswerWebAppQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerWebAppQuery) Execute() (*objects.SentWebAppMessage, error) {
	return MakePostRequest[objects.SentWebAppMessage]("answerWebAppQuery", a)
}

type SavePreparedInlineMessage struct {
	UserId            int
	Result            objects.InlineQueryResult
	AllowUserChats    *bool
	AllowBotChats     *bool
	AllowGroupChats   *bool
	AllowChannelChats *bool
}

func (s SavePreparedInlineMessage) Validate() error {
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if err := s.Result.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SavePreparedInlineMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SavePreparedInlineMessage) Execute() (*objects.PreparedInlineMessage, error) {
	return MakePostRequest[objects.PreparedInlineMessage]("savePreparedInlineMessage", s)
}
