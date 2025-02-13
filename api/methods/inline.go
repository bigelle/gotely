package methods

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/bigelle/gotely/api/objects"
)

type AnswerInlineQuery struct {
	InlineQueryId string                            `json:"inline_query_id"`
	Results       []objects.InlineQueryResult       `json:"results"`
	CacheTime     *int                              `json:"cache_time,omitempty"`
	IsPersonal    *bool                             `json:"is_personal,omitempty"`
	NextOffset    *string                           `json:"next_offset,omitempty"`
	Button        *objects.InlineQueryResultsButton `json:"button,omitempty"`
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

func (s AnswerInlineQuery) Endpoint() string {
	return "answerInlineQuery"
}

func (s AnswerInlineQuery) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s AnswerInlineQuery) ContentType() string {
	return "application/json"
}

type AnswerWebAppQuery struct {
	WebAppQueryId string                    `json:"web_app_query_id"`
	Result        objects.InlineQueryResult `json:"result"`
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

func (s AnswerWebAppQuery) Endpoint() string {
	return "answerWebAppQuery"
}

func (s AnswerWebAppQuery) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s AnswerWebAppQuery) ContentType() string {
	return "application/json"
}

type SavePreparedInlineMessage struct {
	UserId            int                       `json:"user_id"`
	Result            objects.InlineQueryResult `json:"result"`
	AllowUserChats    *bool                     `json:"allow_user_chats,omitempty"`
	AllowBotChats     *bool                     `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   *bool                     `json:"allow_group_chats,omitempty"`
	AllowChannelChats *bool                     `json:"allow_channel_chats,omitempty"`
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

func (s SavePreparedInlineMessage) Endpoint() string {
	return "savePreparedInlineMessage"
}

func (s SavePreparedInlineMessage) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SavePreparedInlineMessage) ContentType() string {
	return "application/json"
}
