package objects

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"slices"

	"github.com/bigelle/gotely"
)

// This object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”.
	// The type of the sticker is independent from its format,
	// which is determined by the fields is_animated and is_video.
	Type string `json:"type"`
	// Sticker width
	Width int `json:"width"`
	// Sticker height
	Height int `json:"height"`
	// True, if the sticker is animated
	IsAnimated bool `json:"is_animated"`
	// True, if the sticker is a video sticker
	IsVideo bool `json:"is_video"`
	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// Optional. Emoji associated with the sticker
	Emoji *string `json:"emoji,omitempty"`
	// Optional. Name of the sticker set to which the sticker belongs
	SetName *string `json:"set_name,omitempty"`
	// Optional. For premium regular stickers, premium animation for the sticker
	PremiumAnimation *File `json:"premium_animation,omitempty"`
	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// Optional. For custom emoji stickers, unique identifier of the custom emoji
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
	// Optional. True, if the sticker must be repainted to a text color in messages,
	// the color of the Telegram Premium badge in emoji status,
	// white color on chat photos, or another appropriate color in other places
	NeedsRepainting *bool `json:"needs_repainting,omitempty"`
	// Optional. File size in bytes
	FileSize *int `json:"file_size,omitempty"`
}

// This object represents a sticker set.
type StickerSet struct {
	// Sticker set name
	Name string `json:"name"`
	// Sticker set title
	Title string `json:"title"`
	// Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type"`
	// List of all set stickers
	Stickers []Sticker `json:"stickers"`
	// Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift *float32 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift *float32 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale *float32 `json:"scale"`
}

func (m MaskPosition) Validate() error {
	var err gotely.ErrFailedValidation
	if m.Point == "" && m.XShift == nil && m.YShift == nil && m.Scale == nil {
		err = append(err, fmt.Errorf("all fields must be non-empty'"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// This object describes a sticker to be added to a sticker set.
type InputSticker struct {
	// The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers,
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// upload a new one using multipart/form-data,
	// or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// Animated and video stickers can't be uploaded via HTTP URL.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Sticker InputFile `json:"sticker"`
	//Format of the added sticker, must be one of “static” for a .WEBP or .PNG image,
	//“animated” for a .TGS animation, “video” for a WEBM video
	Format string `json:"format"`
	// List of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
	// Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters.
	// For “regular” and “custom_emoji” stickers only.
	Keywords *[]string `json:"keywords,omitempty"`
}

func (i InputSticker) Validate() error {
	var err gotely.ErrFailedValidation
	if len(i.EmojiList) < 1 || len(i.EmojiList) > 20 {
		err = append(err, fmt.Errorf("emojiList parameter must be between 1 and 20"))
	}
	if len(*i.Keywords) < 1 || len(*i.Keywords) > 20 {
		err = append(err, fmt.Errorf("keyword parameter must be between 1 and 20"))
	}
	if !slices.Contains([]string{"static", "animated", "video"}, i.Format) {
		err = append(err, fmt.Errorf("format must be 'static', 'animated' or 'video'"))
	}
	if i.MaskPosition != nil {
		if er := i.MaskPosition.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if er := i.Sticker.Validate(); er != nil {
		err = append(err, er)
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (i InputSticker) WriteTo(mw *multipart.Writer) error {
	if err := mw.WriteField("format", i.Format); err != nil {
		return err
	}
	eb, err := json.Marshal(i.EmojiList)
	if err != nil {
		return err
	}
	if err := mw.WriteField("emoji_list", string(eb)); err != nil {
		return err
	}
	if i.MaskPosition != nil {
		if err := gotely.WriteJSONToForm(mw, "mask_position", *i.MaskPosition); err != nil {
			return err
		}
	}

	if i.Keywords != nil {
		if err := gotely.WriteJSONToForm(mw, "keywords", *i.Keywords); err != nil {
			return err
		}
	}

	if err := i.Sticker.WriteTo(mw, "sticker"); err != nil {
		return err
	}
	return nil
}

// This object represents a gift that can be sent by the bot.
type Gift struct {
	// Unique identifier of the gift
	Id string `json:"id"`
	// The sticker that represents the gift
	Sticker Sticker `json:"sticker"`
	// The number of Telegram Stars that must be paid to send the sticker
	StarCount int `json:"star_count"`
	// Optional. The total number of the gifts of this type that can be sent; for limited gifts only
	TotalCount *int `json:"total_count,omitempty"`
	// Optional. The number of remaining gifts of this type that can be sent; for limited gifts only
	RemainingCount *int `json:"remaining_count,omitempty"`
}

// This object represent a list of gifts.
type Gifts struct {
	// The list of gifts
	Gifts []Gift
}
