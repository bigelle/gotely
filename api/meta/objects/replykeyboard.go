package objects

import (
	"fmt"

	"github.com/bigelle/utils.go/ensure"
)

type InlineKeyboardMarkup struct {
	Keyboard []InlineKeyboardRow
}

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
	Text                         string
	Url                          string
	CallbackData                 string
	CallbackGame                 *CallbackGame
	SwitchInlineQuery            string
	SwitchInlineQueryCurrentChat string
	Pay                          bool
	LoginUrl                     *LoginUrl
	WebApp                       *WebAppInfo
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat
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
	if !ensure.NotNil(b.CallbackGame) {
		if err := (*b.CallbackGame).Validate(); err != nil {
			return err
		}
	}

	return nil
}

type SwitchInlineQueryChosenChat struct {
	RequestId         string
	AllowUserChats    bool
	AllowBotChats     bool
	AllowGroupChats   bool
	AllowChannelChats bool
}

//TODO: ReplyKeyboard*, Keyboard*
