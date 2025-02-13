package methods

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/bigelle/gotely/api/objects"
)

type SendSticker struct {
	ChatId               string                   `json:"chat_id"`
	Sticker              objects.InputFile        `json:"sticker"`
	BusinessConnectionId *string                  `json:"business_connection_id,omitempty"`
	MessageThreadId      *int                     `json:"message_thread_id,omitempty"`
	Emoji                *string                  `json:"emoji,omitempty"`
	DisableNotification  *bool                    `json:"disable_notification,omitempty"`
	ProtectContent       *bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   *bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId      *string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          *objects.ReplyMarkup     `json:"reply_markup,omitempty"`
}

func (s SendSticker) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if p, ok := any(s.Sticker).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Sticker).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

type GetStickerSet struct {
	Name string `json:"name"`
}

func (g GetStickerSet) Validate() error {
	if strings.TrimSpace(g.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

type GetCustomEmojiStickers struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"`
}

func (g GetCustomEmojiStickers) Validate() error {
	if len(g.CustomEmojiIds) == 0 {
		return objects.ErrInvalidParam("custom_emoji_ids parameter can't be empty")
	}
	return nil
}

type UploadStickerFile struct {
	UserId        int               `json:"user_id"`
	Sticker       objects.InputFile `json:"sticker"`
	StickerFormat string            `json:"sticker_format"`
}

var allowed_formats = []string{
	"static",
	"animated",
	"video",
}

func (u UploadStickerFile) Validate() error {
	if u.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if err := u.Sticker.Validate(); err != nil {
		return err
	}
	if !slices.Contains(allowed_formats, u.StickerFormat) {
		return objects.ErrInvalidParam("sticker_format must be one of \"static\", \"animated\", \"video\"")
	}
	return nil
}

type CreateNewStickerSet struct {
	UserId          int                    `json:"user_id"`
	Name            string                 `json:"name"`
	Title           string                 `json:"title"`
	Stickers        []objects.InputSticker `json:"stickers"`
	StickerType     *string                `json:"sticker_type,omitempty"`
	NeedsRepainting *bool                  `json:"needs_repainting,omitempty"`
}

var valid_stickerset_name = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
var consecutive_underscores = regexp.MustCompile(`__+`)
var valid_stickerobjects = []string{
	"regular",
	"mask",
	"custom_emoji",
}

func (c CreateNewStickerSet) Validate() error {
	if c.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(c.Name) < 1 || len(c.Name) > 64 {
		return objects.ErrInvalidParam("name parameter must be between 1 and 64 characters")
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

type AddStickerToSet struct {
	UserId  int                  `json:"user_id"`
	Name    string               `json:"name"`
	Sticker objects.InputSticker `json:"sticker"`
}

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

type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
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

type DeleteStickerFromSet struct {
	Sticker string `json:"sticker"`
}

func (d DeleteStickerFromSet) Validate() error {
	if strings.TrimSpace(d.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	return nil
}

type ReplaceStickerInSet struct {
	UserId     int                  `json:"user_id"`
	Name       string               `json:"name"`
	OldSticker string               `json:"old_sticker"`
	Sticker    objects.InputSticker `json:"sticker"`
}

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

type SetStickerEmojiList struct {
	Sticker   string   `json:"sticker"`
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

type SetStickerKeywords struct {
	Sticker  string    `json:"sticker"`
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

type SetStickerMaskPosition struct {
	Sticker      string                `json:"sticker"`
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

type SetStickerSetTitle struct {
	Name  string `json:"name"`
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

type SetStickerSetThumbnail struct {
	Name      string             `json:"name"`
	UserId    int                `json:"user_id"`
	Thumbnail *objects.InputFile `json:"thumbnail,omitempty"`
	Format    string             `json:"format"`
}

var valid_stickerset_thumbnail = []string{
	"static",
	"animated",
	"video",
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
	if !slices.Contains(valid_stickerset_thumbnail, s.Format) {
		return objects.ErrInvalidParam("format parameter must be one of “static”, “animated” or “video”")
	}
	return nil
}

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string  `json:"name"`
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

type DeleteStickerSet struct {
	Name string `json:"name"`
}

func (d DeleteStickerSet) Validate() error {
	if strings.TrimSpace(d.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

type GetAvailableGifts struct {
}

func (g GetAvailableGifts) Validate() error {
	return nil
}

type SendGift struct {
	UserId        int                      `json:"user_id"`
	GiftId        string                   `json:"gift_id"`
	Text          *string                  `json:"text"`
	TextParseMode *string                  `json:"text_parse_mode,omitempty"`
	TextEntities  *[]objects.MessageEntity `json:"text_entities,omitempty"`
}

func (s SendGift) Validate() error {
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
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
