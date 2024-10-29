package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bigelle/tele.go/internal/assertions"
)

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (b BotCommand) Validate() error {
	if assertions.IsStringEmpty(b.Command) {
		return fmt.Errorf("command parameter can't be empty")
	}
	if assertions.IsStringEmpty(b.Description) {
		return fmt.Errorf("description parameter can't be empty")
	}
	return nil
}

type BotCommandScope struct {
	BotCommandScopeInterface
}

type BotCommandScopeInterface interface {
	botCommandScopeContract()
}

func (b BotCommandScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BotCommandScopeInterface)
}

func (b *BotCommandScope) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "all_chat_administrators":
		tmp := BotCommandScopeAllChatAdministrators{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "all_group_chats":
		tmp := BotCommandScopeAllGroupChats{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "all_private_chats":
		tmp := BotCommandScopeAllPrivateChats{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "chat":
		tmp := BotCommandScopeChat{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "chat_administrators":
		tmp := BotCommandScopeChatAdministrators{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "chat_member":
		tmp := BotCommandScopeChatMember{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "default":
		tmp := BotCommandScopeDefault{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	default:
		return errors.New(
			"type must be all_chat_administrators, all_group_chats, all_private_chats, chat, chat_administrators, chat_member or default",
		)
	}
	return nil
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllChatAdministrators) botCommandScopeContract() {}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllGroupChats) botCommandScopeContract() {}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllPrivateChats) botCommandScopeContract() {}

type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
}

func (b BotCommandScopeChat) botCommandScopeContract() {}

func (b BotCommandScopeChat) Validate() error {
	if assertions.IsStringEmpty(b.ChatId) {
		return fmt.Errorf("ChatId parameter can't be empty")
	}
	return nil
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
}

func (b BotCommandScopeChatAdministrators) botCommandScopeContract() {}

func (b BotCommandScopeChatAdministrators) Validate() error {
	if assertions.IsStringEmpty(b.ChatId) {
		return fmt.Errorf("ChatId parameter can't be empty")
	}
	return nil
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
	UserId int64  `json:"user_id"`
}

func (b BotCommandScopeChatMember) botCommandScopeContract() {}

func (b BotCommandScopeChatMember) Validate() error {
	if assertions.IsStringEmpty(b.ChatId) {
		return fmt.Errorf("ChatId parameter can't be empty")
	}
	if b.UserId == 0 {
		return fmt.Errorf("UserId parameter can't be empty")
	}
	return nil
}

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

func (b BotCommandScopeDefault) botCommandScopeContract() {}
