package objects

import (
	"encoding/json"

	"github.com/bigelle/tele.go/api/meta/interfaces"
)

const (
	USERCHATTYPE       = "private"
	GROUPCHATTYPE      = "group"
	CHANNELCHATTYPE    = "channel"
	SUPERGROUPCHATTYPE = "supergroup"
)

type Chat struct {
	Id                                 int64
	Type                               string
	Title                              string
	FirstName                          string
	LastName                           string
	UserName                           string
	Photo                              *ChatPhoto
	Desciption                         string
	InviteLink                         string
	PinnedMessage                      Message
	StickerSetName                     string
	CanSetStickerSet                   bool
	Permissions                        *ChatPermissions
	SlowModeDelay                      int
	Bio                                string
	LinkedChatId                       int64
	Location                           *ChatLocation
	MessageAutoDeleteTime              int
	HasPrivateForwards                 bool
	HasProtectedContent                bool
	JoinToSendMessages                 bool
	JoinByRequest                      bool
	HasRestrictedVoiceAndVideoMessages bool
	IsForum                            bool
	ActiveUserNames                    []string
	EmojiStatusCustomEmojiId           string
	HasAggressiveAntiSpamEnabled       bool
	HasHiddenMembers                   bool
	EmojiStatusExpirationDate          bool
	AvailableReactions                 *[]ReactionType
	AccentColorId                      int
	BackgroundCustomEmojiId            string
	ProfileAccentColorId               bool
	ProfileBackgroundCustomEmojiId     string
	HasVisibleHistory                  bool
	UnrestrictBoostCount               int
	CustomEmojiStickerSetName          string
	BirthDate                          *BirthDate
	BusinessIntro                      *BusinessIntro
	BusinessLocation                   *BusinessLocation
	BusinessOpeningHours               *BusinessOpeningHours
	PersonalChat                       *Chat
}

func (c Chat) IsGroupChat() bool {
	return c.Type == GROUPCHATTYPE
}

func (c Chat) IsChannelChat() bool {
	return c.Type == CHANNELCHATTYPE
}

func (c Chat) IsUserChat() bool {
	return c.Type == USERCHATTYPE
}

func (c Chat) IsSuperGroupChat() bool {
	return c.Type == SUPERGROUPCHATTYPE
}

type ChatFullInfo struct {
	Photo                              ChatPhoto
	Description                        string
	InviteLink                         string
	PinnedMessage                      Message
	StickerSetName                     string
	CanSetStickerSet                   bool
	Permissions                        ChatPermissions
	SlowModeDelay                      int
	Bio                                string
	LinkedChatId                       int64
	Location                           ChatLocation
	MessageAutoDeleteTime              int
	HAsPrivateForwards                 bool
	HasProtectedCount                  bool
	JoinToSendMessage                  bool
	JoinByRequest                      bool
	HasRestrictedVoiceAndVideoMessages bool
	ActiveUsernames                    []string
	EmojiStatusCustomEmojiId           string
	HasAggressiveAntiSpamEnabled       string
	HasHiddenMembers                   bool
	EmojiStatusExpirationDate          bool
	AvailableReactions                 []ReactionType
	AccentColorId                      int
	BackgroundCustomEmojiId            string
	ProfileAccentColorId               bool
	ProfileBackgroundCustomEmojiId     string
	HasVisibleHistory                  bool
	UnrestrictBoostCount               int
	CustomEmojiStickerSetName          string
	BirthDate                          BirthDate
	BusinessIntro                      BusinessIntro
	BusinessLocation                   BusinessLocation
	BusinessOpeningHours               BusinessOpeningHours
	PersonalChat                       Chat
	CanSendPaidMedia                   bool
}

type ChatBackground struct {
	Type BackgroundType
}

type BackgroundType interface {
	interfaces.BotApiObject
}

//TODO: all background related stuff

type BackgroundTypeChatTheme struct {
	Type      string
	ThemeName string
}

func (b BackgroundTypeChatTheme) MarshalJSON() ([]byte, error) {
	type Alias BackgroundTypeChatTheme
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "chat_theme",
		Alias: (*Alias)(&b),
	})
}

type BackgroundTypeFill struct {
	Type             string
	Fill             BackgroundFill
	DarkThemeDimming int
}

func (b BackgroundTypeFill) MarshalJSON() ([]byte, error) {
	type Alias BackgroundTypeFill
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "fill",
		Alias: (*Alias)(&b),
	})
}

type BackgroundTypePattern struct {
	Type       string
	Document   Document
	Fill       BackgroundFill
	Intensity  int
	IsInverted bool
	IsMoving   bool
}

func (b BackgroundTypePattern) MarshalJSON() ([]byte, error) {
	type Alias BackgroundTypePattern
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "pattern",
		Alias: (*Alias)(&b),
	})
}

type BackgroundTypeWallpaper struct {
	Type             string
	Document         Document
	DarkThemeDimming int
	IsBlurred        bool
	IsMoving         bool
}

func (b BackgroundTypeWallpaper) MarshalJSON() ([]byte, error) {
	type Alias BackgroundTypeWallpaper
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "wallpaper",
		Alias: (*Alias)(&b),
	})
}

type BackgroundFill interface {
	//TODO: some contract
}

//TODO: base type and its connection

type BackgroundFillFreeformGradient struct {
	Type   string
	Colors []int
}

func (b BackgroundFillFreeformGradient) MarshalJSON() ([]byte, error) {
	type Alias BackgroundFillFreeformGradient
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "freeform_gradient",
		Alias: (*Alias)(&b),
	})
}

type BackgroundFillGradient struct {
	Type          string
	TopColor      int
	BottomColor   int
	RotationAngle int
}

func (b BackgroundFillGradient) MarshalJSON() ([]byte, error) {
	type Alias BackgroundFillGradient
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "gradient",
		Alias: (*Alias)(&b),
	})
}

type BackgroundFillSolid struct {
	Type  string
	Color int
}

func (b BackgroundFillSolid) MarshalJSON() ([]byte, error) {
	type Alias BackgroundFillSolid
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "solid",
		Alias: (*Alias)(&b),
	})
}
