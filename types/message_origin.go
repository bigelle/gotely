package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/internal"
)

type MessageOrigin struct {
	MessageOriginInterface
}

type MessageOriginInterface interface {
	messageOriginContract()
	internal.Validator
}

func (m MessageOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.MessageOriginInterface)
}

func (m *MessageOrigin) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "user":
		tmp := MessageOriginUser{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	case "hidden_user":
		tmp := MessageOriginHiddenUser{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	case "chat":
		tmp := MessageOriginChat{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	case "channel":
		tmp := MessageOriginChannel{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	default:
		return errors.New("type must be user, hidden_user, chat or channel")
	}
}

type MessageOriginChannel struct {
	Type            string  `json:"type"`
	Date            *int    `json:"date"`
	Chat            *Chat   `json:"chat"`
	MessageId       *int    `json:"message_id"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

func (m MessageOriginChannel) messageOriginContract() {}

func (m MessageOriginChannel) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return fmt.Errorf("Date parameter can't be empty")
	}
	if m.Chat == nil {
		return fmt.Errorf("Chat parameter can't be empty")
	}
	if m.MessageId == nil {
		return fmt.Errorf("MessageId parameter can't be empty")
	}
	return nil
}

type MessageOriginChat struct {
	Type            string  `json:"type"`
	Date            *int    `json:"date"`
	SenderChat      *Chat   `json:"sender_chat"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

func (m MessageOriginChat) messageOriginContract() {}

func (m MessageOriginChat) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return fmt.Errorf("Date parameter can't be empty")
	}
	if m.SenderChat == nil {
		return fmt.Errorf("Chat parameter can't be empty")
	}
	return nil
}

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           *int   `json:"date"`
	SenderUsername string `json:"sender_username"`
}

func (m MessageOriginHiddenUser) messageOriginContract() {}

func (m MessageOriginHiddenUser) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return fmt.Errorf("Date parameter can't be empty")
	}
	if err := assertions.ParamNotEmpty(m.SenderUsername, "SenderUsername"); err != nil {
		return err
	}
	return nil
}

type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       *int   `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

func (m MessageOriginUser) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return fmt.Errorf("Date parameter can't be empty")
	}
	if m.SenderUser == nil {
		return fmt.Errorf("SenderUser parameter can't be empty")
	}
	return nil
}

func (m MessageOriginUser) messageOriginContract() {}
