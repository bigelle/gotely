package objects

import (
	"reflect"
	"strings"
)

type Message struct {
	MessageId                     int
	MessageThreadId               int
	from                          *User
	date                          int
	chat                          *Chat
	forwardFrom                   *User
	forwardFromChat               *Chat
	forwardDate                   int
	text                          string
	entities                      *[]MessageEntity
	captionEntities               *[]MessageEntity
	audio                         *Audio
	document                      *Document
	photo                         *[]PhotoSize
	sticker                       *Sticker
	video                         *Video
	contact                       *Contact
	location                      *Location
	venue                         *Venue
	animation                     *Animation
	pinnedMessage                 *MaybeInaccessibleMessage
	newChatMembers                *[]User
	leftChatMember                *User
	newChatTitle                  string
	newChatPhoto                  *[]PhotoSize
	deleteChatPhoto               bool
	groupChatCreated              bool
	replyToMessage                *Message
	voice                         *Voice
	caption                       string
	superGroupCreated             bool
	migrateToChatId               int64
	migrateFromChatId             int64
	editDate                      int
	game                          *Game
	forwardFromMessageId          int
	invoice                       *Invoice           // TODO: implement it
	SuccessfulPayment             *SuccessfulPayment // TODO: implement it
	videoNote                     *VideoNote
	authorSignature               string
	forwardSignature              string
	mediaGroupId                  string
	connectedWebsite              string
	passportData                  *PassportData // TODO: implement it
	forwardSenderName             string
	poll                          *Poll                          // TODO: implement it
	replyMarkup                   *InlineKeyboardMarkup          // TODO: implement it
	dice                          *Dice                          // TODO: implement it
	viaBot                        *User
	senderChat                    *Chat
	proximityAlertTriggered       *ProximityAlertTriggered       // TODO: implement it
	messageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged // TODO: implement it
	isAutomaticForward            bool
	hasProtectedContent           bool
	webAppData                    *WebAppData                   // TODO: implement it
	videoChatStarted              *VideoChatStarted             // TODO: implement it
	videoChatEnded                *VideoChatEnded               // TODO: implement it
	videoChatParticipantsInvited  *VideoChatParticipantsInvited // TODO: implement it
	videoChatScheduled            *VideoChatScheduled           // TODO: implement it
	isTopicMessage                bool
	forumTopicCreated             *ForumTopicCreated         // TODO: implement it
	forumTopicClosed              *ForumTopicClosed          // TODO: implement it
	forumTopicReopened            *ForumTopicReopened        // TODO: implement it
	forumTopicEdited              *ForumTopicEdited          // TODO: implement it
	GeneralForumTopicHidden       *GeneralForumTopicHidden   // TODO: implement it
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden // TODO: implement it
	WriteAccessAllowed            *WriteAccessAllowed        // TODO: implement it
	hasMediaSpoiler               bool
	userShared                    *UserShared         // TODO: implement it
	chatShared                    *ChatShared         // TODO: implement it
	story                         *Story              // TODO: implement it
	ExternalReplyInfo             *ExternalReplyInfo  // TODO: implement it
	forwardOrigin                 *MessageOrigin      // TODO: implement it
	LinkPreviewOptions            *LinkPreviewOptions // TODO: implement it
	quote                         *TextQuote          // TODO: implement it
	usersShared                   *UsersShared        // TODO: implement it
	GiveawayCreated               *GiveawayCreated    // TODO: implement it
	Giveaway                      *Giveaway           // TODO: implement it
	GiveawayWinners               *GiveawayWinners    // TODO: implement it
	GiveawayCompleted             *GiveawayCompleted  // TODO: implement it
	replyToStory                  *Story              // TODO: implement it
	boostAdded                    *ChatBoostAdded     // TODO: implement it
	senderBoostCount              int
	businessConnectionId          string
	senderBusinessBot             *User
	isFromOffline                 bool
	ChatBackgroundSet             *ChatBackground // TODO: implement it
	effectId                      string
	showCaptionAboveMedia         bool
	PaidMedia                     *PaidMediaInfo   // TODO: implement it
	RefundedPayment               *RefundedPayment // TODO: implement it
}

func (m *Message) GetEntities() []MessageEntity {
	if m.entities != nil {
		for _, e := range *m.entities {
			e.ComputeText(m.text)
		}
	}
	return *m.entities
}

func (m *Message) GetCaptionEntities() []MessageEntity {
	if m.captionEntities != nil {
		for _, e := range *m.captionEntities {
			e.ComputeText(m.text)
		}
	}
	return *m.captionEntities
}

// returns slice of new chat members. if there's no new members, returns empty slice
func (m Message) GetNewChatMembers() []User {
	if *m.newChatMembers != nil {
		return *m.newChatMembers
	}
	return []User{}
}

func (m Message) HasSticker() bool {
	return m.sticker != nil
}

// if json-serialized field is empty, returns default boolian value (false)
func (m Message) IsTopicMessage() bool {
	return m.isTopicMessage
}

func (m Message) IsGroupMessage() bool {
	return m.chat.IsGroupChat()
}

func (m Message) IsUserMesage() bool {
	return m.chat.IsUserChat()
}

func (m Message) IsChannelMessage() bool {
	return m.chat.IsChannelChat()
}

func (m Message) IsSuperGroupMessage() bool {
	return m.chat.IsSuperGroupChat()
}

func (m Message) GetChatId() int64 {
	return m.chat.id
}

func (m Message) HasText() bool {
	return m.text != "" && strings.TrimSpace(m.text) != ""
}

func (m Message) IsCommand() bool {
	if m.HasText() && m.entities != nil {
		for _, en := range *m.entities {
			if !reflect.DeepEqual(en, MessageEntity{}) && en.Offset == 0 &&
				BOTCOMMAND == en.Type {
				return true
			}
		}
	}
	return false
}

func (m Message) HasDocument() bool {
	return m.document != nil
}

func (m Message) HasVideo() bool {
	return m.video != nil
}

func (m Message) HasAudio() bool {
	return m.audio != nil
}

func (m Message) HasVoice() bool {
	return m.voice != nil
}

func (m Message) IsReply() bool {
	return m.replyToMessage != nil
}

func (m Message) HasGame() bool {
	return m.game != nil
}

func (m Message) HasEntities() bool {
	return m.entities != nil && len(*m.entities) != 0
}

func (m Message) HasPhoto() bool {
	return m.photo != nil
}

func (m Message) HasInvoice() bool {
	return m.invoice != nil
}

func (m Message) HasSucessfulPaymen() bool {
	return m.SuccessfulPayment != nil
}

func (m Message) HasContact() bool {
	return m.contact != nil
}

func (m Message) HasVideoNote() bool {
	return m.videoNote != nil
}

func (m Message) HasPassportData() bool {
	return m.passportData != nil
}

func (m Message) HasAnimation() bool {
	return m.animation != nil
}

func (m Message) HasPoll() bool {
	return m.poll != nil
}

func (m Message) HasDice() bool {
	return m.dice != nil
}

func (m Message) HasViaBot() bool {
	return m.viaBot != nil
}

func (m Message) HasReplyMarkup() bool {
	return m.replyMarkup != nil
}

func (m Message) HasMessageAutoDeleteTimerChanged() bool {
	return m.messageAutoDeleteTimerChanged != nil
}

func (m Message) HasWebAppData() bool {
	return m.webAppData != nil
}

func (m Message) HasVideoChatStarted() bool {
	return m.videoChatStarted != nil
}

func (m Message) HasVideoChatEnded() bool {
	return m.videoChatEnded != nil
}

func (m Message) HasVideoChatScheduled() bool {
	return m.videoChatScheduled != nil
}

func (m Message) HasVideoChatParticipantsInvited() bool {
	return m.videoChatParticipantsInvited != nil
}

func (m Message) HasForumTopicCreated() bool {
	return m.forumTopicCreated != nil
}

func (m Message) HasForumTopicClosed() bool {
	return m.forumTopicClosed != nil
}

func (m Message) HasForumTopicReopened() bool {
	return m.forumTopicReopened != nil
}

func (m Message) HasUserShared() bool {
	return m.userShared != nil
}

func (m Message) HasChatShared() bool {
	return m.chatShared != nil
}

func (m Message) HasStory() bool {
	return m.story != nil
}

func (m Message) HasWriteAccessAllowed() bool {
	return m.WriteAccessAllowed != nil
}

func (m Message) HasReplyToStory() bool {
	return m.replyToStory != nil
}

func (m Message) HasBoostAdded() bool {
	return m.boostAdded != nil
}

func (m Message) HasPaidMedia() bool {
	return m.PaidMedia != nil
}

func (m Message) HasCaption() bool {
	return m.caption != ""
}
