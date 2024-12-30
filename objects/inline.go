package objects

import (
	"fmt"
	"regexp"
	"strings"
)

// This object represents an incoming inline query. When the user sends an empty query,
// your bot could return some default or trending results.
type InlineQuery struct {
	//Unique identifier for this query
	Id string `json:"id"`
	//Sender
	From User `json:"from"`
	//Text of the query (up to 256 characters)
	Query string `json:"query"`
	//Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
	//Optional. Type of the chat from which the inline query was sent.
	//Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”.
	//The chat type should be always known for requests sent from official clients and most third-party clients,
	//unless the request was sent from a secret chat
	ChatType *string `json:"chat_type,omitempty"`
	//Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// This object represents a button to be shown above inline query results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	//Label text on the button
	Text string
	//Optional. Description of the Web App that will be launched when the user presses the button.
	//The Web App will be able to switch back to the inline mode using the method switchInlineQuery inside the Web App.
	WebApp *WebAppInfo
	//Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button.
	//1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
	//
	//Example: An inline bot that sends YouTube videos can ask the user to
	//connect the bot to their YouTube account to adapt search results accordingly.
	//To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any.
	//The user presses the button, switches to a private chat with the bot and, in doing so,
	//passes a start parameter that instructs the bot to return an OAuth link. Once done,
	//the bot can offer a switch_inline button so that the user can easily return to the
	//chat where they wanted to use the bot's inline capabilities.
	StartParameter *string
}

var allowed_startparameter = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func (i InlineQueryResultsButton) Validate() error {
	if strings.TrimSpace(i.Text) == "" {
		return ErrInvalidParam("text parameter can't be empty")
	}
	if i.WebApp != nil {
		if err := i.WebApp.Validate(); err != nil {
			return err
		}
	}
	if i.StartParameter != nil {
		if !allowed_startparameter.MatchString(*i.StartParameter) {
			return ErrInvalidParam("start_parameter parameter must contain only A-Z, a-z, 0-9, \"_\" and \"-\"")
		}
		if len(*i.StartParameter) < 1 || len(*i.StartParameter) > 64 {
			return ErrInvalidParam("start_parameter parameter must be between 1 and 64 characters")
		}
	}
	return nil
}

// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types
//
// - InlineQueryResultCachedAudio
//
// - InlineQueryResultCachedDocument
//
// - InlineQueryResultCachedGif
//
// - InlineQueryResultCachedMpeg4Gif
//
// - InlineQueryResultCachedPhoto
//
// - InlineQueryResultCachedSticker
//
// - InlineQueryResultCachedVideo
//
// - InlineQueryResultCachedVoice
//
// - InlineQueryResultArticle
//
// - InlineQueryResultAudio
//
// - InlineQueryResultContact
//
// - InlineQueryResultGame
//
// - InlineQueryResultDocument
//
// - InlineQueryResultGif
//
// - InlineQueryResultLocation
//
// - InlineQueryResultMpeg4Gif
//
// - InlineQueryResultPhoto
//
// - InlineQueryResultVenue
//
// - InlineQueryResultVideo
//
// - InlineQueryResultVoice
//
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
type InlineQueryResult interface {
	Validable
	GetInlineQueryResultType() string
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	//Type of the result, must be article
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	//Title of the result
	Title string `json:"title"`
	//Content of the message to be sent
	InputMessageContent InputMessageContent `json:"input_message_content"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. URL of the result
	Url *string `json:"url,omitempty"`
	//Optional. Pass True if you don't want the URL to be shown in the message
	HideUrl *bool `json:"hide_url,omitempty"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	//Optional. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	//Optional. Thumbnail width
	ThumbnailWidth *int `json:"thumbnail_width,omitempty"`
	//Optional. Thumbnail height
	ThumbnailHeight *int `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultArticle) GetInlineQueryResultType() string {
	return "article"
}

func (i InlineQueryResultArticle) Validate() error {
	if i.Type != "article" {
		return ErrInvalidParam("type must be \"article\"")
	}
	if len([]byte(i.Id)) > 64 {
		return ErrInvalidParam("id parameter must not be longer than 64 bytes")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
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

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	//Type of the result, must be photo
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	PhotoUrl string `json:"photo_url"`
	//URL of the thumbnail for the photo
	ThumbnailUrl string `json:"thumbnail_url"`
	//Optional. Width of the photo
	PhotoWidth *int `json:"photo_width,omitempty"`
	//Optional. Height of the photo
	PhotoHeight *int `json:"photo_height,omitempty"`
	//Optional. Title for the result
	Title *string `json:"title,omitempty"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	//Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the photo caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the photo
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultPhoto) GetInlineQueryResultType() string {
	return "photo"
}

func (i InlineQueryResultPhoto) Validate() error {
	if i.Type != "photo" {
		return ErrInvalidParam("type must be \"photo\"")
	}
	if len([]byte(i.Id)) > 64 {
		return ErrInvalidParam("id parameter can't be empty")
	}
	if strings.TrimSpace(i.PhotoUrl) == "" {
		return ErrInvalidParam("photo_url parameter can't be empty")
	}
	if strings.TrimSpace(i.ThumbnailUrl) == "" {
		return ErrInvalidParam("thumbnail_url parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be enabled if caption_entities are provided")
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

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	//Type of the result, must be gif
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid URL for the GIF file. File size must not exceed 1MB
	GifUrl string `json:"gif_url"`
	//Optional. Width of the GIF
	GifWidth *int `json:"gif_width,omitempty"`
	//Optional. Height of the GIF
	GifHeight *int `json:"gif_height,omitempty"`
	//Optional. Duration of the GIF in seconds
	GifDuration *int `json:"gif_duration,omitempty"`
	//URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url"`
	//Optional. MIME type of the thumbnail, must be one of
	//“image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType *string `json:"thumbnail_mime_type,omitempty"`
	//Optional. Title for the result
	Title *string `json:"title,omitempty"`
	//Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the caption. See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (i InlineQueryResultGif) GetInlineQueryResultType() string {
	return "gif"
}

func (i InlineQueryResultGif) Validate() error {
	if i.Type != "gif" {
		return ErrInvalidParam("type must be \"gif\"")
	}
	if len([]byte(i.Id)) > 64 {
		return ErrInvalidParam("id parameter must not be longer than 64 bytes")
	}
	if strings.TrimSpace(i.GifUrl) == "" {
		return ErrInvalidParam("gif_url parameter can't be empty")
	}
	if strings.TrimSpace(i.ThumbnailUrl) == "" {
		return ErrInvalidParam("thumbnail_url parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	//Type of the result, must be mpeg4_gif
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid URL for the MPEG4 file. File size must not exceed 1MB
	Mpeg4Url string `json:"mpeg4_url"`
	//Optional. Video width
	Mpeg4Width *int `json:"mpeg4_width,omitempty"`
	//Optional. Video height
	Mpeg4Height *int `json:"mpeg4_height,omitempty"`
	//Optional. Video duration in seconds
	Mpeg4Duration *int `json:"mpeg4_duration,omitempty"`
	//URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url"`
	//Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType *string `json:"thumbnail_mime_type,omitempty"`
	//Optional. Title for the result
	Title *string `json:"title,omitempty"`
	//Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the video animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultMpeg4Gif) GetInlineQueryResultType() string {
	return "mpeg4_gif"
}

func (i InlineQueryResultMpeg4Gif) Validate() error {
	if i.Type != "mpeg4_gif" {
		return ErrInvalidParam("type must be \"mpeg4_gif\"")
	}
	b := len([]byte(i.Id))
	if b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.Mpeg4Url) == "" {
		return ErrInvalidParam("mpeg4_url parameter can't be empty")
	}
	if strings.TrimSpace(i.ThumbnailUrl) == "" {
		return ErrInvalidParam("thumbnail_url parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube),
// you must replace its content using input_message_content.
type InlineQueryResultVideo struct {
	//Type of the result, must be video
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid URL for the embedded video player or video file
	VideoUrl string `json:"video_url"`
	//MIME type of the content of the video URL, “text/html” or “video/mp4”
	MimeType string `json:"mime_type"`
	//URL of the thumbnail (JPEG only) for the video
	ThumbnailUrl string `json:"thumbnail_url"`
	//Title for the result
	Title string `json:"title"`
	//Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the video caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Video width
	VideoWidth *int `json:"video_width,omitempty"`
	//Optional. Video height
	VideoHeight *int `json:"video_height,omitempty"`
	//Optional. Video duration in seconds
	VideoDuration *int `json:"video_duration,omitempty"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the video.
	//This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVideo) GetInlineQueryResultType() string {
	return "video"
}

func (i InlineQueryResultVideo) Validate() error {
	if i.Type != "video" {
		return ErrInvalidParam("type must be \"video\"")
	}
	b := len([]byte(i.Id))
	if b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.VideoUrl) == "" {
		return ErrInvalidParam("video_url parameter can't be empty")
	}
	if strings.TrimSpace(i.MimeType) == "" {
		return ErrInvalidParam("mime_type parameter can't be empty")
	}
	if strings.TrimSpace(i.ThumbnailUrl) == "" {
		return ErrInvalidParam("thumbnail_url parameter can't be empty")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse mode can't be enabled if Entities are provided")
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

// Represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	//Type of the result, must be audio
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid URL for the audio file
	AudioUrl string `json:"audio_url"`
	//Title
	Title string `json:"title"`
	//Optional. Caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the audio caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Performer
	Performer *string `json:"performer,omitempty"`
	//Optional. Audio duration in seconds
	AudioDuration *int `json:"audio_duration,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the audio
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultAudio) GetInlineQueryResultType() string {
	return "audio"
}

func (i InlineQueryResultAudio) Validate() error {
	if i.Type != "audio" {
		return ErrInvalidParam("type must be \"audio\"")
	}
	b := len([]byte(i.Id))
	if b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
	}
	if strings.TrimSpace(i.AudioUrl) == "" {
		return ErrInvalidParam("audio_url parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
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

// Represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	//Type of the result, must be voice
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid URL for the voice recording
	VoiceUrl string `json:"voice_url"`
	//Recording title
	Title string `json:"title"`
	//Optional. Caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the voice message caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Recording duration in seconds
	VoiceDuration *int `json:"voice_duration,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the voice recording
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVoice) GetInlineQueryResultType() string {
	return "voice"
}

func (i InlineQueryResultVoice) Validate() error {
	if i.Type != "voice" {
		return ErrInvalidParam("type must be \"voice\"")
	}
	b := len([]byte(i.Id))
	if b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.VoiceUrl) == "" {
		return ErrInvalidParam("voice_url parameter can't be empty")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse mode can't be enabled if Entities are provided")
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

// Represents a link to a file. By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	//Type of the result, must be document
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//Title for the result
	Title string `json:"title"`
	//Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the document caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//A valid URL for the file
	DocumentUrl string `json:"document_url"`
	//MIME type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	//Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	//Optional. Thumbnail width
	ThumbnailWidth *int `json:"thumbnail_width,omitempty"`
	//Optional. Thumbnail height
	ThumbnailHeight *int `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultDocument) GetInlineQueryResultType() string {
	return "document"
}

func (i InlineQueryResultDocument) Validate() error {
	if i.Type != "document" {
		return ErrInvalidParam("type must be \"document\"")
	}
	b := len([]byte(i.Id))
	if b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
	}
	if strings.TrimSpace(i.MimeType) == "" {
		return ErrInvalidParam("mime_type parameter can't be empty")
	}
	if strings.TrimSpace(i.DocumentUrl) == "" {
		return ErrInvalidParam("document_url parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
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

// Represents a location on a map. By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the location
type InlineQueryResultLocation struct {
	//Type of the result, must be location
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	//Location latitude in degrees
	Latitude *float32 `json:"latitude"`
	//Location longitude in degrees
	Longtitude *float32 `json:"longtitude"`
	//Location title
	Title string `json:"title"`
	//Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float32 `json:"horizontal_accuracy,omitempty"`
	//Optional. Period in seconds during which the location can be updated, should be between 60 and 86400,
	//or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int `json:"live_period,omitempty"`
	//Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int `json:"heading,omitempty"`
	//Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters.
	//Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int `json:"proximity_alert_radius,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the location
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	//Optional. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	//Optional. Thumbnail width
	ThumbnailWidth *int `json:"thumbnail_width,omitempty"`
	//Optional. Thumbnail height
	ThumbnailHeight *int `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultLocation) GetInlineQueryResultType() string {
	return "location"
}

func (i InlineQueryResultLocation) Validate() error {
	if i.Type != "location" {
		return ErrInvalidParam("type must be \"location\"")
	}
	b := len([]byte(i.Id))
	if b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if i.Latitude == nil {
		return ErrInvalidParam("latitude parameter can't be empty")
	}
	if i.Longtitude == nil {
		return ErrInvalidParam("longtitude parameter can't be empty")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
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

// Represents a venue. By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	//Type of the result, must be venue
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	//Latitude of the venue location in degrees
	Latitude *float32 `json:"latitude"`
	//Longitude of the venue location in degrees
	Longtitude *float32 `json:"longtitude"`
	//Title of the venue
	Title string `json:"title"`
	//Address of the venue
	Address string `json:"address"`
	//Optional. Foursquare identifier of the venue if known
	FoursquareId *string `json:"foursquare_id,omitempty"`
	//Optional. Foursquare type of the venue, if known.
	//(For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FourSquareType *string `json:"four_square_type,omitempty"`
	//Optional. Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	//Optional. Google Places type of the venue.
	//(See https://developers.google.com/places/web-service/supported_types)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the venue
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	//Optional. Url of the thumbnail for the result
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	//Optional. Thumbnail width
	ThumbnailWidth *int `json:"thumbnail_width,omitempty"`
	//Optional. Thumbnail height
	ThumbnailHeight *int `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultVenue) GetInlineQueryResultType() string {
	return "venue"
}

func (i InlineQueryResultVenue) Validate() error {
	if i.Type != "venue" {
		return ErrInvalidParam("type must be \"venue\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
	}
	if strings.TrimSpace(i.Address) == "" {
		return ErrInvalidParam("address parameter can't be empty")
	}
	if i.Latitude == nil {
		return ErrInvalidParam("latitude parameter can't be empty")
	}
	if i.Longtitude == nil {
		return ErrInvalidParam("longtitude parameter can't be empty")
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

// Represents a contact with a phone number. By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	//Type of the result, must be contact
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	//Contact's phone number
	PhoneNumber string `json:"phone_number"`
	//Contact's first name
	FirstName string `json:"first_name"`
	//Optional. Contact's last name
	LastName *string `json:"last_name,omitempty"`
	//Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard *string `json:"v_card,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the contact
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	//Optional. Url of the thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	//Optional. Thumbnail width
	ThumbnailWidth *int `json:"thumbnail_width,omitempty"`
	//Optional. Thumbnail height
	ThumbnailHeight *int `json:"thumbnail_height,omitempty"`
}

func (i InlineQueryResultContact) GetInlineQueryResultType() string {
	return "contact"
}

func (i InlineQueryResultContact) Validate() error {
	if i.Type != "contact" {
		return ErrInvalidParam("type must be \"contact\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.Id) == "" {
		return ErrInvalidParam("id parameter can't be empty")
	}
	if strings.TrimSpace(i.PhoneNumber) == "" {
		return ErrInvalidParam("phone_number parameter can't be empty")
	}
	if strings.TrimSpace(i.FirstName) == "" {
		return ErrInvalidParam("first_name parameter can't be empty")
	}
	if i.InputMessageContent != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
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

// Represents a Game.
type InlineQueryResultGame struct {
	//Type of the result, must be game
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//Short name of the game
	GameShortName string `json:"game_short_name"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (i InlineQueryResultGame) GetInlineQueryResultType() string {
	return "game"
}

func (i InlineQueryResultGame) Validate() error {
	if i.Type != "game" {
		return ErrInvalidParam("type must be \"game\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.GameShortName) == "" {
		return ErrInvalidParam("game_short_name parameter can't be empty")
	}
	if i.ReplyMarkup != nil {
		if err := (*i.ReplyMarkup).Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	//Type of the result, must be photo
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier of the photo
	PhotoFileId string `json:"photo_file_id"`
	//Optional. Title for the result
	Title *string `json:"title,omitempty"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	//Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the photo caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// /Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the photo
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedPhoto) GetInlineQueryResultType() string {
	return "photo"
}

func (i InlineQueryResultCachedPhoto) Validate() error {
	if i.Type != "photo" {
		return ErrInvalidParam("type must be \"photo\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.PhotoFileId) == "" {
		return ErrInvalidParam("photo_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	//Type of the result, must be gif
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier for the GIF file
	GifFileId string `json:"gif_file_id"`
	//Optional. Title for the result
	Title *string `json:"title,omitempty"`
	//Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedGif) GetInlineQueryResultType() string {
	return "gif"
}

func (i InlineQueryResultCachedGif) Validate() error {
	if i.Type != "gif" {
		return ErrInvalidParam("type must be \"gif\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.GifFileId) == "" {
		return ErrInvalidParam("gif_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be enabled if caption_entities are provided")
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

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	//Type of the result, must be mpeg4_gif
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier for the MPEG4 file
	Mpeg4FileId string `json:"mpeg_4_file_id"`
	//Optional. Title for the result
	Title *string `json:"title,omitempty"`
	//Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the video animation
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedMpeg4Gif) GetInlineQueryResultType() string {
	return "mpeg4_gif"
}

func (i InlineQueryResultCachedMpeg4Gif) Validate() error {
	if i.Type != "mpeg4_gif" {
		return ErrInvalidParam("type must be \"mpeg4_gif\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if i.Mpeg4FileId == "" {
		return ErrInvalidParam("mpeg_4_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	//Type of the result, must be sticker
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier of the sticker
	StickerFileId string `json:"sticker_file_id"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the sticker
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedSticker) GetInlineQueryResultType() string {
	return "sticker"
}

func (i InlineQueryResultCachedSticker) Validate() error {
	if i.Type != "sticker" {
		return ErrInvalidParam("type must be \"sticker\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if i.StickerFileId == "" {
		return ErrInvalidParam("sticker_file_id parameter can't be empty")
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

// Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	//Type of the result, must be document
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//Title for the result
	Title string `json:"title"`
	//A valid file identifier for the file
	DocumentFileId string `json:"document_file_id"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	//Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the document caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedDocument) GetInlineQueryResultType() string {
	return "document"
}

func (i InlineQueryResultCachedDocument) Validate() error {
	if i.Type != "document" {
		return ErrInvalidParam("type must be \"document\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if i.DocumentFileId == "" {
		return ErrInvalidParam("document_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	//Type of the result, must be video
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier for the video file
	VideoFileId string `json:"video_file_id"`
	//Title for the result
	Title string `json:"title"`
	//Optional. Short description of the result
	Description *string `json:"description,omitempty"`
	// /Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the video caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the video
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedVideo) GetInlineQueryResultType() string {
	return "video"
}

func (i InlineQueryResultCachedVideo) Validate() error {
	if i.Type != "video" {
		return ErrInvalidParam("type must be \"video\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.VideoFileId) == "" {
		return ErrInvalidParam("video_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	//Type of the result, must be voice
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier for the voice message
	VoiceFileId string `json:"voice_file_id"`
	//Voice message title
	Title string `json:"title"`
	//Optional. Caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the voice message caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the voice message
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedVoice) GetInlineQueryResultType() string {
	return "voice"
}

func (i InlineQueryResultCachedVoice) Validate() error {
	if i.Type != "voice" {
		return ErrInvalidParam("type must be \"voice\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter can't be empty")
	}
	if strings.TrimSpace(i.VoiceFileId) == "" {
		return ErrInvalidParam("voice_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse mode can't be enabled if Entities are provided")
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

// Represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	//Type of the result, must be audio
	Type string `json:"type"`
	//Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	//A valid file identifier for the audio file
	AudioFileId string `json:"audio_file_id"`
	//Optional. Caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the audio caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	//Optional. Content of the message to be sent instead of the audio
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedAudio) GetInlineQueryResultType() string {
	return "audio"
}

func (i InlineQueryResultCachedAudio) Validate() error {
	if i.Type != "audio" {
		return ErrInvalidParam("type must be \"audio\"")
	}
	if b := len([]byte(i.Id)); b < 1 || b > 64 {
		return ErrInvalidParam("id parameter must be between 1 and 64 bytes")
	}
	if strings.TrimSpace(i.AudioFileId) == "" {
		return ErrInvalidParam("audio_file_id parameter can't be empty")
	}
	if i.ParseMode != nil && i.CaptionEntities != nil {
		return fmt.Errorf("parse_mode can't be used if caption_entities are provided")
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

// This object represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 5 types:
//
// - InputTextMessageContent
//
// - InputLocationMessageContent
//
// - InputVenueMessageContent
//
// - InputContactMessageContent
//
// - InputInvoiceMessageContent
type InputMessageContent interface {
	Validable
	GetInputMessageContentType() string
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	//Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
	//Optional. Mode for parsing entities in the message text.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	Entities *[]MessageEntity `json:"entities,omitempty"`
	//Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (i InputTextMessageContent) GetInputMessageContentType() string {
	return "text"
}

func (i InputTextMessageContent) Validate() error {
	if l := len(i.MessageText); l < 1 || l > 4096 {
		return ErrInvalidParam("message_text parameter must be between 1 and 4096 characters")
	}
	if i.ParseMode != nil && i.Entities != nil {
		return fmt.Errorf("parse_mode can't be used if entities are provided")
	}
	if i.LinkPreviewOptions != nil {
		if err := i.LinkPreviewOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	//Latitude of the location in degrees
	Latitude *float64 `json:"latitude"`
	//Longitude of the location in degrees
	Longtitude *float64 `json:"longtitude"`
	//Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`
	//Optional. Period in seconds during which the location can be updated, should be between 60 and 86400,
	//or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int `json:"live_period,omitempty"`
	//Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int `json:"heading,omitempty"`
	//Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRaidius *int `json:"proximity_alert_raidius,omitempty"`
}

func (i InputLocationMessageContent) GetInputMessageContentType() string {
	return "location"
}

func (i InputLocationMessageContent) Validate() error {
	if i.Latitude == nil {
		return ErrInvalidParam("latitude parameter can't be empty")
	}
	if i.Longtitude == nil {
		return ErrInvalidParam("longtitude parameter can't be empty")
	}
	if i.LivePeriod != nil {
		if (*i.LivePeriod < 60 || *i.LivePeriod > 86400) && *i.LivePeriod != 0x7FFFFFFF {
			return ErrInvalidParam("live_period parameter must be between 60 and 86400 or equal to 0x7FFFFFFF")
		}
	}
	if i.HorizontalAccuracy != nil {
		if *i.HorizontalAccuracy < 0 || *i.HorizontalAccuracy > 1500 {
			return ErrInvalidParam("horizontal_accuracy parameter must be between 0 and 1500 meters")
		}
	}
	if i.Heading != nil {
		if *i.Heading < 1 || *i.Heading > 360 {
			return ErrInvalidParam("heading parameter must be between 0 and 1500")
		}
	}
	if i.ProximityAlertRaidius != nil {
		if *i.ProximityAlertRaidius < 1 || *i.ProximityAlertRaidius > 100000 {
			return ErrInvalidParam("proximity_alert_radius must be between 1 and 100000 meters")
		}
	}
	return nil
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	//Latitude of the venue in degrees
	Latitude *float64 `json:"latitude"`
	//Longitude of the venue in degrees
	Longtitude *float64 `json:"longtitude"`
	//Name of the venue
	Title string `json:"title"`
	//Address of the venue
	Address string `json:"address"`
	//Optional. Foursquare identifier of the venue, if known
	FoursquareId *string `json:"foursquare_id,omitempty"`
	//Optional. Foursquare type of the venue, if known.
	//(For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	//Optional. Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	//ptional. Google Places type of the venue.
	//(See https://developers.google.com/places/web-service/supported_types)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

func (i InputVenueMessageContent) GetInputMessageContentType() string {
	return "venue"
}

func (i InputVenueMessageContent) Validate() error {
	if i.Latitude == nil {
		return ErrInvalidParam("latitude parameter can't be empty")
	}
	if i.Longtitude == nil {
		return ErrInvalidParam("longtitude parameter can't be empty")
	}
	if strings.TrimSpace(i.Title) == "" {
		return fmt.Errorf("title parameter can't be empty")
	}
	if strings.TrimSpace(i.Address) == "" {
		return fmt.Errorf("address parameter can't be empty")
	}
	return nil
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	//Contact's phone number
	PhoneNumber string `json:"phone_number"`
	//Contact's first name
	FirstName string `json:"first_name"`
	//Optional. Contact's last name
	LastName string `json:"last_name"`
	//Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"v_card"`
}

func (i InputContactMessageContent) GetInputMessageContentType() string {
	return "contact"
}

func (i InputContactMessageContent) Validate() error {
	if i.PhoneNumber == "" || strings.TrimSpace(i.PhoneNumber) == "" {
		return fmt.Errorf("phone number can't be empty")
	}
	if i.FirstName == "" || strings.TrimSpace(i.FirstName) == "" {
		return fmt.Errorf("firstname can't be empty")
	}
	return nil
}

// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	//Product name, 1-32 characters
	Title string `json:"title"`
	//Product description, 1-255 characters
	Description string `json:"description"`
	//Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`
	//Optional. Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken *string `json:"provider_token,omitempty"`
	//Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`
	//Price breakdown, a JSON-serialized list of components
	//(e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.).
	//Must contain exactly one item for payments in Telegram Stars.
	Prices []LabeledPrice `json:"prices"`
	//Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double).
	//For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145.
	//See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	//it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	//Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount *int `json:"max_tip_amount,omitempty"`
	//Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double).
	//At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive,
	//passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts *[]int `json:"suggested_tip_amounts,omitempty"`
	//Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider.
	//A detailed description of the required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty"`
	//Optional. URL of the product photo for the invoice.
	//Can be a photo of the goods or a marketing image for a service.
	PhotoUrl *string `json:"photo_url,omitempty"`
	//Optional. Photo size in bytes
	PhotoSize *int `json:"photo_size,omitempty"`
	//Optional. Photo width
	PhotoWidth *int `json:"photo_width,omitempty"`
	//Optional. Photo height
	PhotoHeight *int `json:"photo_height,omitempty"`
	//Optional. Pass True if you require the user's full name to complete the order.
	//Ignored for payments in Telegram Stars.
	NeedName *bool `json:"need_name,omitempty"`
	//Optional. Pass True if you require the user's phone number to complete the order.
	//Ignored for payments in Telegram Stars.
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty"`
	//Optional. Pass True if you require the user's email address to complete the order.
	//Ignored for payments in Telegram Stars.
	NeedEmail *bool `json:"need_email,omitempty"`
	//Optional. Pass True if you require the user's shipping address to complete the order.
	//Ignored for payments in Telegram Stars.
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty"`
	//Optional. Pass True if the user's phone number should be sent to the provider.
	//Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty"`
	//Optional. Pass True if the user's email address should be sent to the provider.
	//Ignored for payments in Telegram Stars.
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty"`
	//Optional. Pass True if the final price depends on the shipping method.
	//Ignored for payments in Telegram Stars.
	IsFlexible *bool `json:"is_flexible,omitempty"`
}

func (i InputInvoiceMessageContent) GetInputMessageContentType() string {
	return "invoice"
}

func (i InputInvoiceMessageContent) Validate() error {
	if l := len(i.Title); l < 1 || l > 32 {
		return ErrInvalidParam("title parameter must be between 1 and 32 characters")
	}
	if l := len(i.Description); l < 1 || l > 255 {
		return ErrInvalidParam("description parameter must be between 1 and 255 characters")
	}
	if l := len([]byte(i.Payload)); l < 1 || l > 128 {
		return ErrInvalidParam("payload parameter must be between 1 and 128 bytes")
	}
	//FIXME: should properly validate currency codes as in
	//https://en.wikipedia.org/wiki/ISO_4217#Active_codes_(list_one)
	if len(i.Currency) > 3 {
		return ErrInvalidParam("currency parameter accepts only valid three-letter ISO 4217 currency codes")
	}
	if i.Currency == "XTR" {
		if len(i.Prices) > 1 {
			return ErrInvalidParam("prices parameter must contain exactly one item for payments in Telegram Stars.")
		}
		if i.MaxTipAmount != nil {
			return ErrInvalidParam("max_tip_amount parameter is not supported for payments in Telegram stars")
		}
	}
	if i.SuggestedTipAmounts != nil {
		if len(*i.SuggestedTipAmounts) > 4 {
			return ErrInvalidParam("suggested_tip_amounts parameter can't be longer than 4 elements")
		}
		for j := 0; j < len(*i.SuggestedTipAmounts); j++ {
			if (*i.SuggestedTipAmounts)[j] <= 0 {
				return ErrInvalidParam("suggested_tip_amounts parameter accepts only positive integers")
			}
			if (*i.SuggestedTipAmounts)[j] > *i.MaxTipAmount {
				return ErrInvalidParam("suggested_tip_amounts parameter accepts only integers that do not exceed the max_tip_amount")
			}
			if j > 0 && (*i.SuggestedTipAmounts)[j] <= (*i.SuggestedTipAmounts)[j-1] {
				return ErrInvalidParam("suggested_tip_amounts parameter should be passed in a strictly increased order")
			}
		}
	}
	return nil
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
//
// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
type ChosenInlineResult struct {
	//The unique identifier for the result that was chosen
	ResultId string `json:"result_id"`
	//The user that chose the result
	From User `json:"from"`
	//Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`
	//Optional. Identifier of the sent inline message.
	//Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	//The query that was used to obtain the result
	Query string `json:"query"`
}

// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	//Optional. Identifier of the sent inline message.
	//Available only if there is an inline keyboard attached to the message.
	InlineMessageId *string `json:"inline_message_id,omitempty"`
}

// Describes an inline message to be sent by a user of a Mini App.
type PreparedInlineMessage struct {
	//Unique identifier of the prepared message
	Id string
	//Expiration date of the prepared message, in Unix time. Expired prepared messages can no longer be used
	ExpirationDate int
}
