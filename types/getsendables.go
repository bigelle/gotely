package types

import (
	"encoding/json"
	"errors"

	"github.com/bigelle/tele.go/assertions"
)

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Performer    *string    `json:"performer,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
}

type BackgroundType struct {
	BackgroundTypeInterface
}

type BackgroundTypeInterface interface {
	backgroundTypeContract()
}

func (b BackgroundType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BackgroundTypeInterface)
}

func (b *BackgroundType) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "fill":
		tmp := BackgroundTypeFill{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	case "chat_theme":
		tmp := BackgroundTypeChatTheme{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	case "pattern":
		tmp := BackgroundTypePattern{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	case "wallpaper":
		tmp := BackgroundTypeWallpaper{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	default:
		return errors.New("type must be fill, chat_theme, pattern or wallpaper")
	}

	return nil
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

func (b BackgroundFill) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BackgroundFillInterface)
}

func (b *BackgroundFill) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "freeform_gradient":
		tmp := BackgroundFillFreeformGradient{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundFillInterface = tmp
	case "gradient":
		tmp := BackgroundFillGradient{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundFillInterface = tmp
	case "solid":
		tmp := BackgroundFillSolid{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundFillInterface = tmp
	default:
		return errors.New("type must be freeform_gradient, gradient or solid")
	}
	return nil
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

type ReplyParameters struct {
	MessageId                int              `json:"message_id"`
	ChatId                   *string          `json:"chat_id,omitempty"`
	AllowSendingWithoutReply *bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    *string          `json:"quote,omitempty"`
	QuoteParseMode           *string          `json:"quote_parse_mode,omitempty"`
	QuoteEntities            *[]MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            *int             `json:"quote_position,omitempty"`
}

func (r ReplyParameters) Validate() error {
	if err := assertions.ParamNotEmpty(*r.ChatId, "ChatId"); err != nil {
		return err
	}
	return nil
}
