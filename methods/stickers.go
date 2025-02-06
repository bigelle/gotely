package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/bigelle/gotely/objects"
)

type SendSticker struct {
	ChatId               string
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
	client               *http.Client
	baseUrl              string
}

func (s *SendSticker) WithClient(c *http.Client) *SendSticker {
	s.client = c
	return s
}

func (s SendSticker) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendSticker) WithApiBaseUrl(u string) *SendSticker {
	s.baseUrl = u
	return s
}

func (s SendSticker) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendSticker) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendSticker) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendSticker", s)
}

type GetStickerSet struct {
	Name    string
	client  *http.Client
	baseUrl string
}

func (s *GetStickerSet) WithClient(c *http.Client) *GetStickerSet {
	s.client = c
	return s
}

func (s GetStickerSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetStickerSet) WithApiBaseUrl(u string) *GetStickerSet {
	s.baseUrl = u
	return s
}

func (s GetStickerSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (g GetStickerSet) Execute(token string) (*objects.StickerSet, error) {
	return SendTelegramPostRequest[objects.StickerSet](token, "getStickerSet", g)
}

type GetCustomEmojiStickers struct {
	CustomEmojiIds []string
	client         *http.Client
	baseUrl        string
}

func (s *GetCustomEmojiStickers) WithClient(c *http.Client) *GetCustomEmojiStickers {
	s.client = c
	return s
}

func (s GetCustomEmojiStickers) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetCustomEmojiStickers) WithApiBaseUrl(u string) *GetCustomEmojiStickers {
	s.baseUrl = u
	return s
}

func (s GetCustomEmojiStickers) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (g GetCustomEmojiStickers) Execute(token string) (*[]objects.Sticker, error) {
	return SendTelegramPostRequest[[]objects.Sticker](token, "getCustomEmojiStickers", g)
}

type UploadStickerFile struct {
	UserId        int
	Sticker       objects.InputFile
	StickerFormat string
	client        *http.Client
	baseUrl       string
}

func (s *UploadStickerFile) WithClient(c *http.Client) *UploadStickerFile {
	s.client = c
	return s
}

func (s UploadStickerFile) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UploadStickerFile) WithApiBaseUrl(u string) *UploadStickerFile {
	s.baseUrl = u
	return s
}

func (s UploadStickerFile) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (u UploadStickerFile) Execute(token string) (*objects.File, error) {
	return SendTelegramPostRequest[objects.File](token, "uploadStickerFile", u)
}

type CreateNewStickerSet struct {
	UserId          int
	Name            string
	Title           string
	Stickers        []objects.InputSticker
	StickerType     *string
	NeedsRepainting *bool
	client          *http.Client
	baseUrl         string
}

func (s *CreateNewStickerSet) WithClient(c *http.Client) *CreateNewStickerSet {
	s.client = c
	return s
}

func (s CreateNewStickerSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CreateNewStickerSet) WithApiBaseUrl(u string) *CreateNewStickerSet {
	s.baseUrl = u
	return s
}

func (s CreateNewStickerSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (c CreateNewStickerSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "createNewStickerSet", c)
}

type AddStickerToSet struct {
	UserId  int
	Name    string
	Sticker objects.InputSticker
	client  *http.Client
	baseUrl string
}

func (s *AddStickerToSet) WithClient(c *http.Client) *AddStickerToSet {
	s.client = c
	return s
}

func (s AddStickerToSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *AddStickerToSet) WithApiBaseUrl(u string) *AddStickerToSet {
	s.baseUrl = u
	return s
}

func (s AddStickerToSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (a AddStickerToSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "addStickerToSet", a)
}

type SetStickerPositionInSet struct {
	Sticker  string
	Position int
	client   *http.Client
	baseUrl  string
}

func (s *SetStickerPositionInSet) WithClient(c *http.Client) *SetStickerPositionInSet {
	s.client = c
	return s
}

func (s SetStickerPositionInSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetStickerPositionInSet) WithApiBaseUrl(u string) *SetStickerPositionInSet {
	s.baseUrl = u
	return s
}

func (s SetStickerPositionInSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetStickerPositionInSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setStickerPositionInSet", s)
}

type DeleteStickerFromSet struct {
	Sticker string
	client  *http.Client
	baseUrl string
}

func (s *DeleteStickerFromSet) WithClient(c *http.Client) *DeleteStickerFromSet {
	s.client = c
	return s
}

func (s DeleteStickerFromSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteStickerFromSet) WithApiBaseUrl(u string) *DeleteStickerFromSet {
	s.baseUrl = u
	return s
}

func (s DeleteStickerFromSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (d DeleteStickerFromSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteStickerFromSet", d)
}

type ReplaceStickerInSet struct {
	UserId     int
	Name       string
	OldSticker string
	Sticker    objects.InputSticker
	client     *http.Client
	baseUrl    string
}

func (s *ReplaceStickerInSet) WithClient(c *http.Client) *ReplaceStickerInSet {
	s.client = c
	return s
}

func (s ReplaceStickerInSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ReplaceStickerInSet) WithApiBaseUrl(u string) *ReplaceStickerInSet {
	s.baseUrl = u
	return s
}

func (s ReplaceStickerInSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (r ReplaceStickerInSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "replaceStickerInSet", r)
}

type SetStickerEmojiList struct {
	Sticker   string
	EmojiList []string
	client    *http.Client
	baseUrl   string
}

func (s *SetStickerEmojiList) WithClient(c *http.Client) *SetStickerEmojiList {
	s.client = c
	return s
}

func (s SetStickerEmojiList) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetStickerEmojiList) WithApiBaseUrl(u string) *SetStickerEmojiList {
	s.baseUrl = u
	return s
}

func (s SetStickerEmojiList) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetStickerEmojiList) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setStickerEmojiList", s)
}

type SetStickerKeywords struct {
	Sticker  string
	Keywords *[]string
	client   *http.Client
	baseUrl  string
}

func (s *SetStickerKeywords) WithClient(c *http.Client) *SetStickerKeywords {
	s.client = c
	return s
}

func (s SetStickerKeywords) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetStickerKeywords) WithApiBaseUrl(u string) *SetStickerKeywords {
	s.baseUrl = u
	return s
}

func (s SetStickerKeywords) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetStickerKeywords) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setStickerKeywords", s)
}

type SetStickerMaskPosition struct {
	Sticker      string
	MaskPosition *objects.MaskPosition
	client       *http.Client
	baseUrl      string
}

func (s *SetStickerMaskPosition) WithClient(c *http.Client) *SetStickerMaskPosition {
	s.client = c
	return s
}

func (s SetStickerMaskPosition) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetStickerMaskPosition) WithApiBaseUrl(u string) *SetStickerMaskPosition {
	s.baseUrl = u
	return s
}

func (s SetStickerMaskPosition) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetStickerMaskPosition) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setStickerMaskPosition", s)
}

type SetStickerSetTitle struct {
	Name    string
	Title   string
	client  *http.Client
	baseUrl string
}

func (s *SetStickerSetTitle) WithClient(c *http.Client) *SetStickerSetTitle {
	s.client = c
	return s
}

func (s SetStickerSetTitle) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetStickerSetTitle) WithApiBaseUrl(u string) *SetStickerSetTitle {
	s.baseUrl = u
	return s
}

func (s SetStickerSetTitle) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetStickerSetTitle) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setStickerSetTitle", s)
}

type SetStickerSetThumbnail struct {
	Name      string
	UserId    int
	Thumbnail *objects.InputFile
	Format    string
	client    *http.Client
	baseUrl   string
}

func (s *SetStickerSetThumbnail) WithClient(c *http.Client) *SetStickerSetThumbnail {
	s.client = c
	return s
}

func (s SetStickerSetThumbnail) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetStickerSetThumbnail) WithApiBaseUrl(u string) *SetStickerSetThumbnail {
	s.baseUrl = u
	return s
}

func (s SetStickerSetThumbnail) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetStickerSetThumbnail) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setStickerSetThumbnail", s)
}

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string
	CustomEmojiId *string
	client        *http.Client
	baseUrl       string
}

func (s *SetCustomEmojiStickerSetThumbnail) WithClient(c *http.Client) *SetCustomEmojiStickerSetThumbnail {
	s.client = c
	return s
}

func (s SetCustomEmojiStickerSetThumbnail) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetCustomEmojiStickerSetThumbnail) WithApiBaseUrl(u string) *SetCustomEmojiStickerSetThumbnail {
	s.baseUrl = u
	return s
}

func (s SetCustomEmojiStickerSetThumbnail) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetCustomEmojiStickerSetThumbnail) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setCustomEmojiStickerSetThumbnail", s)
}

type DeleteStickerSet struct {
	Name    string
	client  *http.Client
	baseUrl string
}

func (s *DeleteStickerSet) WithClient(c *http.Client) *DeleteStickerSet {
	s.client = c
	return s
}

func (s DeleteStickerSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteStickerSet) WithApiBaseUrl(u string) *DeleteStickerSet {
	s.baseUrl = u
	return s
}

func (s DeleteStickerSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (d DeleteStickerSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteStickerSet", d)
}

type GetAvailableGifts struct {
	client  *http.Client
	baseUrl string
}

func (s *GetAvailableGifts) WithClient(c *http.Client) *GetAvailableGifts {
	s.client = c
	return s
}

func (s GetAvailableGifts) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetAvailableGifts) WithApiBaseUrl(u string) *GetAvailableGifts {
	s.baseUrl = u
	return s
}

func (s GetAvailableGifts) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (g GetAvailableGifts) Validate() error {
	return nil
}

func (g GetAvailableGifts) ToRequestBody() ([]byte, error) {
	return json.Marshal(struct{}{})
}

func (g GetAvailableGifts) Execute(token string) (*objects.Gifts, error) {
	return SendTelegramGetRequest[objects.Gifts](token, "getAvailableGifts", g)
}

type SendGift struct {
	UserId        int
	GiftId        string
	Text          *string
	TextParseMode *string
	TextEntities  *[]objects.MessageEntity
	client        *http.Client
	baseUrl       string
}

func (s *SendGift) WithClient(c *http.Client) *SendGift {
	s.client = c
	return s
}

func (s SendGift) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendGift) WithApiBaseUrl(u string) *SendGift {
	s.baseUrl = u
	return s
}

func (s SendGift) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendGift) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "sendGift", s)
}
