package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/bigelle/tele.go/assertions"
)

type Update struct {
	UpdateId               int                          `json:"update_id"`
	Message                *Message                     `json:"message,omitempty"`
	ChannelPost            *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost      *Message                     `json:"edited_channel_post,omitempty"`
	EditedMessage          *Message                     `json:"edited_message,omitempty"`
	BusinessConnection     *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage        *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage  *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessage *BusinessMessagesDeleted     `json:"deleted_business_message,omitempty"`
	MessageReaction        *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount   *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery            *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineQuery      *ChosenInlineResult          `json:"chosen_inline_query,omitempty"`
	CallbackQuery          *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery          *ShippingQuery               `json:"shipping_query,omitempty"`
	PaidMediaPurchased     *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	PreCheckoutQuery       *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	Poll                   *Poll                        `json:"poll,omitempty"`
	PollAnswer             *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember           *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember             *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest        *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost              *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost       *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

type User struct {
	Id                      int64   `json:"id"`
	FirstName               string  `json:"first_name"`
	IsBot                   bool    `json:"is_bot"`
	LastName                *string `json:"last_name,omitempty"`
	UserName                *string `json:"user_name,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportInlineQueries    *bool   `json:"support_inline_queries,omitempty"`
	IsPremium               *bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`
}

type Chat struct {
	Id        int64   `json:"id"`
	Type      string  `json:"type"`
	Title     *string `json:"title,omitempty"`
	UserName  *string `json:"user_name,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	IsForum   *bool   `json:"is_forum,omitempty"`
}

type ChatFullInfo struct {
	Id                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              *string               `json:"title,omitempty"`
	UserName                           *string               `json:"user_name,omitempty"`
	FirstName                          *string               `json:"first_name,omitempty"`
	LastName                           *string               `json:"last_name,omitempty"`
	IsForum                            *bool                 `json:"is_forum,omitempty"`
	AccentColorId                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    *[]string             `json:"active_usernames,omitempty"`
	BirthDate                          *BirthDate            `json:"birth_date,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	AvailableReactions                 *[]ReactionType       `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiId            *string               `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorId               *bool                 `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiId     *string               `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiId           *string               `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          *bool                 `json:"emoji_status_expiration_date,omitempty"`
	Bio                                *string               `json:"bio,omitempty"`
	HasPrivateForwards                 *bool                 `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      *bool                 `json:"join_by_request,omitempty"`
	Description                        *string               `json:"description,omitempty"`
	InviteLink                         *string               `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      *int                  `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       *string               `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   *bool                 `json:"has_hidden_members,omitempty"`
	HasProtectedCount                  *bool                 `json:"has_protected_count,omitempty"`
	HasVisibleHistory                  *bool                 `json:"has_visible_history,omitempty"`
	StickerSetName                     *string               `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatId                       *int64                `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
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

// returns slice of new chat members. if there's no new members, returns empty slice
func (m Message) GetNewChatMembers() []User {
	if *m.NewChatMembers != nil {
		return *m.NewChatMembers
	}
	return []User{}
}

func (m Message) IsCommand() bool {
	if len(*m.Text) != 0 && m.Entities != nil {
		for _, en := range *m.Entities {
			if !reflect.DeepEqual(en, MessageEntity{}) && en.Offset == 0 &&
				en.Type == "bt_command" {
				return true
			}
		}
	}
	return false
}

type MessageId int64

type InaccesibleMessage struct {
	Chat      *Chat `json:"chat"`
	MessageId int   `json:"message_id"`
	Date      int   `json:"date"`
}

func (i InaccesibleMessage) maybeInaccessibleMessageContract() {}

type MaybeInaccessibleMessage interface {
	maybeInaccessibleMessageContract()
}

type MessageEntity struct {
	Type          string  `json:"type"`
	Offset        int     `json:"offset"`
	Length        int     `json:"length"`
	Url           *string `json:"url,omitempty"`
	User          *User   `json:"user,omitempty"`
	Language      *string `json:"language,omitempty"`
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
}

type TextQuote struct {
	Text     string           `json:"text"`
	Position int              `json:"position"`
	Entities *[]MessageEntity `json:"entities,omitempty"`
	IsManual *bool            `json:"is_manual,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageId          *int                `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              *[]PhotoSize        `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    *bool               `json:"has_media_spoiler,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

type ReplyParameters struct {
	MessageId                int              `json:"message_id"`
	ChatId                   *string          `json:"chat_id,omitempty"`
	AllowSendingWithoutReply *bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    *string          `json:"quote,omitempty"`
	QuoteParseMode           *string          `json:"quote_parse_mode,omitempty"`
	QuoteEntities            *[]MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            *int             `json:"quote_position,omitempty"`
}

func (r ReplyParameters) Validate() error {
	if err := assertions.ParamNotEmpty(*r.ChatId, "ChatId"); err != nil {
		return err
	}
	return nil
}

type MessageOrigin struct {
	MessageOriginInterface
}

type MessageOriginInterface interface {
	messageOriginContract()
	Validate() error
}

func (m MessageOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.MessageOriginInterface)
}

func (m *MessageOrigin) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "user":
		tmp := MessageOriginUser{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	case "hidden_user":
		tmp := MessageOriginHiddenUser{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	case "chat":
		tmp := MessageOriginChat{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	case "channel":
		tmp := MessageOriginChannel{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MessageOriginInterface = tmp
		return nil
	default:
		return errors.New("type must be user, hidden_user, chat or channel")
	}
}

type MessageOriginChannel struct {
	Type            string  `json:"type"`
	Date            *int    `json:"date"`
	Chat            *Chat   `json:"chat"`
	MessageId       *int    `json:"message_id"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

func (m MessageOriginChannel) messageOriginContract() {}

func (m MessageOriginChannel) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return assertions.ErrInvalidParam("date parameter can't be empty")
	}
	if m.Chat == nil {
		return assertions.ErrInvalidParam("chat parameter can't be empty")
	}
	if m.MessageId == nil {
		return assertions.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

type MessageOriginChat struct {
	Type            string  `json:"type"`
	Date            *int    `json:"date"`
	SenderChat      *Chat   `json:"sender_chat"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

func (m MessageOriginChat) messageOriginContract() {}

func (m MessageOriginChat) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return assertions.ErrInvalidParam("date parameter can't be empty")
	}
	if m.SenderChat == nil {
		return assertions.ErrInvalidParam("chat parameter can't be empty")
	}
	return nil
}

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           *int   `json:"date"`
	SenderUsername string `json:"sender_username"`
}

func (m MessageOriginHiddenUser) messageOriginContract() {}

func (m MessageOriginHiddenUser) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return assertions.ErrInvalidParam("date parameter can't be empty")
	}
	if err := assertions.ParamNotEmpty(m.SenderUsername, "sender_username"); err != nil {
		return err
	}
	return nil
}

type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       *int   `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

func (m MessageOriginUser) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if m.Date == nil {
		return assertions.ErrInvalidParam("date parameter can't be empty")
	}
	if m.SenderUser == nil {
		return assertions.ErrInvalidParam("sender_user parameter can't be empty")
	}
	return nil
}

func (m MessageOriginUser) messageOriginContract() {}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FilePath     string `json:"file_path"`
	FileSize     *int   `json:"file_size,omitempty"`
}

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Performer    *string    `json:"performer,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Story struct {
	Chat Chat `json:"chat"`
	Id   int  `json:"id"`
}

type Video struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

type Voice struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     *int    `json:"file_size,omitempty"`
}

type PaidMediaInfo struct {
	StarCount string      `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

type PaidMedia struct {
	PaidMediaInterface
}

type PaidMediaInterface interface {
	paidMediaContract()
	Validate() error
}

func (p PaidMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.PaidMediaInterface)
}

func (p *PaidMedia) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "preview":
		tmp := PaidMediaPreview{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PaidMediaInterface = tmp
	case "photo":
		tmp := PaidMediaPhoto{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PaidMediaInterface = tmp
	case "video":
		tmp := PaidMediaVideo{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PaidMediaInterface = tmp
	default:
		return errors.New("type must be preview, photo, video")
	}
	return nil
}

type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

func (p PaidMediaPhoto) paidMediaContract() {}

func (p PaidMediaPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if len(p.Photo) == 0 {
		return assertions.ErrInvalidParam("photo parameter can't be empty")
	}
	return nil
}

type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    *int   `json:"width,omitempty"`
	Height   *int   `json:"height,omitempty"`
	Duration *int   `json:"duration,omitempty"`
}

func (p PaidMediaPreview) paidMediaContract() {}

func (p PaidMediaPreview) Validate() error {
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	return nil
}

type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video *Video `json:"video"`
}

func (p PaidMediaVideo) paidMediaContract() {}

func (p PaidMediaVideo) Validate() error {
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if p.Video == nil {
		return fmt.Errorf("video parameter can't be empty")
	}
	return nil
}

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserId      *int64  `json:"user_id,omitempty"`
	VCard       *string `json:"v_card,omitempty"`
}

type Dice struct {
	Value int    `json:"value"`
	Emoji string `json:"emoji"`
}

type PollOption struct {
	Text         string           `json:"text"`
	VoterCount   int              `json:"voter_count"`
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
}

type InputPollOption struct {
	Text          string           `json:"text"`
	TextParseMode *string          `json:"text_parse_mode,omitempty"`
	TextEntities  *[]MessageEntity `json:"text_entities,omitempty"`
}

func (i InputPollOption) Validate() error {
	if len(i.Text) < 1 || len(i.Text) > 100 {
		return fmt.Errorf("text must be between 1 and 100 characters")
	}
	if len(*i.TextEntities) != 0 && len(*i.TextParseMode) != 0 {
		return fmt.Errorf("parse mode and entities can't be used together")
	}
	return nil
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIds []int  `json:"option_ids"`
	VoterChat *Chat  `json:"voter_chat"`
}

type Poll struct {
	Id                   string           `json:"id"`
	Question             string           `json:"question"`
	Options              []PollOption     `json:"options"`
	TotalVoterCount      int              `json:"total_voter_count"`
	IsClosed             bool             `json:"is_closed"`
	IsAnonymous          bool             `json:"is_anonymous"`
	Type                 string           `json:"type"`
	AllowMultipleAnswers bool             `json:"allow_multiple_answers"`
	QuestionEntities     []MessageEntity  `json:"question_entities"`
	CorrectOptionId      *int             `json:"correct_option_id,omitempty"`
	OpenPeriod           *int             `json:"open_period,omitempty"`
	CloseDate            *int             `json:"close_date,omitempty"`
	Explanation          *string          `json:"explanation,omitempty"`
	ExplanationEntities  *[]MessageEntity `json:"explanation_entities,omitempty"`
}

type Location struct {
	Longitude            float64  `json:"longitude"`
	Latitude             float64  `json:"latitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareId    *string  `json:"foursquare_id,omitempty"`
	FourSquareType  *string  `json:"four_square_type,omitempty"`
	GooglePlaceId   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

type BackgroundFill struct {
	BackgroundFillInterface
}

type BackgroundFillInterface interface {
	backgroundFillContract()
}

func (b BackgroundFill) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BackgroundFillInterface)
}

func (b *BackgroundFill) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "freeform_gradient":
		tmp := BackgroundFillFreeformGradient{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundFillInterface = tmp
	case "gradient":
		tmp := BackgroundFillGradient{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundFillInterface = tmp
	case "solid":
		tmp := BackgroundFillSolid{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundFillInterface = tmp
	default:
		return errors.New("type must be freeform_gradient, gradient or solid")
	}
	return nil
}

type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`
	Colors []int  `json:"colors"`
}

func (b BackgroundFillFreeformGradient) backgroundFillContract() {}

type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

func (b BackgroundFillGradient) backgroundFillContract() {}

type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

func (b BackgroundFillSolid) backgroundFillContract() {}

type BackgroundType struct {
	BackgroundTypeInterface
}

type BackgroundTypeInterface interface {
	backgroundTypeContract()
}

func (b BackgroundType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BackgroundTypeInterface)
}

func (b *BackgroundType) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "fill":
		tmp := BackgroundTypeFill{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	case "chat_theme":
		tmp := BackgroundTypeChatTheme{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	case "pattern":
		tmp := BackgroundTypePattern{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	case "wallpaper":
		tmp := BackgroundTypeWallpaper{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BackgroundTypeInterface = tmp
	default:
		return errors.New("type must be fill, chat_theme, pattern or wallpaper")
	}

	return nil
}

type BackgroundTypeFill struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

func (b BackgroundTypeFill) backgroundTypeContract() {}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

func (b BackgroundTypeChatTheme) backgroundTypeContract() {}

type BackgroundTypeWallpaper struct {
	Type             string   `json:"type"`
	Document         Document `json:"document"`
	DarkThemeDimming int      `json:"dark_theme_dimming"`
	IsBlurred        *bool    `json:"is_blurred,omitempty"`
	IsMoving         *bool    `json:"is_moving,omitempty"`
}

func (b BackgroundTypeWallpaper) backgroundTypeContract() {}

type BackgroundTypePattern struct {
	Type       string         `json:"type"`
	Document   Document       `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int            `json:"intensity"`
	IsInverted *bool          `json:"is_inverted,omitempty"`
	IsMoving   *bool          `json:"is_moving,omitempty"`
}

func (b BackgroundTypePattern) backgroundTypeContract() {}

type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

type ForumTopicClosed struct {
}

type ForumTopicCreated struct {
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicEdited struct {
	Name              *string `json:"name,omitempty"`
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

// placeholder for event
type ForumTopicReopened struct {
}

// placeholder for event
type GeneralForumTopicHidden struct {
}

// placeholder for event
type GeneralForumTopicUnhidden struct {
}

type SharedUser struct {
	UserId    int64        `json:"user_id"`
	FirstName *string      `json:"first_name,omitempty"`
	LastName  *string      `json:"last_name,omitempty"`
	Username  *string      `json:"username,omitempty"`
	Photo     *[]PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestId string       `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

type ChatShared struct {
	RequestId string       `json:"request_id"`
	ChatId    int64        `json:"chat_id"`
	Title     *string      `json:"title,omitempty"`
	Username  *string      `json:"username,omitempty"`
	Photo     *[]PhotoSize `json:"photo,omitempty"`
}

type WriteAccessAllowed struct {
	FromRequest        *bool   `json:"from_request,omitempty"`
	WebAppName         *string `json:"web_app_name,omitempty"`
	FromAttachmentMenu *bool   `json:"from_attachment_menu,omitempty"`
}

// placeholder for event
type VideoChatStarted struct {
}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

type Giveaway struct {
	Chats                         []Chat    `json:"chats"`
	WinnerSelectionDate           int       `json:"winner_selection_date"`
	WinnerCount                   int       `json:"winner_count"`
	OnlyNewMembers                *bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              *bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              *string   `json:"prize_description,omitempty"`
	CountryCodes                  *[]string `json:"country_codes,omitempty"`
	PrizeStarCount                *int      `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      *bool    `json:"is_star_giveaway,omitempty"`
}

type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`
	GiveawayMessageId             int     `json:"giveaway_message_id"`
	WinnersSelectionDate          int     `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []User  `json:"winners"`
	AdditionalChatCount           *int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                *int    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`
	PrizeDescription              *string `json:"prize_description,omitempty"`
}

type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`
	UrlFileId        *string `json:"url_file_id,omitempty"`
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`
}

func (l LinkPreviewOptions) Validate() error {
	if *l.PreferLargeMedia && *l.PreferSmallMedia {
		return fmt.Errorf("PreferSmallMedia and PreferLargeMedia parameters are mutual exclusive")
	}
	return nil
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Plotos     [][]PhotoSize `json:"plotos"`
}

type File struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FileSize     *int64  `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

func (f File) GetFileUrl(botToken string) (string, error) {
	if botToken == "" || strings.TrimSpace(botToken) == "" {
		return "", fmt.Errorf("bot token can't be empty")
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, *f.FilePath), nil
}

type WebAppInfo struct {
	Url string `json:"url"`
}

func (w WebAppInfo) Validate() error {
	if err := assertions.ParamNotEmpty(w.Url, "Url"); err != nil {
		return err
	}
	return nil
}

type ReplyMarkup struct {
	ReplyMarkupInterface
}

type ReplyMarkupInterface interface {
	replyKeyboardContract()
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`
	Selective             *bool              `json:"selective,omitempty"`
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"`
	IsPersistent          *bool              `json:"is_persistent,omitempty"`
}

func (f ReplyKeyboardMarkup) replyKeyboardContract() {}

func (r ReplyKeyboardMarkup) Validate() error {
	if len(*r.InputFieldPlaceholder) < 1 || len(*r.InputFieldPlaceholder) > 64 {
		return fmt.Errorf("InputFieldPlaceholder parameter must be between 1 and 64 characters")
	}
	for _, row := range r.Keyboard {
		for _, key := range row {
			if err := key.Validate(); err != nil {
				return err
			}
		}

	}
	return nil
}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
	RequestContact  *bool                       `json:"request_contact,omitempty"`
	RequestLocation *bool                       `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
}

func (k KeyboardButton) Validate() error {
	if err := assertions.ParamNotEmpty(k.Text, "Text"); err != nil {
		return err
	}

	requestsProvided := 0
	if *k.RequestContact {
		requestsProvided++
	}
	if *k.RequestLocation {
		requestsProvided++
	}
	if k.WebApp != nil {
		if err := k.WebApp.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestPoll != nil {
		if err := k.RequestPoll.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestChat != nil {
		if err := k.RequestChat.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if k.RequestUsers != nil {
		if err := k.RequestUsers.Validate(); err != nil {
			return err
		}
		requestsProvided++
	}
	if requestsProvided > 1 {
		return fmt.Errorf(
			"RequestContact, RequestLocation, WebApp, RequestPoll, RequestUser, RequestChat and RequestUsers are mutually exclusive",
		)
	}

	return nil
}

type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

// currently does not do anything
func (k KeyboardButtonPollType) Validate() error {
	return nil
}

type KeyboardButtonRequestChat struct {
	RequestId                   string                       `json:"request_id"`
	ChatIsChannel               bool                         `json:"chat_is_channel"`
	ChatIsForum                 *bool                        `json:"chat_is_forum,omitempty"`
	ChatHasUsername             *bool                        `json:"chat_has_username,omitempty"`
	ChatIsCreated               *bool                        `json:"chat_is_created,omitempty"`
	UserAdministratorRights     *ChatAdministratorRights     `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights      *ChatAdministratorRights     `json:"bot_administrator_rights,omitempty"`
	BotIsMember                 *bool                        `json:"bot_is_member,omitempty"`
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	RequestTitle                *bool                        `json:"request_title,omitempty"`
	RequestUsername             *bool                        `json:"request_username,omitempty"`
	RequestPhoto                *bool                        `json:"request_photo,omitempty"`
}

func (k KeyboardButtonRequestChat) Validate() error {
	if err := assertions.ParamNotEmpty(k.RequestId, "RequestId"); err != nil {
		return err
	}
	return nil
}

type KeyboardButtonRequestUsers struct {
	RequestId       string `json:"request_id"`
	UserIsBot       *bool  `json:"user_is_bot,omitempty"`
	UserIsPremium   *bool  `json:"user_is_premium,omitempty"`
	MaxQuantity     *int   `json:"max_quantity,omitempty"`
	RequestName     *bool  `json:"request_name,omitempty"`
	RequestUsername *bool  `json:"request_username,omitempty"`
	RequestPhoto    *bool  `json:"request_photo,omitempty"`
}

func (k KeyboardButtonRequestUsers) Validate() error {
	if err := assertions.ParamNotEmpty(k.RequestId, "RequestId"); err != nil {
		return err
	}
	if *k.MaxQuantity < 1 || *k.MaxQuantity > 10 {
		return fmt.Errorf("MaxQuantity parameter must be between 1 and 10")
	}
	return nil
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

func (f ReplyKeyboardRemove) replyKeyboardContract() {}

type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"keyboard"`
}

func (f InlineKeyboardMarkup) replyKeyboardContract() {}

func (m InlineKeyboardMarkup) Validate() error {
	for _, row := range m.Keyboard {
		for _, key := range row {
			if err := key.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          *string                      `json:"url,omitempty"`
	CallbackData                 *string                      `json:"callback_data,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	Pay                          *bool                        `json:"pay,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
}

func (b InlineKeyboardButton) Validate() error {
	if err := assertions.ParamNotEmpty(b.Text, "Text"); err != nil {
		return err
	}
	if b.LoginUrl != nil {
		if err := (*b.LoginUrl).Validate(); err != nil {
			return err
		}
	}
	if b.WebApp != nil {
		if err := (*b.WebApp).Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginUrl struct {
	Url               string  `json:"url"`
	ForwardText       *string `json:"forward_text,omitempty"`
	BotUsername       *string `json:"bot_username,omitempty"`
	RequestWriteAcess *bool   `json:"request_write_acess,omitempty"`
}

func (l LoginUrl) Validate() error {
	if err := assertions.ParamNotEmpty(l.Url, "Url"); err != nil {
		return err
	}
	return nil
}

type SwitchInlineQueryChosenChat struct {
	RequestId         *string `json:"request_id,omitempty"`
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	Text string
}

type CallbackQuery struct {
	Id              string                    `json:"id"`
	From            User                      `json:"from"`
	ChatInstance    string                    `json:"chat_instance"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageId *string                   `json:"inline_message_id,omitempty"`
	Data            *string                   `json:"data,omitempty"`
	GameShortName   *string                   `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	Selective             *bool   `json:"selective,omitempty"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
}

func (f ForceReply) replyKeyboardContract() {}

func (f ForceReply) Validate() error {
	if len(*f.InputFieldPlaceholder) < 1 || len(*f.InputFieldPlaceholder) > 64 {
		return fmt.Errorf("InputFieldPlaceholder parameter must be between 1 and 64 characters")
	}
	return nil
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name,omitempty"`
	ExpireDate              *int    `json:"expire_date,omitempty"`
	MemberLimit             *bool   `json:"member_limit,omitempty"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"`
	SubscriptionPeriod      *int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       *int    `json:"subscription_price,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous         bool  `json:"is_anonymous"`
	CanManageChat       bool  `json:"can_manage_chat"`
	CanDeleteMessages   bool  `json:"can_delete_messages"`
	CanManageVideoChats bool  `json:"can_manage_video_chats"`
	CanRestrictMembers  bool  `json:"can_restrict_members"`
	CanPromoteMembers   bool  `json:"can_promote_members"`
	CanChangeInfo       bool  `json:"can_change_info"`
	CanInviteUsers      bool  `json:"can_invite_users"`
	CanPostStories      bool  `json:"can_post_stories"`
	CanEditStories      bool  `json:"can_edit_stories"`
	CanDeleteStories    bool  `json:"can_delete_stories"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
}

type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          *bool           `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatMember struct {
	ChatMemberInterface
}

type ChatMemberInterface interface {
	GetChatMemberStatus() string
}

func (c *ChatMember) UnmarshalJSON(data []byte) error {
	var raw struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Status {
	case "administrator":
		tmp := ChatMemberAdministrator{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "member":
		tmp := ChatMemberMember{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "owner":
		tmp := ChatMemberOwner{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "restricted":
		tmp := ChatMemberRestricted{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "kicked":
		tmp := ChatMemberBanned{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "left":
		tmp := ChatMemberLeft{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	default:
		fmt.Println(raw.Status)
		return errors.New(
			"status must be administrator, member, owner, restricted, banned or left",
		)
	}

	return nil
}

type ChatMemberAdministrator struct {
	Status      string `json:"status"`
	User        User   `json:"user"`
	CanBeEdited bool   `json:"can_be_edited"`
	ChatAdministratorRights
	CustomTitle *string `json:"custom_title,omitempty"`
}

func (c ChatMemberAdministrator) GetChatMemberStatus() string { return c.Status }

type ChatMemberMember struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate *int   `json:"until_date,omitempty"`
}

func (c ChatMemberMember) GetChatMemberStatus() string { return c.Status }

type ChatMemberOwner struct {
	Status      string  `json:"status"`
	User        *User   `json:"user"`
	IsAnonymous bool    `json:"is_anonymous"`
	CustomTitle *string `json:"custom_title,omitempty"`
}

func (c ChatMemberOwner) GetChatMemberStatus() string { return c.Status }

type ChatMemberRestricted struct {
	Status   string `json:"status"`
	User     *User  `json:"user"`
	IsMember bool   `json:"is_member"`
	ChatPermissions
	UntilDate int `json:"until_date"`
}

func (c ChatMemberRestricted) GetChatMemberStatus() string { return c.Status }

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate int    `json:"until_date"`
}

func (c ChatMemberBanned) GetChatMemberStatus() string { return c.Status }

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

func (c ChatMemberLeft) GetChatMemberStatus() string { return c.Status }

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	User       User            `json:"user"`
	UserChatId int64           `json:"user_chat_id"`
	Date       int             `json:"date"`
	Bio        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     *bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebpagePreviews *bool `json:"can_add_webpage_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
}

type BirthDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message *string  `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHours struct {
	TimeZone     string                         `json:"time_zone"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

type ReactionType struct {
	ReactionTypeInterface
}

type ReactionTypeInterface interface {
	reactionTypeEmojiContract()
}

func (r ReactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ReactionTypeInterface)
}

func (r *ReactionType) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "emoji":
		tmp := ReactionTypeEmoji{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		r.ReactionTypeInterface = tmp
	case "custom_emoji":
		tmp := ReactionTypeCustomEmoji{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		r.ReactionTypeInterface = tmp
	case "paid":
		tmp := ReactionTypePaid{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		r.ReactionTypeInterface = tmp
	default:
		return errors.New("type must be emoji, paid or custom_emoji")
	}
	return nil
}

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"` // = CustomEmojiType
	CustomEmojiId string `json:"custom_emoji_id"`
}

func (r ReactionTypeCustomEmoji) reactionTypeEmojiContract() {}

func (r ReactionTypeCustomEmoji) Vaidate() error {
	if err := assertions.ParamNotEmpty(r.CustomEmojiId, "custom_emoji_id"); err != nil {
		return err
	}
	if r.Type != "custom_emoji" {
		return assertions.ErrInvalidParam("type must be \"custom_emoji\"")
	}
	return nil
}

type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

func (r ReactionTypeEmoji) reactionTypeEmojiContract() {}

func (r ReactionTypeEmoji) Validate() error {
	if err := assertions.ParamNotEmpty(r.Emoji, "emoji"); err != nil {
		return err
	}
	if r.Type != "emoji" {
		return assertions.ErrInvalidParam("type must be \"emoji\"")
	}
	return nil
}

type ReactionTypePaid struct {
	Type string `json:"type"`
}

func (r ReactionTypePaid) reactionTypeEmojiContract() {}

func (r ReactionTypePaid) Validate() error {
	if r.Type != "paid" {
		return assertions.ErrInvalidParam("type must be\"paid\"")
	}
	return nil
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageId int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageId   int            `json:"message_id"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
}

type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

type ForumTopic struct {
	MessageThreadId   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (b BotCommand) Validate() error {
	if err := assertions.ParamNotEmpty(b.Command, "command"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(b.Description, "description"); err != nil {
		return err
	}
	return nil
}

type BotCommandScope struct {
	BotCommandScopeInterface
}

type BotCommandScopeInterface interface {
	botCommandScopeContract()
}

func (b BotCommandScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BotCommandScopeInterface)
}

func (b *BotCommandScope) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "all_chat_administrators":
		tmp := BotCommandScopeAllChatAdministrators{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "all_group_chats":
		tmp := BotCommandScopeAllGroupChats{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "all_private_chats":
		tmp := BotCommandScopeAllPrivateChats{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "chat":
		tmp := BotCommandScopeChat{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "chat_administrators":
		tmp := BotCommandScopeChatAdministrators{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "chat_member":
		tmp := BotCommandScopeChatMember{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	case "default":
		tmp := BotCommandScopeDefault{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		b.BotCommandScopeInterface = tmp
	default:
		return errors.New(
			"type must be all_chat_administrators, all_group_chats, all_private_chats, chat, chat_administrators, chat_member or default",
		)
	}
	return nil
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllChatAdministrators) botCommandScopeContract() {}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllGroupChats) botCommandScopeContract() {}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllPrivateChats) botCommandScopeContract() {}

type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
}

func (b BotCommandScopeChat) botCommandScopeContract() {}

func (b BotCommandScopeChat) Validate() error {
	if err := assertions.ParamNotEmpty(b.ChatId, "chat_id"); err != nil {
		return err
	}
	return nil
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
}

func (b BotCommandScopeChatAdministrators) botCommandScopeContract() {}

func (b BotCommandScopeChatAdministrators) Validate() error {
	if err := assertions.ParamNotEmpty(b.ChatId, "chat_id"); err != nil {
		return err
	}
	return nil
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
	UserId int64  `json:"user_id"`
}

func (b BotCommandScopeChatMember) botCommandScopeContract() {}

func (b BotCommandScopeChatMember) Validate() error {
	if err := assertions.ParamNotEmpty(b.ChatId, "chat_id"); err != nil {
		return err
	}
	if b.UserId == 0 {
		return fmt.Errorf("UserId parameter can't be empty")
	}
	return nil
}

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

func (b BotCommandScopeDefault) botCommandScopeContract() {}

type BotName struct {
	Name string
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

type MenuButton struct {
	MenuButtonInterface
}

type MenuButtonInterface interface {
	menuButtonContract()
	Validate() error
}

func (m MenuButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.MenuButtonInterface)
}

func (m *MenuButton) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "commands":
		tmp := MenuButtonCommands{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MenuButtonInterface = tmp
	case "web_app":
		tmp := MenuButtonWebApp{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MenuButtonInterface = tmp
	case "default":
		tmp := MenuButtonDefault{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		m.MenuButtonInterface = tmp
	default:
		return errors.New("unknow type " + raw.Type + ", type must be commands, web_app or default")
	}
	return nil
}

type MenuButtonCommands struct {
	Type string `json:"type"`
}

func (m MenuButtonCommands) Validate() error {
	return assertions.ParamNotEmpty(m.Type, "Type")
}

func (m MenuButtonCommands) menuButtonContract() {}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

func (m MenuButtonDefault) Validate() error {
	return assertions.ParamNotEmpty(m.Type, "Type")
}

func (m MenuButtonDefault) menuButtonContract() {}

type MenuButtonWebApp struct {
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	WebAppInfo WebAppInfo `json:"web_app_info"`
}

func (m MenuButtonWebApp) menuButtonContract() {}

func (m MenuButtonWebApp) Validate() error {
	if err := assertions.ParamNotEmpty(m.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(m.Text, "text"); err != nil {
		return err
	}
	if err := m.WebAppInfo.Validate(); err != nil {
		return err
	}
	return nil
}

type ChatBoostSource struct {
	ChatBoostSourceInterface
}

type ChatBoostSourceInterface interface {
	GetBoostSource() string
}

func (c *ChatBoostSource) UnmarshalJSON(data []byte) error {
	var raw struct {
		Source string
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Source {
	case "gift_code":
		tmp := ChatBoostSourceGiftCode{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatBoostSourceInterface = tmp
	case "giveaway":
		tmp := ChatBoostSourceGiveaway{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatBoostSourceInterface = tmp
	case "premium":
		tmp := ChatBoostSourcePremium{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatBoostSourceInterface = tmp
	default:
		return errors.New("source must be gift_code, giveaway or premium")
	}
	return nil
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

func (c ChatBoostSourceGiftCode) GetBoostSource() string { return c.Source }

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageId string `json:"giveaway_message_id"`
	User              *User  `json:"user,omitempty"`
	IsUnclaimed       *bool  `json:"is_unclaimed,omitempty"`
	PrizeStarCount    *int   `json:"prize_star_count,omitempty"`
}

func (c ChatBoostSourceGiveaway) GetBoostSource() string { return c.Source }

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

func (c ChatBoostSourcePremium) GetBoostSource() string { return c.Source }

type ChatBoost struct {
	BoostId        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostId    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type BusinessConnection struct {
	Id         string `json:"id"`
	User       *User  `json:"user"`
	UserChatId int64  `json:"user_chat_id"`
	Date       int    `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionId string `json:"business_connection_id"`
	Chat                 Chat   `json:"chat"`
	MessageIds           []int  `json:"message_ids"`
}

type ApiResponse[T any] struct {
	Ok          bool                `json:"ok"`
	ErrorCode   int                 `json:"error_code"`
	Description *string             `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      T                   `json:"result"`
}

type ResponseParameters struct {
	MigrateToChatId *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int   `json:"retry_after,omitempty"`
}

type InputMedia struct {
	InputMediaInterface
}

type InputMediaInterface interface {
	SetInputMedia(media string, isNew bool)
	Validate() error
}

func (i InputMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.InputMediaInterface)
}

func (i *InputMedia) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "animation":
		tmp := InputMediaAnimation{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "audio":
		tmp := InputMediaAudio{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "document":
		tmp := InputMediaDocument{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "photo":
		tmp := InputMediaPhoto{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	case "video":
		tmp := InputMediaVideo{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputMediaInterface = &tmp
	default:
		return errors.New("type must be animation, audio, document, video or photo")
	}
	return nil
}

type InputMediaAnimation struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`
	Caption               *string          `json:"caption,omitempty"`
	ParseMode             *string          `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"`
	Width                 *int             `json:"width,omitempty"`
	Height                *int             `json:"height,omitempty"`
	Duration              *int             `json:"duration,omitempty"`
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`
	isNew                 bool             `json:"-"`
}

func (i *InputMediaAnimation) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaAnimation) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}
	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InputMediaAudio struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumbnail       *InputFile       `json:"thumbnail,omitempty"`
	Caption         *string          `json:"caption,omitempty"`
	ParseMode       *string          `json:"parse_mode,omitempty"`
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int             `json:"duration,omitempty"`
	Performer       *string          `json:"performer,omitempty"`
	Title           *string          `json:"title,omitempty"`
	isNew           bool             `json:"-"`
}

func (i *InputMediaAudio) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaAudio) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}
	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InputMediaDocument struct {
	Type                        string           `json:"type"`
	Media                       string           `json:"media"`
	Thumbnail                   *InputFile       `json:"thumbnail,omitempty"`
	Caption                     *string          `json:"caption,omitempty"`
	ParseMode                   *string          `json:"parse_mode,omitempty"`
	CaptionEntities             *[]MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool            `json:"disable_content_type_detection,omitempty"`
	isNew                       bool             `json:"-"`
}

func (i *InputMediaDocument) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaDocument) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}
	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type InputMediaPhoto struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Caption               *string          `json:"caption,omitempty"`
	ParseMode             *string          `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`
	isNew                 bool             `json:"-"`
}

func (i *InputMediaPhoto) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}

type InputMediaVideo struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Thumbnail             *InputFile       `json:"thumbnail,omitempty"`
	Caption               *string          `json:"caption,omitempty"`
	ParseMode             *string          `json:"parse_mode,omitempty"`
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"`
	Width                 *int             `json:"width,omitempty"`
	Height                *int             `json:"height,omitempty"`
	Duration              *int             `json:"duration,omitempty"`
	SupportsStreaming     *bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`
	isNew                 bool             `json:"-"`
}

func (i *InputMediaVideo) SetInputMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputMediaVideo) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}

type InputFile string

func (i InputFile) Validate() error {
	urlRegex := regexp.MustCompile(`^https?://`)
	attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
	switch {
	case urlRegex.MatchString(string(i)):
		return nil
	case attachmentRegex.MatchString(string(i)):
		return nil
	default:
		return assertions.ErrInvalidParam(
			"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
		)
	}
}

type InputPaidMedia struct {
	InputPaidMediaInterface
}

type InputPaidMediaInterface interface {
	SetInputPaidMedia(media string, isNew bool)
	Validate() error
}

func (i InputPaidMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.InputPaidMediaInterface)
}

func (i *InputPaidMedia) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "photo":
		tmp := InputPaidMediaPhoto{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputPaidMediaInterface = &tmp
	case "video":
		tmp := InputPaidMediaVideo{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i.InputPaidMediaInterface = &tmp
	default:
		return errors.New("type must be photo or video")
	}
	return nil
}

type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	isNew bool   `json:"-"`
}

func (i InputPaidMediaPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}

func (i *InputPaidMediaPhoto) SetInputPaidMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

type InputPaidMediaVideo struct {
	Type              string     `json:"type"`
	Media             string     `json:"media"`
	Thumbnail         *InputFile `json:"thumbnail,omitempty"`
	Width             *int       `json:"width,omitempty"`
	Height            *int       `json:"height,omitempty"`
	Duration          *int       `json:"duration,omitempty"`
	SupportsStreaming *bool      `json:"supports_streaming,omitempty"`
	isNew             bool       `json:"-"`
}

func (i *InputPaidMediaVideo) SetInputPaidMedia(media string, isNew bool) {
	if isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		if urlRegex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

func (i InputPaidMediaVideo) Validate() error {
	if err := assertions.ParamNotEmpty(i.Media, "Media"); err != nil {
		return err
	}

	if i.Thumbnail != nil {
		if err := i.Thumbnail.Validate(); err != nil {
			return err
		}
	}

	if i.isNew {
		urlRegex := regexp.MustCompile(`^https?://`)
		attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
		switch {
		case urlRegex.MatchString(i.Media):
			return nil
		case attachmentRegex.MatchString(i.Media):
			return nil
		default:
			return assertions.ErrInvalidParam(
				"invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
			)
		}
	}

	return nil
}
