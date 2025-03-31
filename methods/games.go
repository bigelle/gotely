package methods

import (
	"fmt"
	"io"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/objects"
)

// Use this method to send a game.
// On success, the sent [objects.Message] is returned.
type SendGame struct {
	// REQUIRED:
	// Unique identifier for the target chat
	ChatId int `json:"chat_id"`
	// REQUIRED:
	// Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	GameShortName string `json:"game_short_name"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown.
	// If not empty, the first button must launch the game.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (s SendGame) Validate() error {
	var err gotely.ErrFailedValidation
	if s.ChatId == 0 {
		err = append(err, fmt.Errorf("chat_id parameter can't be empty"))
	}
	if s.GameShortName == "" {
		err = append(err, fmt.Errorf("game_short_name parameter can't be empty"))
	}
	if s.ReplyMarkup != nil {
		if er := s.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if s.ReplyParameters != nil {
		if er := s.ReplyParameters.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SendGame) Endpoint() string {
	return "sendGame"
}

func (s SendGame) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendGame) ContentType() string {
	return "application/json"
}

// Use this method to set the score of the specified user in a game message.
// On success, if the message is not an inline message, the [objects.Message] is returned, otherwise True is returned.
// Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	// REQUIRED:
	// User identifier
	UserId int `json:"user_id"`
	// REQUIRED:
	// New score, must be non-negative
	Score int `json:"score"`

	// Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	Force *bool `json:"force,omitempty"`
	// Pass True if the game message should not be automatically edited to include the current scoreboard
	DisableEditMessage *bool `json:"disable_edit_message,omitempty"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatId *int `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the sent message
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
}

func (s SetGameScore) Validate() error {
	var err gotely.ErrFailedValidation
	if s.UserId == 0 {
		err = append(err, fmt.Errorf("user_id parameter can't be empty"))
	}
	if s.Score < 0 {
		err = append(err, fmt.Errorf("score parameter must be non-negative"))
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if s.MessageId == nil {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id can't be empty if chat_id and message_id are not specified"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetGameScore) Endpoint() string {
	return "setGameScore"
}

func (s SetGameScore) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetGameScore) ContentType() string {
	return "application/json"
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game.
// Returns an Array of [objects.GameHighScore] objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side.
// Will also return the top three users if the user and their neighbors are not among them.
// Please note that this behavior is subject to change.
type GetGameHighScores struct {
	// REQUIRED:
	// Target user id
	UserId int `json:"user_id"`

	// Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatId *int `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the sent message
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
}

func (s GetGameHighScores) Validate() error {
	var err gotely.ErrFailedValidation
	if s.UserId == 0 {
		err = append(err, fmt.Errorf("user_id parameter can't be empty"))
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if s.MessageId == nil {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id can't be empty if chat_id and message_id are not specified"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s GetGameHighScores) Endpoint() string {
	return "getGameHighScores"
}

func (s GetGameHighScores) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetGameHighScores) ContentType() string {
	return "application/json"
}
