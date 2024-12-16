package methods

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type SendSticker[T int | string, B types.InputFile | string] struct {
	ChatId               T
	Sticker              B
	BusinessConnectionId *string
	MessageThreadId      *int
	Emoji                *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendSticker[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return types.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return types.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Sticker).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Sticker).(string); ok {
		if strings.TrimSpace(p) == "" {
			return types.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendSticker[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendSticker[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendSticker", s)
}

type GetStickerSet struct {
	Name string
}

func (g GetStickerSet) Validate() error {
	if strings.TrimSpace(g.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (g GetStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetStickerSet) Execute() (*types.StickerSet, error) {
	return internal.MakePostRequest[types.StickerSet](telego.GetToken(), "getStickerSet", g)
}

type GetCustomEmojiStickers struct {
	CustomEmojiIds []string
}

func (g GetCustomEmojiStickers) Validate() error {
	if len(g.CustomEmojiIds) == 0 {
		return types.ErrInvalidParam("custom_emoji_ids parameter can't be empty")
	}
	return nil
}

func (g GetCustomEmojiStickers) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetCustomEmojiStickers) Execute() (*[]types.Sticker, error) {
	return internal.MakePostRequest[[]types.Sticker](telego.GetToken(), "getCustomEmojiStickers", g)
}

type UploadStickerFile struct {
	UserId        int
	Sticker       types.InputFile
	StickerFormat string
}

var allowed_formats = []string{
	"static",
	"animated",
	"video",
}

func (u UploadStickerFile) Validate() error {
	if u.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if err := u.Sticker.Validate(); err != nil {
		return err
	}
	if !slices.Contains(allowed_formats, u.StickerFormat) {
		return types.ErrInvalidParam("sticker_format must be one of \"static\", \"animated\", \"video\"")
	}
	return nil
}

func (u UploadStickerFile) ToRequestBody() ([]byte, error) {
	return json.Marshal(u)
}

func (u UploadStickerFile) Execute() (*types.File, error) {
	return internal.MakePostRequest[types.File](telego.GetToken(), "uploadStickerFile", u)
}

type CreateNewStickerSet[T types.InputFile | string] struct {
	UserId          int
	Name            string
	Title           string
	Stickers        []types.InputSticker[T]
	StickerType     *string
	NeedsRepainting *bool
}

var valid_stickerset_name = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
var consecutive_underscores = regexp.MustCompile(`__+`)
var valid_stickertypes = []string{
	"regular",
	"mask",
	"custom_emoji",
}

func (c CreateNewStickerSet[T]) Validate() error {
	if c.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(c.Name) < 1 || len(c.Name) > 64 {
		return types.ErrInvalidParam("name parameter must be between 1 and 64 characters")
	}
	if !valid_stickerset_name.MatchString(c.Name) {
		return types.ErrInvalidParam("name parameter can contain only English letters, digits and underscores")
	}
	if consecutive_underscores.MatchString(c.Name) {
		return types.ErrInvalidParam("name parameter can't contain consecutive underscores")
	}
	if len(c.Title) < 1 || len(c.Title) > 64 {
		return types.ErrInvalidParam("title parameter must be between 1 and 64 characters")
	}
	for _, sticker := range c.Stickers {
		if err := sticker.Validate(); err != nil {
			return err
		}
	}
	if c.StickerType != nil {
		if !slices.Contains(valid_stickertypes, *c.StickerType) {
			return types.ErrInvalidParam("sticker_type must be \"regular\", \"mask\" or \"custom_emoji\"")
		}
	}
	return nil
}

func (c CreateNewStickerSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateNewStickerSet[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "createNewStickerSet", c)
}

type AddStickerToSet[T types.InputFile | string] struct {
	UserId  int
	Name    string
	Sticker types.InputSticker[T]
}

func (a AddStickerToSet[T]) Validate() error {
	if a.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(a.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	if err := a.Sticker.Validate(); err != nil {
		return err
	}
	return nil
}

func (a AddStickerToSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AddStickerToSet[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "addStickerToSet", a)
}

type SetStickerPositionInSet struct {
	Sticker  string
	Position int
}

func (s SetStickerPositionInSet) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return types.ErrInvalidParam("sticker parameter can't be empty")
	}
	if s.Position < 0 {
		return types.ErrInvalidParam("position parameter must be positive")
	}
	return nil
}

func (s SetStickerPositionInSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerPositionInSet) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setStickerPositionInSet", s)
}

type DeleteStickerFromSet struct {
	Sticker string
}

func (d DeleteStickerFromSet) Validate() error {
	if strings.TrimSpace(d.Sticker) == "" {
		return types.ErrInvalidParam("sticker parameter can't be empty")
	}
	return nil
}

func (d DeleteStickerFromSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteStickerFromSet) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteStickerFromSet", d)
}

type ReplaceStickerInSet[T types.InputFile | string] struct {
	UserId     int
	Name       string
	OldSticker string
	Sticker    types.InputSticker[T]
}

func (r ReplaceStickerInSet[T]) Validate() error {
	if r.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(r.OldSticker) == "" {
		return types.ErrInvalidParam("old_sticker parameter can't be empty")
	}
	if strings.TrimSpace(r.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	if err := r.Sticker.Validate(); err != nil {
		return err
	}
	return nil
}

func (r ReplaceStickerInSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(r)
}

func (r ReplaceStickerInSet[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "replaceStickerInSet", r)
}

type SetStickerEmojiList struct {
	Sticker   string
	EmojiList []string
}

func (s SetStickerEmojiList) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return types.ErrInvalidParam("sticker parameter can't be empty")
	}
	if len(s.EmojiList) < 1 || len(s.EmojiList) > 20 {
		return types.ErrInvalidParam("emoji_list parameter can contain only 1-20 elements")
	}
	return nil
}

func (s SetStickerEmojiList) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerEmojiList) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setStickerEmojiList", s)
}

type SetStickerKeywords struct {
	Sticker  string
	Keywords *[]string
}

func (s SetStickerKeywords) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return types.ErrInvalidParam("sticker parameter can't be empty")
	}
	if s.Keywords != nil {
		if len(*s.Keywords) > 20 {
			return types.ErrInvalidParam("keywords parameter can't be longer than 20")
		}
	}
	return nil
}

func (s SetStickerKeywords) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerKeywords) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setStickerKeywords", s)
}

type SetStickerMaskPosition struct {
	Sticker      string
	MaskPosition *types.MaskPosition
}

func (s SetStickerMaskPosition) Validate() error {
	if strings.TrimSpace(s.Sticker) == "" {
		return types.ErrInvalidParam("sticker parameter can't be empty")
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
	return internal.MakePostRequest[bool](telego.GetToken(), "setStickerMaskPosition", s)
}

type SetStickerSetTitle struct {
	Name  string
	Title string
}

func (s SetStickerSetTitle) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	if strings.TrimSpace(s.Title) == "" {
		return types.ErrInvalidParam("title parameter can't be empty")
	}
	return nil
}

func (s SetStickerSetTitle) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerSetTitle) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setStickerSetTitle", s)
}

type SetStickerSetThumbnail[T types.InputFile | string] struct {
	Name      string
	UserId    int
	Thumbnail *T
	Format    string
}

var valid_stickerset_thumbnail = []string{
	"static",
	"animated",
	"video",
}

func (s SetStickerSetThumbnail[T]) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	if s.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.Thumbnail != nil {
		if t, ok := any(*s.Thumbnail).(types.InputFile); ok {
			if err := t.Validate(); err != nil {
				return err
			}
		}
		if t, ok := any(*s.Thumbnail).(string); ok {
			if strings.TrimSpace(t) == "" {
				return types.ErrInvalidParam("thumbnail file id can't be empty")
			}
		}
	}
	if !slices.Contains(valid_stickerset_thumbnail, s.Format) {
		return types.ErrInvalidParam("format parameter must be one of “static”, “animated” or “video”")
	}
	return nil
}

func (s SetStickerSetThumbnail[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetStickerSetThumbnail[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setStickerSetThumbnail", s)
}

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string
	CustomEmojiId *string
}

func (s SetCustomEmojiStickerSetThumbnail) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	if s.CustomEmojiId != nil {
		if strings.TrimSpace(*s.CustomEmojiId) == "" {
			return types.ErrInvalidParam("custom_emoji_id parameter can't be empty")
		}
	}
	return nil
}

func (s SetCustomEmojiStickerSetThumbnail) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetCustomEmojiStickerSetThumbnail) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setCustomEmojiStickerSetThumbnail", s)
}

type DeleteStickerSet struct {
	Name string
}

func (d DeleteStickerSet) Validate() error {
	if strings.TrimSpace(d.Name) == "" {
		return types.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (d DeleteStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteStickerSet) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteStickerSet", d)
}

type GetAvailableGifts struct {
}

func (g GetAvailableGifts) Validate() error {
	return nil
}

func (g GetAvailableGifts) ToRequestBody() ([]byte, error) {
	return json.Marshal(struct{}{})
}

func (g GetAvailableGifts) Execute() (*types.Gifts, error) {
	return internal.MakeGetRequest[types.Gifts](telego.GetToken(), "getAvailableGifts", g)
}

type SendGift struct {
	UserId        int
	GiftId        string
	Text          *string
	TextParseMode *string
	TextEntities  *[]types.MessageEntity
}

func (s SendGift) Validate() error {
	if s.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(s.GiftId) == "" {
		return types.ErrInvalidParam("gift_id parameter can't be empty")
	}
	if s.Text != nil {
		if len(*s.Text) > 255 {
			return types.ErrInvalidParam("text parameter must not be longer than 255 characters")
		}
	}
	if s.TextParseMode != nil && s.TextEntities != nil {
		return types.ErrInvalidParam("parse_mode can't be used if entities are provided")
	}
	return nil
}

func (s SendGift) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendGift) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "sendGift", s)
}
