// TODO: replace marshal json with encoder
package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"regexp"
	"slices"
	"strings"

	"github.com/bigelle/gotely/api"
	"github.com/bigelle/gotely/api/objects"
)

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers.
// On success, the sent [objects.Message] is returned.
type SendSticker struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet,
	// or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data.
	// More information on https://core.telegram.org/bots/api#sending-files.
	// Video and animated stickers can't be sent via an HTTP URL.
	Sticker objects.InputFile `json:"sticker"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Emoji associated with the sticker; only for just uploaded stickers
	Emoji *string `json:"emoji,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`

	contentType string
}

func (s SendSticker) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Sticker.Validate(); err != nil {
		return fmt.Errorf("invalid photo parameter: %w", err)
	}
	return nil
}

func (s SendSticker) Endpoint() string {
	return "sendSticker"
}

func (s SendSticker) Reader() (io.Reader, error) {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", s.ChatId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := s.Sticker.WriteTo(mw, "sticker"); err != nil {
			pw.CloseWithError(err)
			return
		}

		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Emoji != nil {
			if err := mw.WriteField("emoji", *s.Emoji); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", *s.MessageEffectId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyMarkup != nil {
			if err := api.WriteJSONToForm(mw, "reply_markup", *s.ReplyMarkup); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := api.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr, nil
}

func (s SendSticker) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to get a sticker set.
// On success, a [objects.StickerSet] object is returned.
type GetStickerSet struct {
	// REQUIRED:
	// Name of the sticker set
	Name string `json:"name"`
}

func (g GetStickerSet) Validate() error {
	if strings.TrimSpace(g.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (s GetStickerSet) Endpoint() string {
	return "getStickerSet"
}

func (s GetStickerSet) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s GetStickerSet) ContentType() string {
	return "application/json"
}

// Use this method to get information about custom emoji stickers by their identifiers.
// Returns an Array of [objects.Sticker] objects.
type GetCustomEmojiStickers struct {
	// REQUIRED:
	// A JSON-serialized list of custom emoji identifiers.
	// At most 200 custom emoji identifiers can be specified.
	CustomEmojiIds []string `json:"custom_emoji_ids"`
}

func (g GetCustomEmojiStickers) Validate() error {
	if len(g.CustomEmojiIds) == 0 {
		return objects.ErrInvalidParam("custom_emoji_ids parameter can't be empty")
	}
	return nil
}

func (s GetCustomEmojiStickers) Endpoint() string {
	return "getCustomEmojiStickers"
}

func (s GetCustomEmojiStickers) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s GetCustomEmojiStickers) ContentType() string {
	return "application/json"
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet,
// addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times).
// Returns the uploaded [objects.File] on success.
type UploadStickerFile struct {
	// REQUIRED:
	// User identifier of sticker file owner
	UserId int `json:"user_id"`
	// REQUIRED:
	// A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format.
	// See https://core.telegram.org/stickers for technical requirements.
	// More information on https://core.telegram.org/bots/api#sending-files
	Sticker objects.InputFile `json:"sticker"`
	// REQUIRED:
	// Format of the sticker, must be one of “static”, “animated”, “video”
	StickerFormat string `json:"sticker_format"`
}

func (u UploadStickerFile) Validate() error {
	if u.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if err := u.Sticker.Validate(); err != nil {
		return err
	}
	allowed_formats := []string{
		"static",
		"animated",
		"video",
	}
	if !slices.Contains(allowed_formats, u.StickerFormat) {
		return objects.ErrInvalidParam("sticker_format must be one of \"static\", \"animated\", \"video\"")
	}
	return nil
}

func (s UploadStickerFile) Endpoint() string {
	return "uploadStickerFile"
}

func (s UploadStickerFile) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
} // TODO multipart

func (s UploadStickerFile) ContentType() string {
	return "application/json"
}

// Use this method to create a new sticker set owned by a user.
// The bot will be able to edit the sticker set thus created.
// Returns True on success.
type CreateNewStickerSet struct {
	// REQUIRED:
	// User identifier of created sticker set owner
	UserId int `json:"user_id"`
	//REQUIRED:
	//Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals).
	//Can contain only English letters, digits and underscores.
	//Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>".
	//<bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`
	// REQUIRED:
	// Sticker set title, 1-64 characters
	Title string `json:"title"`
	// REQUIRED:
	// A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	Stickers []objects.InputSticker `json:"stickers"`

	// Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”.
	// By default, a regular sticker set is created.
	StickerType *string `json:"sticker_type,omitempty"`
	// Pass True if stickers in the sticker set must be repainted to the color of text when used in messages,
	// the accent color if used as emoji status, white on chat photos,
	// or another appropriate color based on context; for custom emoji sticker sets only
	NeedsRepainting *bool `json:"needs_repainting,omitempty"`
	contentType     string
}

func (c CreateNewStickerSet) Validate() error {
	if c.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(c.Name) < 1 || len(c.Name) > 64 {
		return objects.ErrInvalidParam("name parameter must be between 1 and 64 characters")
	}
	valid_stickerset_name := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	consecutive_underscores := regexp.MustCompile(`__+`)
	valid_stickerobjects := []string{
		"regular",
		"mask",
		"custom_emoji",
	}
	if !valid_stickerset_name.MatchString(c.Name) {
		return objects.ErrInvalidParam("name parameter can contain only English letters, digits and underscores")
	}
	if consecutive_underscores.MatchString(c.Name) {
		return objects.ErrInvalidParam("name parameter can't contain consecutive underscores")
	}
	if len(c.Title) < 1 || len(c.Title) > 64 {
		return objects.ErrInvalidParam("title parameter must be between 1 and 64 characters")
	}
	for _, sticker := range c.Stickers {
		if err := sticker.Validate(); err != nil {
			return err
		}
	}
	if c.StickerType != nil {
		if !slices.Contains(valid_stickerobjects, *c.StickerType) {
			return objects.ErrInvalidParam("sticker_type must be \"regular\", \"mask\" or \"custom_emoji\"")
		}
	}
	return nil
}

func (s CreateNewStickerSet) Endpoint() string {
	return "createNewStickerSet"
}

func (s *CreateNewStickerSet) Reader() (io.Reader, error) {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("user_id", fmt.Sprint(s.UserId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := mw.WriteField("name", s.Name); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := mw.WriteField("title", s.Title); err != nil {
			pw.CloseWithError(err)
			return
		}
		for _, sticker := range s.Stickers {
			if err := sticker.WriteTo(mw); err != nil {
				pw.CloseWithError(err)
				return
			}
		}

		if s.StickerType != nil {
			if err := mw.WriteField("sticker_type", *s.StickerType); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.NeedsRepainting != nil {
			if err := mw.WriteField("needs_repainting", fmt.Sprint(*s.NeedsRepainting)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr, nil
}

func (s CreateNewStickerSet) ContentType() string {
	return "application/json"
}

// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers.
// Other sticker sets can have up to 120 stickers.
// Returns True on success.
type AddStickerToSet struct {
	// REQUIRED:
	// User identifier of sticker set owner
	UserId int `json:"user_id"`
	// REQUIRED:
	// Sticker set name
	Name string `json:"name"`
	// REQUIRED:
	// A JSON-serialized object with information about the added sticker.
	// If exactly the same sticker had already been added to the set, then the set isn't changed.
	Sticker objects.InputSticker `json:"sticker"`
} // TODO multipart

func (a AddStickerToSet) Validate() error {
	if a.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(a.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	if err := a.Sticker.Validate(); err != nil {
		return err
	}
	return nil
}

func (s AddStickerToSet) Endpoint() string {
	return "addStickerToSet"
}

func (s AddStickerToSet) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s AddStickerToSet) ContentType() string {
	return "application/json"
}

// Use this method to move a sticker in a set created by the bot to a specific position.
// Returns True on success.
type SetStickerPositionInSet struct {
	// REQUIRED:
	// File identifier of the sticker
	Sticker string `json:"sticker"`
	// REQUIRED:
	// New sticker position in the set, zero-based
	Position int `json:"position"`
}

func (s SetStickerPositionInSet) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	if s.Position < 0 {
		return objects.ErrInvalidParam("position parameter must be positive")
	}
	return nil
}

func (s SetStickerPositionInSet) Endpoint() string {
	return "setStickerPositionInSet"
}

func (s SetStickerPositionInSet) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetStickerPositionInSet) ContentType() string {
	return "application/json"
}

// Use this method to delete a sticker from a set created by the bot.
// Returns True on success.
type DeleteStickerFromSet struct {
	// REQUIRED:
	Sticker string `json:"sticker"`
}

func (d DeleteStickerFromSet) Validate() error {
	if strings.TrimSpace(d.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	return nil
}

func (s DeleteStickerFromSet) Endpoint() string {
	return "deleteStickerFromSet"
}

func (s DeleteStickerFromSet) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s DeleteStickerFromSet) ContentType() string {
	return "application/json"
}

// Use this method to replace an existing sticker in a sticker set with a new one.
// The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet.
// Returns True on success.
type ReplaceStickerInSet struct {
	// REQUIRED:
	// User identifier of the sticker set owner
	UserId int `json:"user_id"`
	// REQUIRED:
	// Sticker set name
	Name string `json:"name"`
	// REQUIRED:
	// File identifier of the replaced sticker
	OldSticker string `json:"old_sticker"`
	// REQUIRED:
	// A JSON-serialized object with information about the added sticker.
	// If exactly the same sticker had already been added to the set, then the set remains unchanged.
	Sticker objects.InputSticker `json:"sticker"`
} // TODO multipart

func (r ReplaceStickerInSet) Validate() error {
	if r.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(r.OldSticker) == "" {
		return objects.ErrInvalidParam("old_sticker parameter can't be empty")
	}
	if strings.TrimSpace(r.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	if err := r.Sticker.Validate(); err != nil {
		return err
	}
	return nil
}

func (s ReplaceStickerInSet) Endpoint() string {
	return "replaceStickerInSet"
}

func (s ReplaceStickerInSet) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s ReplaceStickerInSet) ContentType() string {
	return "application/json"
}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot.
// Returns True on success.
type SetStickerEmojiList struct {
	// REQUIRED:
	// File identifier of the sticker
	Sticker string `json:"sticker"`
	// REQUIRED:
	// A JSON-serialized list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
}

func (s SetStickerEmojiList) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	if len(s.EmojiList) < 1 || len(s.EmojiList) > 20 {
		return objects.ErrInvalidParam("emoji_list parameter can contain only 1-20 elements")
	}
	return nil
}

func (s SetStickerEmojiList) Endpoint() string {
	return "setStickerEmojiList"
}

func (s SetStickerEmojiList) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetStickerEmojiList) ContentType() string {
	return "application/json"
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerKeywords struct {
	// REQUIRED:
	// File identifier of the sticker
	Sticker string `json:"sticker"`

	// A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
	Keywords *[]string `json:"keywords,omitempty"`
}

func (s SetStickerKeywords) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	if s.Keywords != nil {
		if len(*s.Keywords) > 20 {
			return objects.ErrInvalidParam("keywords parameter can't be longer than 20")
		}
	}
	return nil
}

func (s SetStickerKeywords) Endpoint() string {
	return "setStickerKeywords"
}

func (s SetStickerKeywords) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetStickerKeywords) ContentType() string {
	return "application/json"
}

// Use this method to change the mask position of a mask sticker.
// The sticker must belong to a sticker set that was created by the bot.
// Returns True on success.
type SetStickerMaskPosition struct {
	// REQUIRED:
	// File identifier of the sticker
	Sticker string `json:"sticker"`

	// A JSON-serialized object with the position where the mask should be placed on faces.
	// Omit the parameter to remove the mask position.
	MaskPosition *objects.MaskPosition `json:"mask_position,omitempty"`
}

func (s SetStickerMaskPosition) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	if s.MaskPosition != nil {
		if err := s.MaskPosition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetStickerMaskPosition) Endpoint() string {
	return "setStickerMaskPosition"
}

func (s SetStickerMaskPosition) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetStickerMaskPosition) ContentType() string {
	return "application/json"
}

// Use this method to set the title of a created sticker set. Returns True on success.
type SetStickerSetTitle struct {
	// REQUIRED:
	// Sticker set name
	Name string `json:"name"`
	// REQUIRED:
	// Sticker set title, 1-64 characters
	Title string `json:"title"`
}

func (s SetStickerSetTitle) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	if strings.TrimSpace(s.Title) == "" {
		return objects.ErrInvalidParam("title parameter can't be empty")
	}
	return nil
}

func (s SetStickerSetTitle) Endpoint() string {
	return "setStickerSetTitle"
}

func (s SetStickerSetTitle) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetStickerSetTitle) ContentType() string {
	return "application/json"
}

// Use this method to set the thumbnail of a regular or mask sticker set.
// The format of the thumbnail file must match the format of the stickers in the set.
// Returns True on success.
type SetStickerSetThumbnail struct {
	// REQUIRED:
	// Sticker set name
	Name string `json:"name"`
	// REQUIRED:
	// User identifier of the sticker set owner
	UserId int `json:"user_id"`
	// REQUIRED:
	// Format of the thumbnail, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, or “video” for a .WEBM video
	Format string `json:"format"`

	// A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px,
	// or a .TGS animation with a thumbnail up to 32 kilobytes in size
	// (see https://core.telegram.org/stickers#animation-requirements for animated sticker technical requirements),
	// or a .WEBM video with the thumbnail up to 32 kilobytes in size;
	// see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements.
	// Pass a file_id as a String to send a file that already exists on the Telegram servers,
	// pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	// More information on Sending Files https://core.telegram.org/bots/api#sending-files.
	// Animated and video sticker set thumbnails can't be uploaded via HTTP URL.
	// If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Thumbnail *objects.InputFile `json:"thumbnail,omitempty"`
}

func (s SetStickerSetThumbnail) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.Thumbnail != nil {
		if t, ok := any(*s.Thumbnail).(objects.InputFile); ok {
			if err := t.Validate(); err != nil {
				return err
			}
		}
		if t, ok := any(*s.Thumbnail).(string); ok {
			if strings.TrimSpace(t) == "" {
				return objects.ErrInvalidParam("thumbnail file id can't be empty")
			}
		}
	}
	valid_stickerset_thumbnail := []string{
		"static",
		"animated",
		"video",
	}
	if !slices.Contains(valid_stickerset_thumbnail, s.Format) {
		return objects.ErrInvalidParam("format parameter must be one of “static”, “animated” or “video”")
	}
	return nil
}

func (s SetStickerSetThumbnail) Endpoint() string {
	return "setStickerSetThumbnail"
}

func (s SetStickerSetThumbnail) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetStickerSetThumbnail) ContentType() string {
	return "application/json"
}

// Use this method to set the thumbnail of a custom emoji sticker set.
// Returns True on success.
type SetCustomEmojiStickerSetThumbnail struct {
	// REQUIRED:
	// Sticker set name
	Name string `json:"name"`

	// Custom emoji identifier of a sticker from the sticker set;
	// pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
}

func (s SetCustomEmojiStickerSetThumbnail) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	if s.CustomEmojiId != nil {
		if strings.TrimSpace(*s.CustomEmojiId) == "" {
			return objects.ErrInvalidParam("custom_emoji_id parameter can't be empty")
		}
	}
	return nil
}

func (s SetCustomEmojiStickerSetThumbnail) Endpoint() string {
	return "setCustomEmojiStickerSetThumbnail"
}

func (s SetCustomEmojiStickerSetThumbnail) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SetCustomEmojiStickerSetThumbnail) ContentType() string {
	return "application/json"
}

// Use this method to delete a sticker set that was created by the bot.
// Returns True on success.
type DeleteStickerSet struct {
	// REQUIRED:
	// Sticker set name
	Name string `json:"name"`
}

func (d DeleteStickerSet) Validate() error {
	if strings.TrimSpace(d.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (s DeleteStickerSet) Endpoint() string {
	return "deleteStickerSet"
}

func (s DeleteStickerSet) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s DeleteStickerSet) ContentType() string {
	return "application/json"
}

// Returns the list of gifts that can be sent by the bot to users and channel chats.
// Requires no parameters.
// Returns a [objects.Gifts] object.
type GetAvailableGifts struct{}

func (g GetAvailableGifts) Validate() error {
	return nil
}

func (s GetAvailableGifts) Endpoint() string {
	return "getAvailableGifts"
}

func (s GetAvailableGifts) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s GetAvailableGifts) ContentType() string {
	return "application/json"
}

// Sends a gift to the given user or channel chat.
// The gift can't be converted to Telegram Stars by the receiver.
// Returns True on success.
type SendGift struct {
	// REQUIRED:
	// Identifier of the gift
	GiftId string `json:"gift_id"`
	// REQUIRED:
	// Text that will be shown along with the gift; 0-128 characters
	Text *string `json:"text"`

	// Required if chat_id is not specified. Unique identifier of the target user who will receive the gift.
	UserId *int `json:"user_id"`
	// Required if user_id is not specified.
	// Unique identifier for the chat or username of the channel (in the format @channelusername) that will receive the gift.
	ChatId *string `json:"chat_id,omitempty,"`
	// Pass True to pay for the gift upgrade from the bot's balance,
	// thereby making the upgrade free for the receiver
	PayForUpgrade *bool `json:"pay_for_upgrade,omitempty"`
	// Mode for parsing entities in the text. See formatting options for more details.
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty,"`
	// A JSON-serialized list of special entities that appear in the gift text.
	// It can be specified instead of text_parse_mode.
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextEntities *[]objects.MessageEntity `json:"text_entities,omitempty,"`
}

func (s SendGift) Validate() error {
	if s.UserId != nil {
		if *s.UserId < 1 {
			return objects.ErrInvalidParam("user_id parameter can't be empty")
		}
	}
	if s.ChatId != nil {
		if *s.ChatId == "" {
			return objects.ErrInvalidParam("user_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.GiftId) == "" {
		return objects.ErrInvalidParam("gift_id parameter can't be empty")
	}
	if s.Text != nil {
		if len(*s.Text) > 255 {
			return objects.ErrInvalidParam("text parameter must not be longer than 255 characters")
		}
	}
	if s.TextParseMode != nil && s.TextEntities != nil {
		return objects.ErrInvalidParam("parse_mode can't be used if entities are provided")
	}
	return nil
}

func (s SendGift) Endpoint() string {
	return "sendGift"
}

func (s SendGift) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SendGift) ContentType() string {
	return "application/json"
}

// Verifies a user on behalf of the organization which is represented by the bot.
// Returns True on success.
type VerifyUser struct {
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Custom description for the verification; 0-70 characters.
	// Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty"`
}

func (v VerifyUser) Validate() error {
	if v.UserId <= 0 {
		return objects.ErrInvalidParam("user_id can't be empty or negative")
	}
	if v.CustomDescription != nil {
		if len(*v.CustomDescription) > 70 {
			return objects.ErrInvalidParam("custom_description must be between 0 and 70 characters")
		}
	}
	return nil
}

func (s VerifyUser) Endpoint() string {
	return "verifyUser"
}

func (s VerifyUser) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s VerifyUser) ContentType() string {
	return "application/json"
}

// Verifies a chat on behalf of the organization which is represented by the bot.
// Returns True on success.
type VerifyChat struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`

	// Custom description for the verification; 0-70 characters.
	// Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty"`
}

func (v VerifyChat) Validate() error {
	if v.ChatId == "" {
		return objects.ErrInvalidParam("user_id can't be empty or negative")
	}
	if v.CustomDescription != nil {
		if len(*v.CustomDescription) > 70 {
			return objects.ErrInvalidParam("custom_description must be between 0 and 70 characters")
		}
	}
	return nil
}

func (s VerifyChat) Endpoint() string {
	return "verifyChat"
}

func (s VerifyChat) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s VerifyChat) ContentType() string {
	return "application/json"
}

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot.
// Returns True on success.
type RemoveUserVerification struct {
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (r RemoveUserVerification) Validate() error {
	if r.UserId <= 0 {
		return objects.ErrInvalidParam("user_id can't be empty or negative")
	}
	return nil
}

func (s RemoveUserVerification) Endpoint() string {
	return "removeUserVerification"
}

func (s RemoveUserVerification) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s RemoveUserVerification) ContentType() string {
	return "application/json"
}

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot.
// Returns True on success.
type RemoveChatVerification struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (r RemoveChatVerification) Validate() error {
	if r.ChatId == "" {
		return objects.ErrInvalidParam("user_id can't be empty or negative")
	}
	return nil
}

func (s RemoveChatVerification) Endpoint() string {
	return "removeChatVerification"
}

func (s RemoveChatVerification) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s RemoveChatVerification) ContentType() string {
	return "application/json"
}
