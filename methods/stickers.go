package methods

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/bigelle/gotely/objects"
)

type SendSticker[T int | string] struct {
	ChatId               T
	Sticker              objects.InputFile
	BusinessConnectionId *string
	MessageThreadId      *int
	Emoji                *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendSticker[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (s SendSticker[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendSticker[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendSticker", s)
}

type GetStickerSet struct {
	Name string
}

func (g GetStickerSet) Validate() error {
	if strings.TrimSpace(g.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (g GetStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetStickerSet) Execute() (*objects.StickerSet, error) {
	return MakePostRequest[objects.StickerSet]("getStickerSet", g)
}

type GetCustomEmojiStickers struct {
	CustomEmojiIds []string
}

func (g GetCustomEmojiStickers) Validate() error {
	if len(g.CustomEmojiIds) == 0 {
		return objects.ErrInvalidParam("custom_emoji_ids parameter can't be empty")
	}
	return nil
}

func (g GetCustomEmojiStickers) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetCustomEmojiStickers) Execute() (*[]objects.Sticker, error) {
	return MakePostRequest[[]objects.Sticker]("getCustomEmojiStickers", g)
}

type UploadStickerFile struct {
	UserId        int
	Sticker       objects.InputFile
	StickerFormat string
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

func (u UploadStickerFile) ToRequestBody() ([]byte, error) {
	return json.Marshal(u)
}

func (u UploadStickerFile) Execute() (*objects.File, error) {
	return MakePostRequest[objects.File]("uploadStickerFile", u)
}

type CreateNewStickerSet struct {
	UserId          int
	Name            string
	Title           string
	Stickers        []objects.InputSticker
	StickerType     *string
	NeedsRepainting *bool
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

func (c CreateNewStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateNewStickerSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("createNewStickerSet", c)
}

type AddStickerToSet struct {
	UserId  int
	Name    string
	Sticker objects.InputSticker
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

func (a AddStickerToSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AddStickerToSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("addStickerToSet", a)
}

type SetStickerPositionInSet struct {
	Sticker  string
	Position int
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

func (s SetStickerPositionInSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerPositionInSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("setStickerPositionInSet", s)
}

type DeleteStickerFromSet struct {
	Sticker string
}

func (d DeleteStickerFromSet) Validate() error {
	if strings.TrimSpace(d.Sticker) == "" {
		return objects.ErrInvalidParam("sticker parameter can't be empty")
	}
	return nil
}

func (d DeleteStickerFromSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteStickerFromSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteStickerFromSet", d)
}

type ReplaceStickerInSet struct {
	UserId     int
	Name       string
	OldSticker string
	Sticker    objects.InputSticker
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

func (r ReplaceStickerInSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(r)
}

func (r ReplaceStickerInSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("replaceStickerInSet", r)
}

type SetStickerEmojiList struct {
	Sticker   string
	EmojiList []string
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

func (s SetStickerEmojiList) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerEmojiList) Execute() (*bool, error) {
	return MakePostRequest[bool]("setStickerEmojiList", s)
}

type SetStickerKeywords struct {
	Sticker  string
	Keywords *[]string
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

func (s SetStickerKeywords) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerKeywords) Execute() (*bool, error) {
	return MakePostRequest[bool]("setStickerKeywords", s)
}

type SetStickerMaskPosition struct {
	Sticker      string
	MaskPosition *objects.MaskPosition
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

func (s SetStickerMaskPosition) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerMaskPosition) Execute() (*bool, error) {
	return MakePostRequest[bool]("setStickerMaskPosition", s)
}

type SetStickerSetTitle struct {
	Name  string
	Title string
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

func (s SetStickerSetTitle) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerSetTitle) Execute() (*bool, error) {
	return MakePostRequest[bool]("setStickerSetTitle", s)
}

type SetStickerSetThumbnail struct {
	Name      string
	UserId    int
	Thumbnail *objects.InputFile
	Format    string
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

func (s SetStickerSetThumbnail) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerSetThumbnail) Execute() (*bool, error) {
	return MakePostRequest[bool]("setStickerSetThumbnail", s)
}

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string
	CustomEmojiId *string
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

func (s SetCustomEmojiStickerSetThumbnail) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetCustomEmojiStickerSetThumbnail) Execute() (*bool, error) {
	return MakePostRequest[bool]("setCustomEmojiStickerSetThumbnail", s)
}

type DeleteStickerSet struct {
	Name string
}

func (d DeleteStickerSet) Validate() error {
	if strings.TrimSpace(d.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (d DeleteStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteStickerSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteStickerSet", d)
}

type GetAvailableGifts struct {
}

func (g GetAvailableGifts) Validate() error {
	return nil
}

func (g GetAvailableGifts) ToRequestBody() ([]byte, error) {
	return json.Marshal(struct{}{})
}

func (g GetAvailableGifts) Execute() (*objects.Gifts, error) {
	return MakeGetRequest[objects.Gifts]("getAvailableGifts", g)
}

type SendGift struct {
	UserId        int
	GiftId        string
	Text          *string
	TextParseMode *string
	TextEntities  *[]objects.MessageEntity
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

func (s SendGift) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendGift) Execute() (*bool, error) {
	return MakePostRequest[bool]("sendGift", s)
}
