package methods

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/bigelle/gotely/api/objects"
)

type SendGame struct {
	ChatId               int                           `json:"chat_id"`
	GameShortName        string                        `json:"game_short_name"`
	BusinessConnectionId *string                       `json:"business_connection_id,omitempty"`
	MessageThreadId      *int                          `json:"message_thread_id,omitempty"`
	DisableNotification  *bool                         `json:"disable_notification"`
	ProtectContent       *bool                         `json:"protect_content,omitempty"`
	AllowPaidBroadcast   *bool                         `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId      *string                       `json:"message_effect_id,omitempty"`
	ReplyParameters      *objects.ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup          *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (s SendGame) Validate() error {
	if s.ChatId == 0 {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.GameShortName) == "" {
		return objects.ErrInvalidParam("game_short_name parameter can't be empty")
	}
	if s.ReplyMarkup != nil {
		if err := s.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendGame) Endpoint() string {
	return "sendGame"
}

func (s SendGame) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SendGame) ContentType() string {
	return "application/json"
}

type SetGameHighScore struct {
	UserId             int     `json:"user_id"`
	Score              int     `json:"score"`
	Force              *bool   `json:"force,omitempty"`
	DisableEditMessage *bool   `json:"disable_edit_message,omitempty"`
	ChatId             *int    `json:"chat_id,omitempty"`
	MessageId          *int    `json:"message_id,omitempty"`
	InlineMessageId    *string `json:"inline_message_id,omitempty"`
}

func (s SetGameHighScore) Validate() error {
	if s.UserId == 0 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.Score < 0 {
		return objects.ErrInvalidParam("score parameter must be non-negative")
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		}
		if s.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id can't be empty if chat_id and message_id are not specified")
		}
	}
	return nil
}

func (s SetGameHighScore) Endpoint() string {
	return "setGameHighScore"
}

func (s SetGameHighScore) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetGameHighScore) ContentType() string {
	return "application/json"
}

type GetGameHighScores struct {
	UserId          int     `json:"user_id"`
	ChatId          *int    `json:"chat_id,omitempty"`
	MessageId       *int    `json:"message_id,omitempty"`
	InlineMessageId *string `json:"inline_message_id,omitempty"`
}

func (s GetGameHighScores) Validate() error {
	if s.UserId == 0 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		}
		if s.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id can't be empty if chat_id and message_id are not specified")
		}
	}
	return nil
}

func (s GetGameHighScores) Endpoint() string {
	return "getGameHighScores"
}

func (s GetGameHighScores) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s GetGameHighScores) ContentType() string {
	return "application/json"
}
