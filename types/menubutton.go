package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal/assertions"
)

type MenuButton struct {
	MenuButtonInterface
}

type MenuButtonInterface interface {
	menuButtonContract()
	telego.Validator
}

func (m MenuButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.MenuButtonInterface)
}

func (m *MenuButton) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "commands":
		tmp := MenuButtonCommands{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MenuButtonInterface = tmp
	case "web_app":
		tmp := MenuButtonWebApp{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MenuButtonInterface = tmp
	case "default":
		tmp := MenuButtonDefault{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MenuButtonInterface = tmp
	default:
		return errors.New("unknow type " + raw.Type + ", type must be commands, web_app or default")
	}
	return nil
}

type MenuButtonCommands struct {
	Type string `json:"type"`
}

func (m MenuButtonCommands) Validate() error {
	return assertions.ParamNotEmpty(m.Type, "Type")
}

func (m MenuButtonCommands) menuButtonContract() {}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

func (m MenuButtonDefault) Validate() error {
	return assertions.ParamNotEmpty(m.Type, "Type")
}

func (m MenuButtonDefault) menuButtonContract() {}

type MenuButtonWebApp struct {
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	WebAppInfo WebAppInfo `json:"web_app_info"`
}

func (m MenuButtonWebApp) menuButtonContract() {}

func (m MenuButtonWebApp) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if assertions.IsStringEmpty(m.Text) {
		return fmt.Errorf("text parameter can't be empty")
	}
	if err := m.WebAppInfo.Validate(); err != nil {
		return err
	}
	return nil
}
