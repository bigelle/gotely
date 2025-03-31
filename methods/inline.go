package methods

import (
	"fmt"
	"io"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/objects"
)

// Use this method to send answers to an inline query.
// On success, True is returned.
//
// No more than 50 results per query are allowed.
type AnswerInlineQuery struct {
	// REQUIRED:
	// Unique identifier for the answered query
	InlineQueryId string `json:"inline_query_id"`
	// REQUIRED:
	// A JSON-serialized array of results for the inline query
	Results []objects.InlineQueryResult `json:"results"`

	// The maximum amount of time in seconds that the result of the inline query may be cached on the server.
	// Defaults to 300.
	CacheTime *int `json:"cache_time,omitempty"`
	// Pass True if results may be cached on the server side only for the user that sent the query.
	// By default, results may be returned to any user who sends the same query.
	IsPersonal *bool `json:"is_personal,omitempty"`
	// Pass the offset that a client should send in the next query with the same text to receive more results.
	// Pass an empty string if there are no more results or if you don't support pagination.
	// Offset length can't exceed 64 bytes.
	NextOffset *string `json:"next_offset,omitempty"`
	// A JSON-serialized object describing a button to be shown above inline query results
	Button *objects.InlineQueryResultsButton `json:"button,omitempty"`
}

func (a AnswerInlineQuery) Validate() error {
	var err gotely.ErrFailedValidation
	if a.InlineQueryId == "" {
		err = append(err, fmt.Errorf("inline_query_id parameter can't be empty"))
	}
	for _, res := range a.Results {
		if er := res.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if a.NextOffset != nil {
		if len([]byte(*a.NextOffset)) > 64 {
			err = append(err, fmt.Errorf("next_offset parameter can't exceed 64 bytes"))
		}
	}
	if a.Button != nil {
		if er := a.Button.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s AnswerInlineQuery) Endpoint() string {
	return "answerInlineQuery"
}

func (s AnswerInlineQuery) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s AnswerInlineQuery) ContentType() string {
	return "application/json"
}

// Use this method to set the result of an interaction with a Web App and send
// a corresponding message on behalf of the user to the chat from which the query originated.
// On success, a [objects.SentWebAppMessage] object is returned.
type AnswerWebAppQuery struct {
	// REQUIRED:
	// Unique identifier for the query to be answered
	WebAppQueryId string `json:"web_app_query_id"`
	// REQUIRED:
	// A JSON-serialized object describing the message to be sent
	Result objects.InlineQueryResult `json:"result"`
}

func (a AnswerWebAppQuery) Validate() error {
	var err gotely.ErrFailedValidation
	if a.WebAppQueryId == "" {
		err = append(err, fmt.Errorf("web_app_query_id parameter can't be empty"))
	}
	if er := a.Result.Validate(); er != nil {
		err = append(err, er)
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s AnswerWebAppQuery) Endpoint() string {
	return "answerWebAppQuery"
}

func (s AnswerWebAppQuery) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s AnswerWebAppQuery) ContentType() string {
	return "application/json"
}

// Stores a message that can be sent by a user of a Mini App.
// Returns a [objects.PreparedInlineMessage] object.
type SavePreparedInlineMessage struct {
	// REQUIRED:
	// Unique identifier of the target user that can use the prepared message
	UserId int `json:"user_id"`
	// REQUIRED:
	// A JSON-serialized object describing the message to be sent
	Result objects.InlineQueryResult `json:"result"`

	// Pass True if the message can be sent to private chats with users
	AllowUserChats *bool `json:"allow_user_chats,omitempty"`
	// Pass True if the message can be sent to private chats with bots
	AllowBotChats *bool `json:"allow_bot_chats,omitempty"`
	// Pass True if the message can be sent to group and supergroup chats
	AllowGroupChats *bool `json:"allow_group_chats,omitempty"`
	// Pass True if the message can be sent to channel chats
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty"`
}

func (s SavePreparedInlineMessage) Validate() error {
	var err gotely.ErrFailedValidation
	if s.UserId < 1 {
		err = append(err, fmt.Errorf("user_id parameter can't be empty"))
	}
	if er := s.Result.Validate(); er != nil {
		err = append(err, er)
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SavePreparedInlineMessage) Endpoint() string {
	return "savePreparedInlineMessage"
}

func (s SavePreparedInlineMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SavePreparedInlineMessage) ContentType() string {
	return "application/json"
}
