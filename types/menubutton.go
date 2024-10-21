package types

import (
	"encoding/json"
	"fmt"

	"github.com/bigelle/tele.go/interfaces"
	"github.com/bigelle/tele.go/internal/assertions"
)

type MenuButton struct {
	MenuButtonInterface
}

type MenuButtonInterface interface {
	menuButtonContract()
	interfaces.Validator
}

func (m *MenuButton) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "commands":
		m.MenuButtonInterface = new(MenuButtonCommands)
	case "default":
		m.MenuButtonInterface = new(MenuButtonDefault)
	case "web_app":
		m.MenuButtonInterface = new(MenuButtonWebApp)
	}
	return json.Unmarshal(raw.Attributes, m.MenuButtonInterface)
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
