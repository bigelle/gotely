package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal/assertions"
)

type InputMedia struct {
	InputMediaInterface
}

type InputMediaInterface interface {
	SetInputMedia(media string, isNew bool)
	telego.Validator
}

func (i InputMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.InputMediaInterface)
}

func (i *InputMedia) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "animation":
		tmp := InputMediaAnimation{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "audio":
		tmp := InputMediaAudio{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "document":
		tmp := InputMediaDocument{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "photo":
		tmp := InputMediaPhoto{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "video":
		tmp := InputMediaVideo{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	default:
		return errors.New("type must be animation, audio, document, video or photo")
	}
	return nil
}

type InputMediaAnimation struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`
	Caption               *string          `json:"caption,omitempty"`
	ParseMode             *string          `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"`
	Width                 *int             `json:"width,omitempty"`
	Height                *int             `json:"height,omitempty"`
	Duration              *int             `json:"duration,omitempty"`
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`
	isNew                 bool             `json:"-"`
}

func (i *InputMediaAnimation) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaAnimation) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}
	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InputMediaAudio struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumbnail       *InputFile       `json:"thumbnail,omitempty"`
	Caption         *string          `json:"caption,omitempty"`
	ParseMode       *string          `json:"parse_mode,omitempty"`
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int             `json:"duration,omitempty"`
	Performer       *string          `json:"performer,omitempty"`
	Title           *string          `json:"title,omitempty"`
	isNew           bool             `json:"-"`
}

func (i *InputMediaAudio) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaAudio) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}
	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InputMediaDocument struct {
	Type                        string           `json:"type"`
	Media                       string           `json:"media"`
	Thumbnail                   *InputFile       `json:"thumbnail,omitempty"`
	Caption                     *string          `json:"caption,omitempty"`
	ParseMode                   *string          `json:"parse_mode,omitempty"`
	CaptionEntities             *[]MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool            `json:"disable_content_type_detection,omitempty"`
	isNew                       bool             `json:"-"`
}

func (i *InputMediaDocument) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaDocument) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}
	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InputMediaPhoto struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Caption               *string          `json:"caption,omitempty"`
	ParseMode             *string          `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`
	isNew                 bool             `json:"-"`
}

func (i *InputMediaPhoto) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}

type InputMediaVideo struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`
	Caption               *string          `json:"caption,omitempty"`
	ParseMode             *string          `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"`
	Width                 *int             `json:"width,omitempty"`
	Height                *int             `json:"height,omitempty"`
	Duration              *int             `json:"duration,omitempty"`
	SupportsStreaming     *bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`
	isNew                 bool             `json:"-"`
}

func (i *InputMediaVideo) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaVideo) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}

type InputPaidMedia struct {
	InputPaidMediaInterface `json:"input_paid_media_interface"`
}

type InputPaidMediaInterface interface {
	SetInputPaidMedia(media string, isNew bool)
	telego.Validator
}

func (i InputPaidMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.InputPaidMediaInterface)
}

func (i *InputPaidMedia) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "photo":
		tmp := InputPaidMediaPhoto{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputPaidMediaInterface = &tmp
	case "video":
		tmp := InputPaidMediaVideo{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputPaidMediaInterface = &tmp
	default:
		return errors.New("type must be photo or video")
	}
	return nil
}

type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	isNew bool   `json:"-"`
}

func (i InputPaidMediaPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}

func (i *InputPaidMediaPhoto) SetInputPaidMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

type InputPaidMediaVideo struct {
	Type              string     `json:"type"`
	Media             string     `json:"media"`
	Thumbnail         *InputFile `json:"thumbnail,omitempty"`
	Width             *int       `json:"width,omitempty"`
	Height            *int       `json:"height,omitempty"`
	Duration          *int       `json:"duration,omitempty"`
	SupportsStreaming *bool      `json:"supports_streaming,omitempty"`
	isNew             bool       `json:"-"`
}

func (i *InputPaidMediaVideo) SetInputPaidMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputPaidMediaVideo) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return fmt.Errorf(
				"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}
