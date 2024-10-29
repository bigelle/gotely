package types

import (
	"fmt"

	"github.com/bigelle/utils.go/ensure"

	"github.com/bigelle/tele.go/internal/assertions"
)

type InlineKeyboardMarkup struct {
	Keyboard []InlineKeyboardRow `json:"keyboard"`
}

func (f InlineKeyboardMarkup) replyKeyboardContract() {}

func (m InlineKeyboardMarkup) Validate() error {
	for _, buttons := range m.Keyboard {
		if err := buttons.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineKeyboardRow []InlineKeyboardButton

func (r InlineKeyboardRow) Validate() error {
	for _, button := range r {
		if err := button.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          *string                      `json:"url,omitempty"`
	CallbackData                 *string                      `json:"callback_data,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	Pay                          *bool                        `json:"pay,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
}

func (b InlineKeyboardButton) Validate() error {
	if !ensure.NotEmpty(b.Text) {
		return fmt.Errorf("text parameter can't be empty")
	}
	if !ensure.NotNil(b.LoginUrl) {
		if err := (*b.LoginUrl).Validate(); err != nil {
			return err
		}
	}
	if !ensure.NotNil(b.WebApp) {
		if err := (*b.WebApp).Validate(); err != nil {
			return err
		}
	}
	//	if !ensure.NotNil(b.CallbackGame) {
	//		if err := (*b.CallbackGame).Validate(); err != nil {
	//			return err
	//		}
	//	}

	return nil
}

type SwitchInlineQueryChosenChat struct {
	RequestId         *string `json:"request_id,omitempty"`
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"`
}

type ReplyKeyboardInterface interface {
	replyKeyboardContract()
}

type ReplyKeyboard struct {
	ReplyKeyboardInterface
}

type ReplyKeyboardMarkup struct {
	Keyboard              []KeyboardRow `json:"keyboard"`
	ResizeKeyboard        *bool         `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       *bool         `json:"one_time_keyboard,omitempty"`
	Selective             *bool         `json:"selective,omitempty"`
	InputFieldPlaceholder *string       `json:"input_field_placeholder,omitempty"`
	IsPersistent          *bool         `json:"is_persistent,omitempty"`
}

func (f ReplyKeyboardMarkup) replyKeyboardContract() {}

func (r ReplyKeyboardMarkup) Validate() error {
	if len(*r.InputFieldPlaceholder) < 1 || len(*r.InputFieldPlaceholder) > 64 {
		return fmt.Errorf("InputFieldPlaceholder parameter must be between 1 and 64 characters")
	}
	for _, row := range r.Keyboard {
		if err := row.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	Selective             *bool   `json:"selective,omitempty"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
}

func (f ForceReply) replyKeyboardContract() {}

func (f ForceReply) Validate() error {
	if len(*f.InputFieldPlaceholder) < 1 || len(*f.InputFieldPlaceholder) > 64 {
		return fmt.Errorf("InputFieldPlaceholder parameter must be between 1 and 64 characters")
	}
	return nil
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

func (f ReplyKeyboardRemove) replyKeyboardContract() {}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
	RequestContact  *bool                       `json:"request_contact,omitempty"`
	RequestLocation *bool                       `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	RequestUser     *KeyboardButtonRequestUser  `json:"request_user,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
}

func (k KeyboardButton) Validate() error {
	if err := assertions.ParamNotEmpty(k.Text, "Text"); err != nil {
		return err
	}

	requestsProvided := 0
	if *k.RequestContact {
		requestsProvided++
	}
	if *k.RequestLocation {
		requestsProvided++
	}
	if k.WebApp != nil {
		if err := k.WebApp.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestPoll != nil {
		if err := k.RequestPoll.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestUser != nil {
		if err := k.RequestUser.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestChat != nil {
		if err := k.RequestChat.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestUsers != nil {
		if err := k.RequestUsers.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if requestsProvided > 1 {
		return fmt.Errorf(
			"RequestContact, RequestLocation, WebApp, RequestPoll, RequestUser, RequestChat and RequestUsers are mutually exclusive",
		)
	}

	return nil
}

type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

// currently does not do anything
func (k KeyboardButtonPollType) Validate() error {
	return nil
}

type KeyboardButtonRequestChat struct {
	RequestId                   string                       `json:"request_id"`
	ChatIsChannel               bool                         `json:"chat_is_channel"`
	ChatIsForum                 *bool                        `json:"chat_is_forum,omitempty"`
	ChatHasUsername             *bool                        `json:"chat_has_username,omitempty"`
	ChatIsCreated               *bool                        `json:"chat_is_created,omitempty"`
	UserAdministratorRights     *ChatAdministratorRights     `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights      *ChatAdministratorRights     `json:"bot_administrator_rights,omitempty"`
	BotIsMember                 *bool                        `json:"bot_is_member,omitempty"`
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	RequestTitle                *bool                        `json:"request_title,omitempty"`
	RequestUsername             *bool                        `json:"request_username,omitempty"`
	RequestPhoto                *bool                        `json:"request_photo,omitempty"`
}

func (k KeyboardButtonRequestChat) Validate() error {
	if err := assertions.ParamNotEmpty(k.RequestId, "RequestId"); err != nil {
		return err
	}
	return nil
}

type KeyboardButtonRequestUser struct {
	RequestId     string `json:"request_id"`
	UserIsBot     bool   `json:"user_is_bot"`
	UserIsPremium bool   `json:"user_is_premium"`
}

func (k KeyboardButtonRequestUser) Validate() error {
	if err := assertions.ParamNotEmpty(k.RequestId, "RequestId"); err != nil {
		return err
	}
	return nil
}

type KeyboardButtonRequestUsers struct {
	RequestId       string `json:"request_id"`
	UserIsBot       *bool  `json:"user_is_bot,omitempty"`
	UserIsPremium   *bool  `json:"user_is_premium,omitempty"`
	MaxQuantity     *int   `json:"max_quantity,omitempty"`
	RequestName     *bool  `json:"request_name,omitempty"`
	RequestUsername *bool  `json:"request_username,omitempty"`
	RequestPhoto    *bool  `json:"request_photo,omitempty"`
}

func (k KeyboardButtonRequestUsers) Validate() error {
	if err := assertions.ParamNotEmpty(k.RequestId, "RequestId"); err != nil {
		return err
	}
	if *k.MaxQuantity < 1 || *k.MaxQuantity > 10 {
		return fmt.Errorf("MaxQuantity parameter must be between 1 and 10")
	}
	return nil
}

type KeyboardRow []KeyboardButton

func (k KeyboardRow) Validate() error {
	for _, button := range k {
		if err := button.Validate(); err != nil {
			return err
		}
	}
	return nil
}
