package methods

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Sticker).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Sticker).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("photo parameter can't be empty")
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
		return errors.ErrInvalidParam("name parameter can't be empty")
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
		return errors.ErrInvalidParam("custom_emoji_ids parameter can't be empty")
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
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if err := u.Sticker.Validate(); err != nil {
		return err
	}
	if !slices.Contains(allowed_formats, u.StickerFormat) {
		return errors.ErrInvalidParam("sticker_format must be one of \"static\", \"animated\", \"video\"")
	}
	return nil
}

func (u UploadStickerFile) ToRequestBody() ([]byte, error) {
	return json.Marshal(u)
}

func (u UploadStickerFile) Execute() (*types.File, error) {
	return internal.MakePostRequest[types.File](telego.GetToken(), "uploadStickerFile", u)
}
