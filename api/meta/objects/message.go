package objects

import (
	"reflect"
	"strings"
)

type Message struct {
	MessageId                     int
	MessageThreadId               int
	From                          *User
	Date                          int
	Chat                          *Chat
	ForwardFrom                   *User
	ForwardFromChat               *Chat
	ForwardDate                   int
	Text                          string
	Entities                      *[]MessageEntity
	CaptionEntities               *[]MessageEntity
	Audio                         *Audio
	Document                      *Document
	Photo                         *[]PhotoSize
	Sticker                       *Sticker
	video                         *Video
	Contact                       *Contact
	Location                      *Location
	Venue                         *Venue
	Animation                     *Animation
	PinnedMessage                 *MaybeInaccessibleMessage
	NewChatMembers                *[]User
	LeftChatMember                *User
	NewChatTitle                  string
	NewChatPhoto                  *[]PhotoSize
	DeleteChatPhoto               bool
	GroupChatCreated              bool
	ReplyToMessage                *Message
	Voice                         *Voice
	Caption                       string
	SuperGroupCreated             bool
	MigrateToChatId               int64
	MigrateFromChatId             int64
	EditDate                      int
	Game                          *Game
	ForwardFromMessageId          int
	Invoice                       *Invoice
	SuccessfulPayment             *SuccessfulPayment
	VideoNote                     *VideoNote
	AuthorSignature               string
	ForwardSignature              string
	MediaGroupId                  string
	ConnectedWebsite              string
	PassportData                  *PassportData
	ForwardSenderName             string
	Poll                          *Poll
	ReplyMarkup                   *InlineKeyboardMarkup
	Dice                          *Dice
	ViaBot                        *User
	SenderChat                    *Chat
	ProximityAlertTriggered       *ProximityAlertTriggered
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged
	IsAutomaticForward            bool
	HasProtectedContent           bool
	WebAppData                    *WebAppData
	VideoChatStarted              *VideoChatStarted
	VideoChatEnded                *VideoChatEnded
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited
	VideoChatScheduled            *VideoChatScheduled
	IsTopicMessage                bool
	ForumTopicCreated             *ForumTopicCreated
	ForumTopicClosed              *ForumTopicClosed
	ForumTopicReopened            *ForumTopicReopened
	ForumTopicEdited              *ForumTopicEdited
	GeneralForumTopicHidden       *GeneralForumTopicHidden
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden
	WriteAccessAllowed            *WriteAccessAllowed
	HasMediaSpoiler               bool
	UserShared                    *UserShared
	ChatShared                    *ChatShared
	Story                         *Story
	ExternalReplyInfo             *ExternalReplyInfo
	ForwardOrigin                 *MessageOrigin
	LinkPreviewOptions            *LinkPreviewOptions
	Quote                         *TextQuote
	UsersShared                   *UsersShared
	GiveawayCreated               *GiveawayCreated
	Giveaway                      *Giveaway
	GiveawayWinners               *GiveawayWinners
	GiveawayCompleted             *GiveawayCompleted
	ReplyToStory                  *Story
	BoostAdded                    *ChatBoostAdded
	SenderBoostCount              int
	BusinessConnectionId          string
	SenderBusinessBot             *User
	IsFromOffline                 bool
	ChatBackgroundSet             *ChatBackground
	EffectId                      string
	ShowCaptionAboveMedia         bool
	PaidMedia                     *PaidMediaInfo
	RefundedPayment               *RefundedPayment
}

func (m *Message) GetEntities() []MessageEntity {
	if m.Entities != nil {
		for _, e := range *m.Entities {
			e.ComputeText(m.Text)
		}
	}
	return *m.Entities
}

func (m *Message) GetCaptionEntities() []MessageEntity {
	if m.CaptionEntities != nil {
		for _, e := range *m.CaptionEntities {
			e.ComputeText(m.Text)
		}
	}
	return *m.CaptionEntities
}

// returns slice of new chat members. if there's no new members, returns empty slice
func (m Message) GetNewChatMembers() []User {
	if *m.NewChatMembers != nil {
		return *m.NewChatMembers
	}
	return []User{}
}

func (m Message) HasSticker() bool {
	return m.Sticker != nil
}

func (m Message) IsGroupMessage() bool {
	return m.Chat.IsGroupChat()
}

func (m Message) IsUserMessage() bool {
	return m.Chat.IsUserChat()
}

func (m Message) IsChannelMessage() bool {
	return m.Chat.IsChannelChat()
}

func (m Message) IsSuperGroupMessage() bool {
	return m.Chat.IsSuperGroupChat()
}

func (m Message) GetChatId() int64 {
	return m.Chat.Id
}

func (m Message) HasText() bool {
	return m.Text != "" && strings.TrimSpace(m.Text) != ""
}

func (m Message) IsCommand() bool {
	if m.HasText() && m.Entities != nil {
		for _, en := range *m.Entities {
			if !reflect.DeepEqual(en, MessageEntity{}) && en.Offset == 0 &&
				BOTCOMMAND == en.Type {
				return true
			}
		}
	}
	return false
}

func (m Message) HasDocument() bool {
	return m.Document != nil
}

func (m Message) HasVideo() bool {
	return m.video != nil
}

func (m Message) HasAudio() bool {
	return m.Audio != nil
}

func (m Message) HasVoice() bool {
	return m.Voice != nil
}

func (m Message) IsReply() bool {
	return m.ReplyToMessage != nil
}

func (m Message) HasGame() bool {
	return m.Game != nil
}

func (m Message) HasEntities() bool {
	return m.Entities != nil && len(*m.Entities) != 0
}

func (m Message) HasPhoto() bool {
	return m.Photo != nil
}

func (m Message) HasInvoice() bool {
	return m.Invoice != nil
}

func (m Message) HasSucessfulPaymen() bool {
	return m.SuccessfulPayment != nil
}

func (m Message) HasContact() bool {
	return m.Contact != nil
}

func (m Message) HasVideoNote() bool {
	return m.VideoNote != nil
}

func (m Message) HasPassportData() bool {
	return m.PassportData != nil
}

func (m Message) HasAnimation() bool {
	return m.Animation != nil
}

func (m Message) HasPoll() bool {
	return m.Poll != nil
}

func (m Message) HasDice() bool {
	return m.Dice != nil
}

func (m Message) HasViaBot() bool {
	return m.ViaBot != nil
}

func (m Message) HasReplyMarkup() bool {
	return m.ReplyMarkup != nil
}

func (m Message) HasMessageAutoDeleteTimerChanged() bool {
	return m.MessageAutoDeleteTimerChanged != nil
}

func (m Message) HasWebAppData() bool {
	return m.WebAppData != nil
}

func (m Message) HasVideoChatStarted() bool {
	return m.VideoChatStarted != nil
}

func (m Message) HasVideoChatEnded() bool {
	return m.VideoChatEnded != nil
}

func (m Message) HasVideoChatScheduled() bool {
	return m.VideoChatScheduled != nil
}

func (m Message) HasVideoChatParticipantsInvited() bool {
	return m.VideoChatParticipantsInvited != nil
}

func (m Message) HasForumTopicCreated() bool {
	return m.ForumTopicCreated != nil
}

func (m Message) HasForumTopicClosed() bool {
	return m.ForumTopicClosed != nil
}

func (m Message) HasForumTopicReopened() bool {
	return m.ForumTopicReopened != nil
}

func (m Message) HasUserShared() bool {
	return m.UserShared != nil
}

func (m Message) HasChatShared() bool {
	return m.ChatShared != nil
}

func (m Message) HasStory() bool {
	return m.Story != nil
}

func (m Message) HasWriteAccessAllowed() bool {
	return m.WriteAccessAllowed != nil
}

func (m Message) HasReplyToStory() bool {
	return m.ReplyToStory != nil
}

func (m Message) HasBoostAdded() bool {
	return m.BoostAdded != nil
}

func (m Message) HasPaidMedia() bool {
	return m.PaidMedia != nil
}

func (m Message) HasCaption() bool {
	return m.Caption != ""
}
