package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bigelle/tele.go/interfaces"
	"github.com/bigelle/tele.go/internal/assertions"
)

type InlineQuery struct {
	Id       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType *string   `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            User      `json:"from"`
	Query           string    `json:"query"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId *string   `json:"inline_message_id,omitempty"`
}

type InlineQueryResult struct {
	InlineQueryResultInterface
}

type InlineQueryResultInterface interface {
	interfaces.Validator
	inlineQueryResultContract() //NOTE: maybe should do something special
}

func (i *InlineQueryResult) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "article":
		i.InlineQueryResultInterface = new(InlineQueryResultArticle)
	case "audio":
		i.InlineQueryResultInterface = new(InlineQueryResultAudio)
	case "contact":
		i.InlineQueryResultInterface = new(InlineQueryResultContact)
	case "game":
		i.InlineQueryResultInterface = new(InlineQueryResultGame)
	case "document":
		i.InlineQueryResultInterface = new(InlineQueryResultDocument)
	case "gif":
		i.InlineQueryResultInterface = new(InlineQueryResultGif)
	case "location":
		i.InlineQueryResultInterface = new(InlineQueryResultLocation)
	case "mpeg4_gif":
		i.InlineQueryResultInterface = new(InlineQueryResultMpeg4Gif)
	case "photo":
		i.InlineQueryResultInterface = new(InlineQueryResultPhoto)
	case "venue":
		i.InlineQueryResultInterface = new(InlineQueryResultVenue)
	case "video":
		i.InlineQueryResultInterface = new(InlineQueryResultVideo)
	case "voice":
		i.InlineQueryResultInterface = new(InlineQueryResultVoice)
	default:
		return fmt.Errorf(
			"Type must be article, audio, contact, game, document, gif, location, mpeg4_gif, photo, venue, video or voice",
		)
	}
	return json.Unmarshal(raw.Attributes, i.InlineQueryResultInterface)
}

type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 *string               `json:"url,omitempty"`
	HideUrl             *bool                 `json:"hide_url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultArticle) inlineQueryResultContract() {}

func (i InlineQueryResultArticle) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if err := i.InputMessageContent.Validate(); err != nil {
		return err
	}
	if i.ReplyMarkup != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	AudioUrl            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int                  `json:"audio_duration,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     *[]MessageEntity      `json:"caption_entities,omitempty"`
}

func (i InlineQueryResultAudio) inlineQueryResultContract() {}

func (i InlineQueryResultAudio) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.AudioUrl, "AudioUrl"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if *i.ParseMode != "" && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("Parse mode can't be enabled if Entities are provided")
	}
	if i.InputMessageContent != nil {
		if err := (*i.InputMessageContent).Validate(); err != nil {
			return err
		}
	}
	if i.ReplyMarkup != nil {
		if err := (*i.ReplyMarkup).Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
	VCard               *string               `json:"v_card,omitempty"`
}

func (i InlineQueryResultContact) inlineQueryResultContract() {}

func (i InlineQueryResultContact) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.PhoneNumber, "PhoneNumber"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.FirstName, "FirstName"); err != nil {
		return err
	}
	if i.InputMessageContent != nil {
		if err := (*i.InputMessageContent).Validate(); err != nil {
			return err
		}
	}
	if i.ReplyMarkup != nil {
		if err := (*i.ReplyMarkup).Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	MimeType            string                `json:"mime_type"`
	DocumentUrl         string                `json:"document_url"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     *[]MessageEntity      `json:"caption_entities,omitempty"`
}

func (i InlineQueryResultDocument) inlineQueryResultContract() {}

func (i InlineQueryResultDocument) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.MimeType, "MimeType"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.DocumentUrl, "DocumentUrl"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("Parse mode can't be enabled if Entities are provided")
	}
	if i.InputMessageContent != nil {
		if err := (*i.InputMessageContent).Validate(); err != nil {
			return err
		}
	}
	if i.ReplyMarkup != nil {
		if err := (*i.ReplyMarkup).Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	Id            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (i InlineQueryResultGame) inlineQueryResultContract() {}

func (i InlineQueryResultGame) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.GameShortName, "GameShortName"); err != nil {
		return err
	}
	if i.ReplyMarkup != nil {
		if err := (*i.ReplyMarkup).Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultGif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	GifUrl                string                `json:"gif_url"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	GifWidth              *int                  `json:"gif_width,omitempty"`
	GifHeight             *int                  `json:"gif_height,omitempty"`
	ThumbnailMimeType     *string               `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	GifDuration           *int                  `json:"gif_duration,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
}

func (i InlineQueryResultGif) inlineQueryResultContract() {}

func (i InlineQueryResultGif) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.GifUrl, "GifUrl"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.ThumbnailUrl, "ThumbnailUrl"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	Id                   string                `json:"id"`
	Latitude             *float32              `json:"latitude"`
	Longtitude           *float32              `json:"longtitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   *float32              `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int                  `json:"live_period,omitempty"`
	Heading              *int                  `json:"heading,omitempty"`
	ProximityAlertRadius *int                  `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl         *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      *int                  `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultLocation) inlineQueryResultContract() {}

func (i InlineQueryResultLocation) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if i.Latitude == nil {
		return fmt.Errorf("Latitude parameter can't be nil")
	}
	if i.Longtitude == nil {
		return fmt.Errorf("Longtitude parameter can't be nil")
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	Mpeg4Url              string                `json:"mpeg4_url"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	Mpeg4Width            *int                  `json:"mpeg4_width,omitempty"`
	Mpeg4Height           *int                  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         *int                  `json:"mpeg4_duration,omitempty"`
	ThumbnailMimeType     *string               `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultMpeg4Gif) inlineQueryResultContract() {}

func (i InlineQueryResultMpeg4Gif) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Mpeg4Url, "Mpeg4Url"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.ThumbnailUrl, "ThumbnailUrl"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be used if Entities are provided")
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	PhotoUrl              string                `json:"photo_url"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	PhotoWidth            *int                  `json:"photo_width,omitempty"`
	PhotoHeight           *int                  `json:"photo_height,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultPhoto) inlineQueryResultContract() {}

func (i InlineQueryResultPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.PhotoUrl, "PhotoUrl"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.ThumbnailUrl, "ThumbnailUrl"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Latitude            *float32              `json:"latitude"`
	Longtitude          *float32              `json:"longtitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareId        *string               `json:"foursquare_id,omitempty"`
	FourSquareType      *string               `json:"four_square_type,omitempty"`
	GooglePlaceId       *string               `json:"google_place_id,omitempty"`
	GooglePlaceType     *string               `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultVenue) inlineQueryResultContract() {}

func (i InlineQueryResultVenue) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Address, "Address"); err != nil {
		return err
	}
	if i.Latitude == nil {
		return fmt.Errorf("Latitude parameter can't be nil")
	}
	if i.Longtitude == nil {
		return fmt.Errorf("Longtitude parameter can't be nil")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil

}

type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	VideoUrl              string                `json:"video_url"`
	MimeType              string                `json:"mime_type"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	VideoWidth            *int                  `json:"video_width,omitempty"`
	VideoHeight           *int                  `json:"video_height,omitempty"`
	VideoDuration         *int                  `json:"video_duration,omitempty"`
	Description           *string               `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVideo) inlineQueryResultContract() {}

func (i InlineQueryResultVideo) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.VideoUrl, "VideoUrl"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.MimeType, "MimeType"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.ThumbnailUrl, "ThumbnailUrl"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil

}

type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VoiceUrl            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     *[]MessageEntity      `json:"caption_entities,omitempty"`
	VoiceDuration       *int                  `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVoice) inlineQueryResultContract() {}

func (i InlineQueryResultVoice) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.VoiceUrl, "VoiceUrl"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Title, "Title"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCached struct {
	InlineQueryResultCachedInterface
}

type InlineQueryResultCachedInterface interface {
	interfaces.Validator
	inlineQueryResultCachedContract()
}

func (i *InlineQueryResultCached) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "photo":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedPhoto)
	case "gif":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedGif)
	case "mpeg4_gif":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedMpeg4Gif)
	case "sticker":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedSticker)
	case "document":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedDocument)
	case "video":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedVideo)
	case "voice":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedVoice)
	case "audio":
		i.InlineQueryResultCachedInterface = new(InlineQueryResultCachedAudio)
	}
	return json.Unmarshal(raw.Attributes, &i.InlineQueryResultCachedInterface)
}

type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	PhotoFileId           string                `json:"photo_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedPhoto) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.PhotoFileId, "PhotoUrl"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	GifFileId             string                `json:"gif_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedGif) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedGif) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.GifFileId, "GifFileId"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	Mpeg4FileId           string                `json:"mpeg_4_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedMpeg4Gif) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedMpeg4Gif) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.Mpeg4FileId, "Mpeg4FileId"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	StickerFileId       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedSticker) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedSticker) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.StickerFileId, "StickerFileId"); err != nil {
		return err
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	VideoFileId           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedVideo) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedVideo) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.VideoFileId, "VideoFileId"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedVoice struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	VoiceFileId           string                `json:"voice_file_id"`
	Title                 string                `json:"title"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedVoice) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedVoice) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.VoiceFileId, "VoiceFileId"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedAudio struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	AudioFileId           string                `json:"audio_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedAudio) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedAudio) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.AudioFileId, "AudioFileId"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InlineQueryResultCachedDocument struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	DocumentFileId        string                `json:"document_file_id"`
	Title                 string                `json:"title"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedDocument) inlineQueryResultCachedContract() {}

func (i InlineQueryResultCachedDocument) Validate() error {
	if err := assertions.ParamNotEmpty(i.Id, "Id"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(i.DocumentFileId, "DocumentFileId"); err != nil {
		return err
	}
	if !assertions.IsStringEmpty(*i.ParseMode) && len(*i.CaptionEntities) != 0 {
		return fmt.Errorf("ParseMode can't be enabled if Entities are provided")
	}
	if i.ReplyMarkup != nil {
		if err := i.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InputMessageContent struct {
	InputMessageContentInterface
}

type InputMessageContentInterface interface {
	interfaces.Validator
	inputMessageContentContract()
}

func (i *InputMessageContent) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if _, ok := raw["message_text"]; ok {
		i.InputMessageContentInterface = new(InputTextMessageContent)
		return json.Unmarshal(data, &i.InputMessageContentInterface)
	}
	if _, ok := raw["address"]; ok {
		i.InputMessageContentInterface = new(InputVenueMessageContent)
		return json.Unmarshal(data, &i.InputMessageContentInterface)
	}
	if _, ok := raw["latitude"]; ok {
		i.InputMessageContentInterface = new(InputLocationMessageContent)
		return json.Unmarshal(data, &i.InputMessageContentInterface)
	}
	if _, ok := raw["phone_number"]; ok {
		i.InputMessageContentInterface = new(InputContactMessageContent)
		return json.Unmarshal(data, &i.InputMessageContentInterface)
	}
	if _, ok := raw["provider_token"]; ok {
		i.InputMessageContentInterface = new(InputInvoiceMessageContent)
		return json.Unmarshal(data, &i.InputMessageContentInterface)
	}
	return fmt.Errorf("Unrecognized type: %T", i.InputMessageContentInterface)
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	VCard       string `json:"v_card"`
}

func (i InputContactMessageContent) inputMessageContentContract() {}

func (i InputContactMessageContent) Validate() error {
	if i.PhoneNumber == "" || strings.TrimSpace(i.PhoneNumber) == "" {
		return fmt.Errorf("phone number can't be empty")
	}
	if i.FirstName == "" || strings.TrimSpace(i.FirstName) == "" {
		return fmt.Errorf("firstname can't be empty")
	}
	return nil
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              *int           `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       *[]int         `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoSize                 *int           `json:"photo_size,omitempty"`
	PhotoWidth                *int           `json:"photo_width,omitempty"`
	PhotoHeight               *int           `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
	ProviderToken             *string        `json:"provider_token,omitempty"`
}

func (i InputInvoiceMessageContent) inputMessageContentContract() {}

func (i InputInvoiceMessageContent) Validate() error {
	if i.Title == "" || strings.TrimSpace(i.Title) == "" || len(i.Title) > 32 {
		return fmt.Errorf("title parameter must be between 1 and 32 characters")
	}
	if i.Description == "" || strings.TrimSpace(i.Description) == "" || len(i.Description) > 255 {
		return fmt.Errorf("Description parameter must be between 1 and 255 characters")
	}
	if i.Payload == "" || strings.TrimSpace(i.Payload) == "" || len(i.Payload) > 128 {
		return fmt.Errorf("payload parameter must be between 1 and 128 characters")
	}
	if i.Currency == "" {
		return fmt.Errorf("currency parameter can't be empty")
	}
	if len(i.Prices) == 0 {
		return fmt.Errorf("prices parameter can't be empty")
	}
	for _, label := range i.Prices {
		if err := label.Validate(); err != nil {
			return err
		}
	}
	if len(*i.SuggestedTipAmounts) != 0 || len(*i.SuggestedTipAmounts) > 4 {
		return fmt.Errorf("only up to 4 suggested tip amounts are allowed")
	}
	return nil
}

type InputLocationMessageContent struct {
	Latitude              float64  `json:"latitude"`
	Longtitude            float64  `json:"longtitude"`
	LivePeriod            *int     `json:"live_period,omitempty"`
	HorizontalAccuracy    *float64 `json:"horizontal_accuracy,omitempty"`
	Heading               *int     `json:"heading,omitempty"`
	ProximityAlertRaidius *int     `json:"proximity_alert_raidius,omitempty"`
}

func (i InputLocationMessageContent) inputMessageContentContract() {}

func (i InputLocationMessageContent) Validate() error {
	if i.LivePeriod != nil {
		if (*i.LivePeriod < 60 || *i.LivePeriod > 86400) && *i.LivePeriod != 0x7FFFFFFF {
			return fmt.Errorf("LivePeriod parameter must be between 60 and 86400 or be 0x7FFFFFFF")
		}
	}
	if i.HorizontalAccuracy != nil {
		if *i.HorizontalAccuracy < 0 || *i.HorizontalAccuracy > 1500 {
			return fmt.Errorf("Horizontal accuracy must be between 0 and 1500")
		}
	}
	if i.Heading != nil {
		if *i.Heading < 1 || *i.Heading > 360 {
			return fmt.Errorf("Heading Accuracy must be between 1 and 360")
		}
	}
	if i.ProximityAlertRaidius != nil {
		if *i.ProximityAlertRaidius < 1 || *i.ProximityAlertRaidius > 100000 {
			return fmt.Errorf(
				"Approaching Notification distance parameter must be between 1 and 100000",
			)
		}
	}
	return nil
}

type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          *string             `json:"parse_mode,omitempty"`
	Entities           *[]MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (i InputTextMessageContent) inputMessageContentContract() {}

func (i InputTextMessageContent) Validate() error {
	if i.MessageText == "" || strings.TrimSpace(i.MessageText) == "" {
		return fmt.Errorf("MessageText parameter can't be empty")
	}
	if *i.ParseMode != "" && len(*i.Entities) != 0 {
		return fmt.Errorf("parse mode can't be enabled if entities are provided")
	}
	if i.LinkPreviewOptions != nil {
		if err := i.LinkPreviewOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InputVenueMessageContent struct {
	Latitude        float32 `json:"latitude"`
	Longtitude      float32 `json:"longtitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareId    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceId   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

func (i InputVenueMessageContent) inputMessageContentContract() {}

func (i InputVenueMessageContent) Validate() error {
	if i.Title == "" || strings.TrimSpace(i.Title) == "" {
		return fmt.Errorf("title parameter can't be empty")
	}
	if i.Address == "" || strings.TrimSpace(i.Address) == "" {
		return fmt.Errorf("address parameter can't be empty")
	}
	return nil
}
