package types

import (
	"reflect"

	"github.com/bigelle/tele.go/internal/assertions"
)

type MaybeInaccessibleMessage interface {
	maybeInaccessibleMessageContract()
}

type Message struct {
	MessageId                     int                            `json:"message_id"`
	Chat                          Chat                           `json:"chat"`
	MessageThreadId               *int                           `json:"message_thread_id,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	Date                          *int                           `json:"date,omitempty"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardDate                   *int                           `json:"forward_date,omitempty"`
	Text                          *string                        `json:"text,omitempty"`
	Entities                      *[]MessageEntity               `json:"entities,omitempty"`
	CaptionEntities               *[]MessageEntity               `json:"caption_entities,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	Photo                         *[]PhotoSize                   `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	NewChatMembers                *[]User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  *[]PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       *string                        `json:"caption,omitempty"`
	SuperGroupCreated             *bool                          `json:"super_group_created,omitempty"`
	MigrateToChatId               *int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             *int64                         `json:"migrate_from_chat_id,omitempty"`
	EditDate                      *int                           `json:"edit_date,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	ForwardFromMessageId          *int                           `json:"forward_from_message_id,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	AuthorSignature               *string                        `json:"author_signature,omitempty"`
	ForwardSignature              *string                        `json:"forward_signature,omitempty"`
	MediaGroupId                  *string                        `json:"media_group_id,omitempty"`
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ForwardSenderName             *string                        `json:"forward_sender_name,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	IsAutomaticForward            *bool                          `json:"is_automatic_forward,omitempty"`
	HasProtectedContent           *bool                          `json:"has_protected_content,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	IsTopicMessage                *bool                          `json:"is_topic_message,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler,omitempty"`
	UserShared                    *SharedUser                    `json:"user_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	ExternalReplyInfo             *ExternalReplyInfo             `json:"external_reply_info,omitempty"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	SenderBoostCount              *int                           `json:"sender_boost_count,omitempty"`
	BusinessConnectionId          *string                        `json:"business_connection_id,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	IsFromOffline                 *bool                          `json:"is_from_offline,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	EffectId                      *string                        `json:"effect_id,omitempty"`
	ShowCaptionAboveMedia         *bool                          `json:"show_caption_above_media,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
}

func (m Message) maybeInaccessibleMessageContract() {}

func (m *Message) GetEntities() []MessageEntity {
	if m.Entities != nil {
		for _, e := range *m.Entities {
			e.ComputeText(*m.Text)
		}
	}
	return *m.Entities
}

func (m *Message) GetCaptionEntities() []MessageEntity {
	if m.CaptionEntities != nil {
		for _, e := range *m.CaptionEntities {
			e.ComputeText(*m.Text)
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

func (m Message) IsCommand() bool {
	if !assertions.IsStringEmpty(*m.Text) && m.Entities != nil {
		for _, en := range *m.Entities {
			if !reflect.DeepEqual(en, MessageEntity{}) && en.Offset == 0 &&
				en.Type == "bt_command" {
				return true
			}
		}
	}
	return false
}

type InaccesibleMessage struct {
	Chat      *Chat `json:"chat"`
	MessageId int   `json:"message_id"`
	Date      int   `json:"date"`
}

func (i InaccesibleMessage) maybeInaccessibleMessageContract() {}

type MessageEntity struct {
	Type          string  `json:"type"`
	Offset        int     `json:"offset"`
	Length        int     `json:"length"`
	Url           *string `json:"url,omitempty"`
	User          *User   `json:"user,omitempty"`
	Language      *string `json:"language,omitempty"`
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
	Text          *string `json:"text,omitempty"`
}

func (m *MessageEntity) ComputeText(message string) {
	if message != "" {
		*m.Text = message[m.Offset : m.Offset+(m.Length)]
	}
}

type MessageId struct {
	MessageId int64 `json:"message_id"`
}
