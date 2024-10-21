package types

import (
	"encoding/json"
	"fmt"
)

type BackgroundType struct {
	BackgroundTypeInterface
}

type BackgroundTypeInterface interface {
	backgroundTypeContract()
}

func (b *BackgroundType) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "fill":
		b.BackgroundTypeInterface = new(BackgroundTypeFill)
	case "chat_theme":
		b.BackgroundTypeInterface = new(BackgroundTypeChatTheme)
	case "pattern":
		b.BackgroundTypeInterface = new(BackgroundTypePattern)
	case "wallpaper":
		b.BackgroundTypeInterface = new(BackgroundTypeWallpaper)
	default:
		return fmt.Errorf("Type must be fill, chat_theme, pattern or wallpaper")
	}

	return json.Unmarshal(raw.Attributes, b.BackgroundTypeInterface)
}

type BackgroundTypeFill struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

func (b BackgroundTypeFill) backgroundTypeContract() {}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

func (b BackgroundTypeChatTheme) backgroundTypeContract() {}

type BackgroundTypeWallpaper struct {
	Type             string   `json:"type"`
	Document         Document `json:"document"`
	DarkThemeDimming int      `json:"dark_theme_dimming"`
	IsBlurred        *bool    `json:"is_blurred,omitempty"`
	IsMoving         *bool    `json:"is_moving,omitempty"`
}

func (b BackgroundTypeWallpaper) backgroundTypeContract() {}

type BackgroundTypePattern struct {
	Type       string         `json:"type"`
	Document   Document       `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int            `json:"intensity"`
	IsInverted *bool          `json:"is_inverted,omitempty"`
	IsMoving   *bool          `json:"is_moving,omitempty"`
}

func (b BackgroundTypePattern) backgroundTypeContract() {}

type BackgroundFill struct {
	BackgroundFillInterface
}

type BackgroundFillInterface interface {
	backgroundFillContract()
}

func (b *BackgroundFill) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "freeform_gradient":
		b.BackgroundFillInterface = new(BackgroundFillFreeformGradient)
	case "gradient":
		b.BackgroundFillInterface = new(BackgroundFillGradient)
	case "solid":
		b.BackgroundFillInterface = new(BackgroundFillSolid)
	default:
		return fmt.Errorf("Type must be freeform_gradient, gradient or solid")
	}
	return json.Unmarshal(raw.Attributes, b.BackgroundFillInterface)
}

type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`
	Colors []int  `json:"colors"`
}

func (b BackgroundFillFreeformGradient) backgroundFillContract() {}

type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

func (b BackgroundFillGradient) backgroundFillContract() {}

type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

func (b BackgroundFillSolid) backgroundFillContract() {}
