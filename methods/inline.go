package methods

import (
	"encoding/json"
	"net/http"
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
	client        *http.Client
	baseUrl       string
}

func (s *AnswerInlineQuery) WithClient(c *http.Client) *AnswerInlineQuery {
	s.client = c
	return s
}

func (s AnswerInlineQuery) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *AnswerInlineQuery) WithApiBaseUrl(u string) *AnswerInlineQuery {
	s.baseUrl = u
	return s
}

func (s AnswerInlineQuery) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (a AnswerInlineQuery) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "answerInlineQuery", a)
}

type AnswerWebAppQuery struct {
	WebAppQueryId string
	Result        objects.InlineQueryResult
	client        *http.Client
	baseUrl       string
}

func (s *AnswerWebAppQuery) WithClient(c *http.Client) *AnswerWebAppQuery {
	s.client = c
	return s
}

func (s AnswerWebAppQuery) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *AnswerWebAppQuery) WithApiBaseUrl(u string) *AnswerWebAppQuery {
	s.baseUrl = u
	return s
}

func (s AnswerWebAppQuery) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (a AnswerWebAppQuery) Execute(token string) (*objects.SentWebAppMessage, error) {
	return SendTelegramPostRequest[objects.SentWebAppMessage](token, "answerWebAppQuery", a)
}

type SavePreparedInlineMessage struct {
	UserId            int
	Result            objects.InlineQueryResult
	AllowUserChats    *bool
	AllowBotChats     *bool
	AllowGroupChats   *bool
	AllowChannelChats *bool
	client            *http.Client
	baseUrl           string
}

func (s *SavePreparedInlineMessage) WithClient(c *http.Client) *SavePreparedInlineMessage {
	s.client = c
	return s
}

func (s SavePreparedInlineMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SavePreparedInlineMessage) WithApiBaseUrl(u string) *SavePreparedInlineMessage {
	s.baseUrl = u
	return s
}

func (s SavePreparedInlineMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SavePreparedInlineMessage) Execute(token string) (*objects.PreparedInlineMessage, error) {
	return SendTelegramPostRequest[objects.PreparedInlineMessage](token, "savePreparedInlineMessage", s)
}
