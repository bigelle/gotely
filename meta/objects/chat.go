package objects

const (
	USERCHATTYPE       = "private"
	GROUPCHATTYPE      = "group"
	CHANNELCHATTYPE    = "channel"
	SUPERGROUPCHATTYPE = "supergroup"
)

type Chat struct {
	id                                 int64
	Type                               string
	title                              string
	firstName                          string
	lastName                           string
	userName                           string
	photo                              *ChatPhoto
	desciption                         string
	inviteLink                         string
	pinnedMessage                      Message
	stickerSetName                     string
	canSetStickerSet                   bool
	permissions                        *ChatPermissions
	slowModeDelay                      int
	bio                                string
	linkedChatId                       int64
	location                           *ChatLocation
	messageAutoDeleteTime              int
	hasPrivateForwards                 bool
	hasProtectedContent                bool
	joinToSendMessages                 bool
	joinByRequest                      bool
	hasRestrictedVoiceAndVideoMessages bool
	isForum                            bool
	activeUserNames                    []string
	emojiStatusCustomEmojiId           string
	hasAggressiveAntiSpamEnabled       bool
	hasHiddenMembers                   bool
	emojiStatusExpirationDate          bool
	availableReactions                 *[]ReactionType
	accentColorId                      int
	backgroundCustomEmojiId            string
	profileAccentColorId               bool
	profileBackgroundCustomEmojiId     string
	hasVisibleHistory                  bool
	unrestrictBoostCount               int
	customEmojiStickerSetName          string
	birthDate                          *BirthDate
	businessIntro                      *BusinessIntro
	businessLocation                   *BusinessLocation
	BusinessOpeningHours               *BusinessOpeningHours
	personalChat                       *Chat
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
