package types

import (
	"encoding/json"
	"errors"

	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/internal"
)

type MenuButton struct {
	MenuButtonInterface
}

type MenuButtonInterface interface {
	menuButtonContract()
	internal.Validator
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
	if err := assertions.ParamNotEmpty(m.Text, "text"); err != nil {
		return err
	}
	if err := m.WebAppInfo.Validate(); err != nil {
		return err
	}
	return nil
}
