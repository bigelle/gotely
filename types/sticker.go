package types

import (
	"fmt"
	"slices"
)

type Sticker struct {
	FileId           string        `json:"file_id"`
	FileUniqueId     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	FileSize         *int          `json:"file_size,omitempty"`
	Emoji            *string       `json:"emoji,omitempty"`
	SetName          *string       `json:"set_name,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	CustomEmojiId    *string       `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`
}

type StickerSet struct {
	StickerType string     `json:"sticker_type"`
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	Stickers    []Sticker  `json:"stickers"`
	IsAnimated  bool       `json:"is_animated"`
	IsVideo     bool       `json:"is_video"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
}

type InputSticker struct {
	Sticker      InputFile     `json:"sticker"`
	EmojiList    []string      `json:"emoji_list"`
	Format       string        `json:"format"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     *[]string     `json:"keywords,omitempty"`
}

func (i InputSticker) Validate() error {
	if len(i.EmojiList) < 1 || len(i.EmojiList) > 20 {
		return fmt.Errorf("EmojiList parameter must be between 1 and 20")
	}
	if len(*i.Keywords) < 1 || len(*i.Keywords) > 20 {
		return fmt.Errorf("Keyword parameter must be between 1 and 20")
	}
	if !slices.Contains([]string{"static", "animated", "video"}, i.Format) {
		return fmt.Errorf("Format must be 'static', 'animated' or 'video'")
	}
	if i.MaskPosition != nil {
		if err := i.MaskPosition.Validate(); err != nil {
			return err
		}
	}
	if err := i.Sticker.Validate(); err != nil {
		return err
	}
	return nil
}

type MaskPosition struct {
	Point  string   `json:"point"`
	XShift *float32 `json:"x_shift"`
	YShift *float32 `json:"y_shift"`
	Scale  *float32 `json:"scale"`
}

func (m MaskPosition) Validate() error {
	if m.Point == "" && m.XShift == nil && m.YShift == nil && m.Scale == nil {
		return fmt.Errorf("all fields must be non-empty'")
	}
	return nil
}
