package objects

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

type Validable interface {
	Validate() error
}

type ErrInvalidParam string

func (e ErrInvalidParam) Error() string {
	return string(e)
}

type ErrInvalidSubtypeConversion struct {
	Expected string
	Got      string
}

func (e ErrInvalidSubtypeConversion) Error() string {
	return fmt.Sprintf("failed subtype conversion; expected: %s, got: %s", e.Expected, e.Got)
}

// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	//The update's unique identifier.
	//Update identifiers start from a certain positive number and increase sequentially.
	//This identifier becomes especially handy if you're using webhooks,
	//since it allows you to ignore repeated updates or to restore the correct update sequence,
	//should they get out of order.
	//If there are no new updates for at least a week,
	//then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateId int `json:"update_id"`

	//Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	//Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	//Optional. New version of a channel post that is known to the bot and was edited.
	//This update may at times be triggered by changes to message fields
	//that are either unavailable or not actively used by your bot.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	//Optional. New version of a message that is known to the bot and was edited.
	//This update may at times be triggered by changes to message fields
	//that are either unavailable or not actively used by your bot.
	EditedMessage *Message `json:"edited_message,omitempty"`

	//Optional. The bot was connected to or disconnected from a business account,
	//or a user edited an existing connection with the bot
	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`

	//Optional. New message from a connected business account
	BusinessMessage *Message `json:"business_message,omitempty"`

	//Optional. New version of a message from a connected business account
	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`

	//Optional. Messages were deleted from a connected business account
	DeletedBusinessMessage *BusinessMessagesDeleted `json:"deleted_business_message,omitempty"`

	//Optional. A reaction to a message was changed by a user.
	//The bot must be an administrator in the chat and must explicitly specify "message_reaction"
	//in the list of allowed_updates to receive these updates.
	//The update isn't received for reactions set by bots.
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`

	//Optional. Reactions to a message with anonymous reactions were changed.
	//The bot must be an administrator in the chat and must explicitly specify "message_reaction_count"
	//in the list of allowed_updates to receive these updates.
	//The updates are grouped and can be sent with delay up to a few minutes.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

	//Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	//Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	//Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	ChosenInlineQuery *ChosenInlineResult `json:"chosen_inline_query,omitempty"`

	//Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	//Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	//Optional. A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`

	//Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	//Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`

	//Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	//Optional. The bot's chat member status was updated in a chat.
	//For private chats, this update is received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	//Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and
	//must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	//Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	//Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`

	//Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}

// This object represents a Telegram user or bot.
type User struct {
	//Unique identifier for this user or bot.
	//This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int64 `json:"id"`

	//User's or bot's first name
	FirstName string `json:"first_name"`

	//True, if this user is a bot
	IsBot bool `json:"is_bot"`

	//Optional. User's or bot's last name
	LastName *string `json:"last_name,omitempty"`

	//Optional. User's or bot's username
	UserName *string `json:"user_name,omitempty"`

	//Optional. IETF language tag of the user's language
	LanguageCode *string `json:"language_code,omitempty"`

	//Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups *bool `json:"can_join_groups,omitempty"`

	//Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages *bool `json:"can_read_all_group_messages,omitempty"`

	//Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportInlineQueries *bool `json:"support_inline_queries,omitempty"`

	//Optional. True, if this user is a Telegram Premium user
	IsPremium *bool `json:"is_premium,omitempty"`

	//Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu *bool `json:"added_to_attachment_menu,omitempty"`

	//Optional. True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
	CanConnectToBusiness *bool `json:"can_connect_to_business,omitempty"`

	//Optional. True, if the bot has a main Web App. Returned only in getMe.
	HasMainWebApp *bool `json:"has_main_web_app,omitempty"`
}

// This object represents a chat.
type Chat struct {
	//Unique identifier for this chat. This number may have more than 32 significant bits
	//and some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int64 `json:"id"`

	//Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	//Optional. Title, for supergroups, channels and group chats
	Title *string `json:"title,omitempty"`

	//Optional. Username, for private chats, supergroups and channels if available
	UserName *string `json:"user_name,omitempty"`

	//Optional. First name of the other party in a private chat
	FirstName *string `json:"first_name,omitempty"`

	//Optional. Last name of the other party in a private chat
	LastName *string `json:"last_name,omitempty"`

	//Optional. True, if the supergroup chat is a forum (has topics enabled)
	IsForum *bool `json:"is_forum,omitempty"`
}

// This object contains full information about a chat.
type ChatFullInfo struct {
	//Unique identifier for this chat.
	//This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int64 `json:"id"`
	//Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	//Optional. Title, for supergroups, channels and group chats
	Title *string `json:"title,omitempty"`
	//Optional. Username, for private chats, supergroups and channels if available
	UserName *string `json:"user_name,omitempty"`
	//Optional. First name of the other party in a private chat
	FirstName *string `json:"first_name,omitempty"`
	//Optional. Last name of the other party in a private chat
	LastName *string `json:"last_name,omitempty"`
	//Optional. True, if the supergroup chat is a forum (has topics enabled)
	IsForum *bool `json:"is_forum,omitempty"`
	//Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview.
	//See https://core.telegram.org/bots/api#accent-colors for more details.
	AccentColorId int `json:"accent_color_id"`
	//The maximum number of reactions that can be set on a message in the chat
	MaxReactionCount int `json:"max_reaction_count"`
	//Optional. Chat photo
	Photo *ChatPhoto `json:"photo,omitempty"`
	//Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels
	ActiveUsernames *[]string `json:"active_usernames,omitempty"`
	//Optional. For private chats, the date of birth of the user
	BirthDate *BirthDate `json:"birth_date,omitempty"`
	//Optional. For private chats with business accounts, the intro of the business
	BusinessIntro *BusinessIntro `json:"business_intro,omitempty"`
	//Optional. For private chats with business accounts, the location of the business
	BusinessLocation *BusinessLocation `json:"business_location,omitempty"`
	//Optional. For private chats with business accounts, the opening hours of the business
	BusinessOpeningHours *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	//Optional. For private chats, the personal channel of the user
	PersonalChat *Chat `json:"personal_chat,omitempty"`
	//Optional. List of available reactions allowed in the chat. If omitted, then all emoji reactions are allowed.
	AvailableReactions *[]ReactionType `json:"available_reactions,omitempty"`
	//Optional. Custom emoji identifier of the emoji chosen by the chat for the reply header and link preview background
	BackgroundCustomEmojiId *string `json:"background_custom_emoji_id,omitempty"`
	//Optional. Identifier of the accent color for the chat's profile background.
	//See https://core.telegram.org/bots/api#profile-accent-colors for more details.
	ProfileAccentColorId *bool `json:"profile_accent_color_id,omitempty"`
	//Optional. Custom emoji identifier of the emoji chosen by the chat for its profile background
	ProfileBackgroundCustomEmojiId *string `json:"profile_background_custom_emoji_id,omitempty"`
	//Optional. Custom emoji identifier of the emoji status of the chat or the other party in a private chat
	EmojiStatusCustomEmojiId *string `json:"emoji_status_custom_emoji_id,omitempty"`
	//Optional. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any
	EmojiStatusExpirationDate *bool `json:"emoji_status_expiration_date,omitempty"`
	//Optional. Bio of the other party in a private chat
	Bio *string `json:"bio,omitempty"`
	//Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user
	HasPrivateForwards *bool `json:"has_private_forwards,omitempty"`
	//Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat
	HasRestrictedVoiceAndVideoMessages *bool `json:"has_restricted_voice_and_video_messages,omitempty"`
	//Optional. True, if users need to join the supergroup before they can send messages
	JoinToSendMessages *bool `json:"join_to_send_messages,omitempty"`
	//Optional. True, if all users directly joining the supergroup without using an invite link need to be approved by supergroup administrators
	JoinByRequest *bool `json:"join_by_request,omitempty"`
	//Optional. Description, for groups, supergroups and channel chats
	Description *string `json:"description,omitempty"`
	//Optional. Primary invite link, for groups, supergroups and channel chats
	InviteLink *string `json:"invite_link,omitempty"`
	//Optional. The most recent pinned message (by sending date)
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	//Optional. Default chat member permissions, for groups and supergroups
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	//Optional. True, if paid media messages can be sent or forwarded to the channel chat. The field is available only for channel chats.
	CanSendPaidMedia *bool `json:"can_send_paid_media,omitempty"`
	//Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds
	SlowModeDelay *int `json:"slow_mode_delay,omitempty"`
	//Optional. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions
	UnrestrictBoostCount *int `json:"unrestrict_boost_count,omitempty"`
	//Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds
	MessageAutoDeleteTime *int `json:"message_auto_delete_time,omitempty"`
	//Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators.
	HasAggressiveAntiSpamEnabled *string `json:"has_aggressive_anti_spam_enabled,omitempty"`
	//Optional. True, if non-administrators can only get the list of bots and administrators in the chat
	HasHiddenMembers *bool `json:"has_hidden_members,omitempty"`
	//Optional. True, if messages from the chat can't be forwarded to other chats
	HasProtectedCount *bool `json:"has_protected_count,omitempty"`
	//Optional. True, if new chat members will have access to old messages; available only to chat administrators
	HasVisibleHistory *bool `json:"has_visible_history,omitempty"`
	//Optional. For supergroups, name of the group sticker set
	StickerSetName *string `json:"sticker_set_name,omitempty"`
	//Optional. True, if the bot can change the group sticker set
	CanSetStickerSet *bool `json:"can_set_sticker_set,omitempty"`
	//Optional. For supergroups, the name of the group's custom emoji sticker set. Custom emoji from this set can be used by all users and bots in the group.
	CustomEmojiStickerSetName *string `json:"custom_emoji_sticker_set_name,omitempty"`
	//Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa;
	//for supergroups and channel chats. This identifier may be greater than 32 bits
	//and some programming languages may have difficulty/silent defects in interpreting it.
	//But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	LinkedChatId *int64 `json:"linked_chat_id,omitempty"`
	//Optional. For supergroups, the location to which the supergroup is connected
	Location *ChatLocation `json:"location,omitempty"`
}

// This object represents a message.
type Message struct {
	//Unique message identifier inside this chat.
	//In specific instances (e.g., message containing a video sent to a big chat),
	//the server might automatically schedule a message instead of sending it immediately.
	//In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageId int `json:"message_id"`
	//Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Optional. Sender of the message; may be empty for messages sent to channels.
	//For backward compatibility, if the message was sent on behalf of a chat,
	//the field contains a fake sender user in non-channel chats
	From *User `json:"from,omitempty"`
	//Optional. Sender of the message when sent on behalf of a chat.
	//For example, the supergroup itself for messages sent by its anonymous administrators or
	//a linked channel for messages automatically forwarded to the channel's discussion group.
	//For backward compatibility, if the message was sent on behalf of a chat,
	//the field from contains a fake sender user in non-channel chats.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	//Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBoostCount *int `json:"sender_boost_count,omitempty"`
	//Optional. The bot that actually sent the message on behalf of the business account.
	//Available only for outgoing messages sent on behalf of the connected business account.
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`
	//Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	Date int `json:"date"`
	//Optional. Unique identifier of the business connection from which the message was received.
	//If non-empty, the message belongs to a chat of the corresponding business account that is
	//independent from any potential bot chat which might share the same identifier.
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Chat the message belongs to
	Chat Chat `json:"chat"`
	//Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`
	//Optional. True, if the message is sent to a forum topic
	IsTopicMessage *bool `json:"is_topic_message,omitempty"`
	//Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	IsAutomaticForward *bool `json:"is_automatic_forward,omitempty"`
	//Optional. For replies in the same chat and message thread, the original message.
	//Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	//Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`
	//Optional. For replies that quote part of the original message, the quoted part of the message
	Quote *TextQuote `json:"quote,omitempty"`
	//Optional. For replies to a story, the original story
	ReplyToStory *Story `json:"reply_to_story,omitempty"`
	//Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`
	//Optional. Date the message was last edited in Unix time
	EditDate *int `json:"edit_date,omitempty"`
	//Optional. True, if the message can't be forwarded
	HasProtectedContent *bool `json:"has_protected_content,omitempty"`
	//Optional. True, if the message was sent by an implicit action, for example,
	//as an away or a greeting business message, or as a scheduled message
	IsFromOffline *bool `json:"is_from_offline,omitempty"`
	//Optional. The unique identifier of a media message group this message belongs to
	MediaGroupId *string `json:"media_group_id,omitempty"`
	//Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	AuthorSignature *string `json:"author_signature,omitempty"`
	//Optional. For text messages, the actual UTF-8 text of the message
	Text *string `json:"text,omitempty"`
	//Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities *[]MessageEntity `json:"entities,omitempty"`
	//Optional. Options used for link preview generation for the message,
	//if it is a text message and link preview options were changed
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	//Optional. Unique identifier of the message effect added to the message
	EffectId *string `json:"effect_id,omitempty"`
	//Optional. Message is an animation, information about the animation.
	//For backward compatibility, when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`
	//Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`
	//Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`
	//Optional. Message contains paid media; information about the paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`
	//Optional. Message is a photo, available sizes of the photo
	Photo *[]PhotoSize `json:"photo,omitempty"`
	//Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	//Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`
	//Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`
	//Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`
	//Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`
	//Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	Caption *string `json:"caption,omitempty"`
	//Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler *bool `json:"has_media_spoiler,omitempty"`
	//Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`
	//Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`
	//Optional. Message is a game, information about the game. More about games » https://core.telegram.org/bots/api#games
	Game *Game `json:"game,omitempty"`
	//Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`
	//Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Venue *Venue `json:"venue,omitempty"`
	//Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`
	//Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	NewChatMembers *[]User `json:"new_chat_members,omitempty"`
	//Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	//Optional. A chat title was changed to this value
	NewChatTitle *string `json:"new_chat_title,omitempty"`
	//Optional. A chat photo was change to this value
	NewChatPhoto *[]PhotoSize `json:"new_chat_photo,omitempty"`
	//Optional. Service message: the chat photo was deleted
	DeleteChatPhoto *bool `json:"delete_chat_photo,omitempty"`
	//Optional. Service message: the group has been created
	GroupChatCreated *bool `json:"group_chat_created,omitempty"`
	//Optional. Service message: the supergroup has been created.
	//This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created.
	//It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	SuperGroupCreated *bool `json:"super_group_created,omitempty"`
	//Optional. Service message: the channel has been created.
	//This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created.
	//It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated *bool `json:"channel_chat_created"`
	//Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	//Optional. The group has been migrated to a supergroup with the specified identifier.
	//This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId *int64 `json:"migrate_to_chat_id,omitempty"`
	//Optional. The supergroup has been migrated from a group with the specified identifier.
	//This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId *int64 `json:"migrate_from_chat_id,omitempty"`
	//Optional. Specified message was pinned. Note that the Message object in this field
	//will not contain further reply_to_message fields even if it itself is a reply.
	PinnedMessage *MaybeInaccessibleMessage `json:"pinned_message,omitempty"`
	//Optional. Message is an invoice for a payment, information about the invoice.
	//More about payments » https://core.telegram.org/bots/api#payments
	Invoice *Invoice `json:"invoice,omitempty"`
	//Optional. Message is a service message about a successful payment, information about the payment.
	//More about payments » https://core.telegram.org/bots/api#payments
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	//Optional. Message is a service message about a refunded payment, information about the payment.
	//More about payments » https://core.telegram.org/bots/api#payments
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`
	//Optional. Service message: users were shared with the bot
	UsersShared *UsersShared `json:"users_shared,omitempty"`
	//Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared,omitempty"`
	//Optional. The domain name of the website on which the user has logged in.
	//More about Telegram Login » https://core.telegram.org/widgets/login
	ConnectedWebsite *string `json:"connected_website,omitempty"`
	//Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu,
	//launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`
	//Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`
	//Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	ForwardFrom             *User                    `json:"forward_from,omitempty"`
	//Optional. Service message: user boosted the chat
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`
	//Optional. Service message: chat background set
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`
	//Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`
	//Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`
	//Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`
	//Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`
	//Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`
	//Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`
	//Optional. Service message: a scheduled giveaway was created
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`
	//Optional. The message is a scheduled giveaway message
	Giveaway *Giveaway `json:"giveaway,omitempty"`
	//Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`
	//Optional. Service message: a giveaway without public winners was completed
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`
	//Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	//Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`
	//Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	//Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	//Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
	//Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m Message) IsCommand() bool {
	if len(*m.Text) != 0 && m.Entities != nil {
		for _, en := range *m.Entities {
			if !reflect.DeepEqual(en, MessageEntity{}) && en.Offset == 0 &&
				en.Type == "bot_command" {
				return true
			}
		}
	}
	return false
}

// This object represents a unique message identifier.
type MessageId struct {
	//Unique message identifier. In specific instances (e.g., message containing a video sent to a big chat),
	//the server might automatically schedule a message instead of sending it immediately.
	//In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageId int
}

// This object describes a message that was deleted or is otherwise inaccessible to the bot.
type InaccesibleMessage struct {
	//Chat the message belonged to
	Chat Chat `json:"chat"`
	//Unique message identifier inside the chat
	MessageId int `json:"message_id"`
	//Always 0. The field can be used to differentiate regular and inaccessible messages.
	Date int `json:"date"`
}

// This object describes a message that can be inaccessible to the bot. It can be one of:
//
// - Message
//
// - InaccessibleMessage
type MaybeInaccessibleMessage struct {
	MessageId                     int                            `json:"message_id"`
	MessageThreadId               *int                           `json:"message_thread_id,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBoostCount              *int                           `json:"sender_boost_count,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	Date                          int                            `json:"date"`
	BusinessConnectionId          *string                        `json:"business_connection_id,omitempty"`
	Chat                          Chat                           `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	IsTopicMessage                *bool                          `json:"is_topic_message,omitempty"`
	IsAutomaticForward            *bool                          `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      *int                           `json:"edit_date,omitempty"`
	HasProtectedContent           *bool                          `json:"has_protected_content,omitempty"`
	IsFromOffline                 *bool                          `json:"is_from_offline,omitempty"`
	MediaGroupId                  *string                        `json:"media_group_id,omitempty"`
	AuthorSignature               *string                        `json:"author_signature,omitempty"`
	Text                          *string                        `json:"text,omitempty"`
	Entities                      *[]MessageEntity               `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	EffectId                      *string                        `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	Photo                         *[]PhotoSize                   `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       *string                        `json:"caption,omitempty"`
	CaptionEntities               *[]MessageEntity               `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         *bool                          `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                *[]User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  *[]PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`
	SuperGroupCreated             *bool                          `json:"super_group_created,omitempty"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               *int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             *int64                         `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

func (m MaybeInaccessibleMessage) IsAccessible() bool {
	return m.Date != 0
}

func (m MaybeInaccessibleMessage) MessageType() (*Message, error) {
	if !m.IsAccessible() {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "message",
			Got:      "inaccesible_message",
		}
	}
	msg := Message(m)
	return &msg, nil
}

func (m MaybeInaccessibleMessage) InaccesibleMessageType() (*InaccesibleMessage, error) {
	if m.IsAccessible() {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "inaccesible_message",
			Got:      "message",
		}
	}
	return &InaccesibleMessage{
		Chat:      m.Chat,
		MessageId: m.MessageId,
		Date:      m.Date,
	}, nil
}

type MessageEntity struct {
	//Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag or #hashtag@chatusername),
	//“cashtag” ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot), “url” (https://telegram.org),
	//“email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text),
	//“underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation),
	//“expandable_blockquote” (collapsed-by-default block quotation), “code” (monowidth string), “pre” (monowidth block),
	//“text_link” (for clickable text URLs), “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`
	//Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	//Length of the entity in UTF-16 code units
	Length int `json:"length"`
	//Optional. For “text_link” only, URL that will be opened after user taps on the text
	Url *string `json:"url,omitempty"`
	//Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`
	//Optional. For “pre” only, the programming language of the entity text
	Language *string `json:"language,omitempty"`
	//Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
}

// This object contains information about the quoted part of a message that is replied to by the given message.
type TextQuote struct {
	//Text of the quoted part of a message that is replied to by the given message
	Text string `json:"text"`
	//Optional. Special entities that appear in the quote.
	//Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Entities *[]MessageEntity `json:"entities,omitempty"`
	//Approximate quote position in the original message in UTF-16 code units as specified by the sender
	Position int `json:"position"`
	//Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
	IsManual *bool `json:"is_manual,omitempty"`
}

// This object contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	//Origin of the message replied to by the given message
	Origin MessageOrigin `json:"origin"`
	//Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	Chat *Chat `json:"chat,omitempty"`
	//Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	MessageId *int `json:"message_id,omitempty"`
	//Optional. Options used for link preview generation for the original message, if it is a text message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	//Optional. Message is an animation, information about the animation
	Animation *Animation `json:"animation,omitempty"`
	//Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`
	//Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`
	//Optional. Message contains paid media; information about the paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`
	//Optional. Message is a photo, available sizes of the photo
	Photo *[]PhotoSize `json:"photo,omitempty"`
	//Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	//Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`
	//Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`
	//Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`
	//Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`
	//Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler *bool `json:"has_media_spoiler,omitempty"`
	//Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`
	//Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`
	//Optional. Message is a game, information about the game.
	//More about games » https://core.telegram.org/bots/api#games
	Game *Game `json:"game,omitempty"`
	//Optional. Message is a scheduled giveaway, information about the giveaway
	Giveaway *Giveaway `json:"giveaway,omitempty"`
	//Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`
	//Optional. Message is an invoice for a payment, information about the invoice.
	//More about payments » https://core.telegram.org/bots/api#payments
	Invoice *Invoice `json:"invoice,omitempty"`
	//Optional. Message is a shared location, information about the location
	Location *Location `json:"location"`
	//Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`
	//Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`
}

// Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	//Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
	MessageId int `json:"message_id"`
	//Optional. If the message to be replied to is from a different chat,
	//unique identifier for the chat or username of the channel (in the format @channelusername).
	//Not supported for messages sent on behalf of a business account.
	ChatId *string `json:"chat_id,omitempty"`
	//Optional. Pass True if the message should be sent even if the specified message to be replied to is not found.
	//Always False for replies in another chat or forum topic. Always True for messages sent on behalf of a business account.
	AllowSendingWithoutReply *bool `json:"allow_sending_without_reply,omitempty"`
	//Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing.
	//The quote must be an exact substring of the message to be replied to, including bold, italic, underline,
	//strikethrough, spoiler, and custom_emoji entities.
	//The message will fail to send if the quote isn't found in the original message.
	Quote *string `json:"quote,omitempty"`
	//Optional. Mode for parsing entities in the quote.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	QuoteParseMode *string `json:"quote_parse_mode,omitempty"`
	//Optional. A JSON-serialized list of special entities that appear in the quote.
	//It can be specified instead of quote_parse_mode.
	QuoteEntities *[]MessageEntity `json:"quote_entities,omitempty"`
	//Optional. Position of the quote in the original message in UTF-16 code units
	QuotePosition *int `json:"quote_position,omitempty"`
}

func (r ReplyParameters) Validate() error {
	if r.ChatId != nil {
		if strings.TrimSpace(*r.ChatId) == "" {
			return ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if r.MessageId < 1 {
		return ErrInvalidParam("message_id parameter can't be empty")
	}
	if r.QuoteParseMode != nil && r.QuoteEntities != nil {
		return ErrInvalidParam("quote)parse_mode can't be used if quote_entities are provided")
	}
	return nil
}

// This object describes the origin of a message. It can be one of
//
// -MessageOriginUser
//
// -MessageOriginHiddenUser
//
// -MessageOriginChat
//
// -MessageOriginChannel
type MessageOrigin struct {
	Type            string  `json:"type"`
	Date            *int    `json:"date,omitempty"`
	SenderUser      *User   `json:"sender_user,omitempty"`
	SenderUserName  *string `json:"sender_user_name,omitempty"`
	SenderChat      *Chat   `json:"sender_chat,omitempty"`
	AuthorSignature *string `json:"author_signature,omitempty"`
	Chat            *Chat   `json:"chat,omitempty"`
	MessageId       *int    `json:"message_id,omitempty"`
}

func (m MessageOrigin) UserType() (*MessageOriginUser, error) {
	if m.Type != "user" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "user",
			Got:      m.Type,
		}
	}
	return &MessageOriginUser{
		Type:       m.Type,
		Date:       *m.Date,
		SenderUser: *m.SenderUser,
	}, nil
}

func (m MessageOrigin) HiddenUserType() (*MessageOriginHiddenUser, error) {
	if m.Type != "hidden_user" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "user",
			Got:      m.Type,
		}
	}
	return &MessageOriginHiddenUser{
		Type:           m.Type,
		Date:           *m.Date,
		SenderUsername: *m.SenderUserName,
	}, nil
}

func (m MessageOrigin) ChatType() (*MessageOriginChat, error) {
	if m.Type != "chat" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "chat",
			Got:      m.Type,
		}
	}
	return &MessageOriginChat{
		Type:            m.Type,
		Date:            *m.Date,
		SenderChat:      *m.SenderChat,
		AuthorSignature: m.AuthorSignature,
	}, nil
}

func (m MessageOrigin) ChannelType() (*MessageOriginChannel, error) {
	if m.Type != "channel" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "channel",
			Got:      m.Type,
		}
	}
	return &MessageOriginChannel{
		Type:            m.Type,
		Date:            *m.Date,
		Chat:            *m.Chat,
		MessageId:       *m.MessageId,
		AuthorSignature: m.AuthorSignature,
	}, nil
}

// The message was originally sent by a known user.
type MessageOriginUser struct {
	//Type of the message origin, always “user”
	Type string `json:"type"`
	//Date the message was sent originally in Unix time
	Date int `json:"date"`
	//User that sent the message originally
	SenderUser User `json:"sender_user"`
}

// The message was originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	//Type of the message origin, always “hidden_user”
	Type string `json:"type"`
	//Date the message was sent originally in Unix time
	Date int `json:"date"`
	//Name of the user that sent the message originally
	SenderUsername string `json:"sender_username"`
}

// The message was originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {
	//Type of the message origin, always “chat”
	Type string `json:"type"`
	//Date the message was sent originally in Unix time
	Date int `json:"date"`
	//Chat that sent the message originally
	SenderChat Chat `json:"sender_chat"`
	//Optional. For messages originally sent by an anonymous chat administrator, original message author signature
	AuthorSignature *string `json:"author_signature,omitempty"`
}

// The message was originally sent to a channel chat.
type MessageOriginChannel struct {
	//Type of the message origin, always “channel”
	Type string `json:"type"`
	//Date the message was sent originally in Unix time
	Date int `json:"date"`
	//Channel chat to which the message was originally sent
	Chat Chat `json:"chat"`
	//Unique message identifier inside the chat
	MessageId int `json:"message_id"`
	//Optional. Signature of the original post author
	AuthorSignature *string `json:"author_signature,omitempty"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Photo width
	Width int `json:"width"`
	//Photo height
	Height int `json:"height"`
	//Optional. File size in bytes
	FileSize *int `json:"file_size,omitempty"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Video width as defined by the sender
	Width int `json:"width"`
	//Video height as defined by the sender
	Height int `json:"height"`
	//Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`
	//Optional. Animation thumbnail as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	//Optional. Original animation filename as defined by the sender
	FileName *string `json:"file_name,omitempty"`
	//Optional. MIME type of the file as defined by the sender
	MimeType *string `json:"mime_type,omitempty"`
	//Optional. File size in bytes. It can be bigger than 2^31 and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Duration of the audio in seconds as defined by the sender
	Duration int `json:"duration"`
	//Optional. Performer of the audio as defined by the sender or by audio tags
	Performer *string `json:"performer,omitempty"`
	//Optional. Title of the audio as defined by the sender or by audio tags
	Title *string `json:"title,omitempty"`
	//Optional. Original filename as defined by the sender
	FileName *string `json:"file_name,omitempty"`
	//Optional. MIME type of the file as defined by the sender
	MimeType *string `json:"mime_type,omitempty"`
	//Optional. File size in bytes. It can be bigger than 2^31 and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int64 `json:"file_size,omitempty"`
	//Optional. Thumbnail of the album cover to which the music file belongs
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Optional. Document thumbnail as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	//Optional. Original filename as defined by the sender
	FileName *string `json:"file_name,omitempty"`
	//Optional. MIME type of the file as defined by the sender
	MimeType *string `json:"mime_type,omitempty"`
	//Optional. File size in bytes. It can be bigger than 2^31 and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a story.
type Story struct {
	//Chat that posted the story
	Chat Chat `json:"chat"`
	//Unique identifier for the story in the chat
	Id int `json:"id"`
}

// This object represents a video file.
type Video struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Video width as defined by the sender
	Width int `json:"width"`
	//Video height as defined by the sender
	Height int `json:"height"`
	//Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`
	//Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	//Optional. Original filename as defined by the sender
	FileName *string `json:"file_name,omitempty"`
	//Optional. MIME type of the file as defined by the sender
	MimeType *string `json:"mime_type,omitempty"`
	//Optional. File size in bytes. It can be bigger than 2^31 and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Video width and height (diameter of the video message) as defined by the sender
	Length int `json:"length"`
	//Duration of the video in seconds as defined by the sender
	Duration int `json:"duration"`
	//Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	//Optional. File size in bytes
	FileSize *int `json:"file_size,omitempty"`
}

// This object represents a voice note.
type Voice struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Duration of the audio in seconds as defined by the sender
	Duration int `json:"duration"`
	//Optional. MIME type of the file as defined by the sender
	MimeType *string `json:"mime_type,omitempty"`
	//Optional. File size in bytes. It can be bigger than 2^31 and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int `json:"file_size,omitempty"`
}

// Describes the paid media added to a message.
type PaidMediaInfo struct {
	//The number of Telegram Stars that must be paid to buy access to the media
	StarCount string `json:"star_count"`
	//Information about the paid media
	PaidMedia []PaidMedia `json:"paid_media"`
}

// This object describes paid media. Currently, it can be one of
//
// - PaidMediaPreview
//
// - PaidMediaPhoto
//
// - PaidMediaVideo
type PaidMedia struct {
	Type     string       `json:"type"`
	Width    *int         `json:"width,omitempty"`
	Height   *int         `json:"height,omitempty"`
	Duration *int         `json:"duration,omitempty"`
	Photo    *[]PhotoSize `json:"photo,omitempty"`
	Video    *Video       `json:"video,omitempty"`
}

func (p PaidMedia) PreviewType() (*PaidMediaPreview, error) {
	if p.Type != "preview" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "preview",
			Got:      p.Type,
		}
	}
	return &PaidMediaPreview{
		Type:     p.Type,
		Width:    p.Width,
		Height:   p.Height,
		Duration: p.Duration,
	}, nil
}

func (p PaidMedia) PhotoType() (*PaidMediaPhoto, error) {
	if p.Type != "photo" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "photo",
			Got:      p.Type,
		}
	}
	return &PaidMediaPhoto{
		Type:  p.Type,
		Photo: *p.Photo,
	}, nil
}

func (p PaidMedia) VideoType() (*PaidMediaVideo, error) {
	if p.Type != "video" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "video",
			Got:      p.Type,
		}
	}
	return &PaidMediaVideo{
		Type:  p.Type,
		Video: p.Video,
	}, nil
}

// The paid media isn't available before the payment.
type PaidMediaPreview struct {
	//Type of the paid media, always “preview”
	Type string `json:"type"`
	//Optional. Media width as defined by the sender
	Width *int `json:"width,omitempty"`
	//Optional. Media height as defined by the sender
	Height *int `json:"height,omitempty"`
	//Optional. Duration of the media in seconds as defined by the sender
	Duration *int `json:"duration,omitempty"`
}

// The paid media is a photo.
type PaidMediaPhoto struct {
	//Type of the paid media, always “photo”
	Type string `json:"type"`
	//The photo
	Photo []PhotoSize `json:"photo"`
}

// The paid media is a video.
type PaidMediaVideo struct {
	//Type of the paid media, always “video”
	Type string `json:"type"`
	//The video
	Video *Video `json:"video"`
}

// This object represents a phone contact.
type Contact struct {
	//Contact's phone number
	PhoneNumber string `json:"phone_number"`
	//Contact's first name
	FirstName string `json:"first_name"`
	//Optional. Contact's last name
	LastName *string `json:"last_name,omitempty"`
	//Optional. Contact's user identifier in Telegram.
	//This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserId *int64 `json:"user_id,omitempty"`
	//Optional. Additional data about the contact in the form of a vCard
	VCard *string `json:"v_card,omitempty"`
}

// This object represents an animated emoji that displays a random value.
type Dice struct {
	//Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value int `json:"value"`
	//Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`
}

// This object contains information about one answer option in a poll.
type PollOption struct {
	//Option text, 1-100 characters
	Text string `json:"text"`
	//Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	VoterCount int `json:"voter_count"`
	//Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
}

// This object contains information about one answer option in a poll to be sent.
type InputPollOption struct {
	//Option text, 1-100 characters
	Text string `json:"text"`
	//Optional. Mode for parsing entities in the text. See https://core.telegram.org/bots/api#formatting-options for more details.
	//Currently, only custom emoji entities are allowed
	TextParseMode *string `json:"text_parse_mode,omitempty"`
	//Optional. A JSON-serialized list of special entities that appear in the poll option text.
	//It can be specified instead of text_parse_mode
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
}

func (i InputPollOption) Validate() error {
	if len(i.Text) < 1 || len(i.Text) > 100 {
		return fmt.Errorf("text must be between 1 and 100 characters")
	}
	if i.TextParseMode != nil && i.TextEntities != nil {
		return fmt.Errorf("parse_mode parameter can't be used if entities are provided")
	}
	return nil
}

// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	//Unique poll identifier
	PollId string `json:"poll_id"`
	//Optional. The chat that changed the answer to the poll, if the voter is anonymous
	VoterChat *Chat `json:"voter_chat"`
	//Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	User *User `json:"user"`
	//0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIds []int `json:"option_ids"`
}

// This object contains information about a poll.
type Poll struct {
	//Unique poll identifier
	Id string `json:"id"`
	//Poll question, 1-300 characters
	Question string `json:"question"`
	//Optional. Special entities that appear in the question.
	//Currently, only custom emoji entities are allowed in poll questions
	QuestionEntities []MessageEntity `json:"question_entities"`
	// /List of poll options
	Options []PollOption `json:"options"`
	//Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`
	//True, if the poll is closed
	IsClosed bool `json:"is_closed"`
	//True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`
	//Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`
	//True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	//Optional. 0-based identifier of the correct answer option.
	//Available only for polls in the quiz mode, which are closed, or
	//was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionId *int `json:"correct_option_id,omitempty"`
	//Optional. Text that is shown when a user chooses an incorrect answer or
	//taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation *string `json:"explanation,omitempty"`
	//Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	ExplanationEntities *[]MessageEntity `json:"explanation_entities,omitempty"`
	//Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod *int `json:"open_period,omitempty"`
	//Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate *int `json:"close_date,omitempty"`
}

// This object represents a point on the map.
type Location struct {
	// 	Latitude as defined by the sender
	Latitude float64 `json:"latitude"`
	//Longitude as defined by the sender
	Longitude float64 `json:"longitude"`
	//Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`
	//Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	LivePeriod *int `json:"live_period,omitempty"`
	//Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	Heading *int `json:"heading,omitempty"`
	//Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
	ProximityAlertRadius *int `json:"proximity_alert_radius,omitempty"`
}

// This object represents a venue.
type Venue struct {
	//Venue location. Can't be a live location
	Location Location `json:"location"`
	//Name of the venue
	Title string `json:"title"`
	//Address of the venue
	Address string `json:"address"`
	//Optional. Foursquare identifier of the venue
	FoursquareId *string `json:"foursquare_id,omitempty"`
	//Optional. Foursquare type of the venue.
	//(For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FourSquareType *string `json:"four_square_type,omitempty"`
	//Optional. Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	//Optional. Google Places type of the venue. (See supported types: https://developers.google.com/places/web-service/supported_types)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

// Describes data sent from a Web App to the bot.
type WebAppData struct {
	//The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`
	//Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// This object represents the content of a service message,
// sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	//User that triggered the alert
	Traveler User `json:"traveler"`
	//User that set the alert
	Watcher User `json:"watcher"`
	//The distance between the users
	Distance int `json:"distance"`
}

// This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	//New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// This object represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

// This object describes the way a background is filled based on the selected colors. Currently, it can be one of
//
// - BackgroundFillSolid
//
// - BackgroundFillGradient
//
// - BackgroundFillFreeformGradient
type BackgroundFill struct {
	Type          string `json:"type"`
	Color         *int   `json:"color,omitempty"`
	TopColor      *int   `json:"top_color,omitempty"`
	BottomColor   *int   `json:"bottom_color,omitempty"`
	RotationAngle *int   `json:"rotation_angle,omitempty"`
	Colors        *[]int `json:"colors,omitempty"`
}

func (b BackgroundFill) SolidType() (*BackgroundFillSolid, error) {
	if b.Type != "solid" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "solid",
			Got:      b.Type,
		}
	}
	return &BackgroundFillSolid{
		Type:  b.Type,
		Color: *b.Color,
	}, nil
}

func (b BackgroundFill) GradientType() (*BackgroundFillGradient, error) {
	if b.Type != "gradient" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "gradient",
			Got:      b.Type,
		}
	}
	return &BackgroundFillGradient{
		Type:          b.Type,
		TopColor:      *b.TopColor,
		BottomColor:   *b.BottomColor,
		RotationAngle: *b.RotationAngle,
	}, nil
}

func (b BackgroundFill) FreeformGradientType() (*BackgroundFillFreeformGradient, error) {
	if b.Type != "freeform_gradient" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "freeform_gradient",
			Got:      b.Type,
		}
	}
	return &BackgroundFillFreeformGradient{
		Type:   b.Type,
		Colors: *b.Colors,
	}, nil
}

// The background is filled using the selected color.
type BackgroundFillSolid struct {
	//Type of the background fill, always “solid”
	Type string `json:"type"`
	//The color of the background fill in the RGB24 format
	Color int `json:"color"`
}

// The background is a gradient fill.
type BackgroundFillGradient struct {
	//Type of the background fill, always “gradient”
	Type string `json:"type"`
	//Top color of the gradient in the RGB24 format
	TopColor int `json:"top_color"`
	//Bottom color of the gradient in the RGB24 format
	BottomColor int `json:"bottom_color"`
	//Clockwise rotation angle of the background fill in degrees; 0-359
	RotationAngle int `json:"rotation_angle"`
}

// The background is a freeform gradient that rotates after every message in the chat.
type BackgroundFillFreeformGradient struct {
	//Type of the background fill, always “freeform_gradient”
	Type string `json:"type"`
	//A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
	Colors []int `json:"colors"`
}

// This object describes the type of a background. Currently, it can be one of
//
// - BackgroundTypeFill
//
// - BackgroundTypeWallpaper
//
// - BackgroundTypePattern
//
// - BackgroundTypeChatTheme
type BackgroundType struct {
	Type             string          `json:"type"`
	Fill             *BackgroundFill `json:"fill,omitempty"`
	DarkThemeDimming *int            `json:"dark_theme_dimming,omitempty"`
	Document         *Document       `json:"document,omitempty"`
	IsBlurred        *bool           `json:"is_blurred,omitempty"`
	IsMoving         *bool           `json:"is_moving,omitempty"`
	Intensity        *int            `json:"intensity,omitempty"`
	IsInverted       *bool           `json:"is_inverted,omitempty"`
	ThemeName        *string         `json:"theme_name,omitempty"`
}

func (b BackgroundType) FillType() (*BackgroundTypeFill, error) {
	if b.Type != "fill" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "fill",
			Got:      b.Type,
		}
	}
	return &BackgroundTypeFill{
		Type:             b.Type,
		Fill:             *b.Fill,
		DarkThemeDimming: *b.DarkThemeDimming,
	}, nil
}

func (b BackgroundType) WallpaperType() (*BackgroundTypeWallpaper, error) {
	if b.Type != "wallpaper" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "wallpaper",
			Got:      b.Type,
		}
	}
	return &BackgroundTypeWallpaper{
		Type:             b.Type,
		Document:         *b.Document,
		DarkThemeDimming: *b.DarkThemeDimming,
		IsBlurred:        b.IsBlurred,
		IsMoving:         b.IsMoving,
	}, nil
}

func (b BackgroundType) PatternType() (*BackgroundTypePattern, error) {
	if b.Type != "pattern" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "pattern",
			Got:      b.Type,
		}
	}
	return &BackgroundTypePattern{
		Type:       b.Type,
		Document:   *b.Document,
		Fill:       *b.Fill,
		Intensity:  *b.Intensity,
		IsInverted: b.IsInverted,
		IsMoving:   b.IsMoving,
	}, nil
}

func (b BackgroundType) ChatThemeType() (*BackgroundTypeChatTheme, error) {
	if b.Type != "chat_theme" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "chat_theme",
			Got:      b.Type,
		}
	}
	return &BackgroundTypeChatTheme{
		Type:      b.Type,
		ThemeName: *b.ThemeName,
	}, nil
}

// The background is automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	//Type of the background, always “fill”
	Type string `json:"type"`
	//The background fill
	Fill BackgroundFill `json:"fill"`
	//Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
}

// The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	//Type of the background, always “wallpaper”
	Type string `json:"type"`
	//Document with the wallpaper
	Document Document `json:"document"`
	//Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
	//Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsBlurred *bool `json:"is_blurred,omitempty"`
	//Optional. True, if the background moves slightly when the device is tilted
	IsMoving *bool `json:"is_moving,omitempty"`
}

// The background is a PNG or TGV (gzipped subset of SVG with MIME type “application/x-tgwallpattern”)
// pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	//Type of the background, always “pattern”
	Type string `json:"type"`
	//Document with the pattern
	Document Document `json:"document"`
	//The background fill that is combined with the pattern
	Fill BackgroundFill `json:"fill"`
	//Intensity of the pattern when it is shown above the filled background; 0-100
	Intensity int `json:"intensity"`
	//Optional. True, if the background fill must be applied only to the pattern itself.
	//All other pixels are black in this case. For dark themes only
	IsInverted *bool `json:"is_inverted,omitempty"`
	//Optional. True, if the background moves slightly when the device is tilted
	IsMoving *bool `json:"is_moving,omitempty"`
}

// The background is taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	//Type of the background, always “chat_theme”
	Type string `json:"type"`
	//Name of the chat theme, which is usually an emoji
	ThemeName string `json:"theme_name"`
}

// This object represents a chat background.
type ChatBackground struct {
	//Type of the background
	Type BackgroundType `json:"type"`
}

// This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	//Name of the topic
	Name string `json:"name"`
	//Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`
	//Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

// This object represents a service message about a forum topic closed in the chat.
// Currently holds no information.
type ForumTopicClosed struct {
}

// This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	//Optional. New name of the topic, if it was edited
	Name *string `json:"name,omitempty"`
	//Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

// This object represents a service message about a forum topic reopened in the chat.
// Currently holds no information.
type ForumTopicReopened struct {
}

// This object represents a service message about General forum topic hidden in the chat.
// Currently holds no information.
type GeneralForumTopicHidden struct {
}

// This object represents a service message about General forum topic unhidden in the chat.
// Currently holds no information.
type GeneralForumTopicUnhidden struct {
}

// This object contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
type SharedUser struct {
	//Identifier of the shared user. This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers.
	//The bot may not have access to the user and could be unable to use this identifier,
	//unless the user is already known to the bot by some other means.
	UserId int64 `json:"user_id"`
	//Optional. First name of the user, if the name was requested by the bot
	FirstName *string `json:"first_name,omitempty"`
	//Optional. Last name of the user, if the name was requested by the bot
	LastName *string `json:"last_name,omitempty"`
	//Optional. Username of the user, if the username was requested by the bot
	Username *string `json:"username,omitempty"`
	//Optional. Available sizes of the chat photo, if the photo was requested by the bot
	Photo *[]PhotoSize `json:"photo,omitempty"`
}

// This object contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	//Identifier of the request
	RequestId string `json:"request_id"`
	//Information about users shared with the bot.
	Users []SharedUser `json:"users"`
}

// This object contains information about a chat that was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	//Identifier of the request
	RequestId string `json:"request_id"`
	//Identifier of the shared chat. This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	//The bot may not have access to the chat and could be unable to use this identifier,
	//unless the chat is already known to the bot by some other means.
	ChatId int64 `json:"chat_id"`
	//Optional. Title of the chat, if the title was requested by the bot.
	Title *string `json:"title,omitempty"`
	//Optional. Username of the chat, if the username was requested by the bot and available.
	Username *string `json:"username,omitempty"`
	//Optional. Available sizes of the chat photo, if the photo was requested by the bot
	Photo *[]PhotoSize `json:"photo,omitempty"`
}

// This object represents a service message about a user allowing a bot to write messages after
// adding it to the attachment menu, launching a Web App from a link,
// or accepting an explicit request from a Web App sent by the method requestWriteAccess.
type WriteAccessAllowed struct {
	//Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	FromRequest *bool `json:"from_request,omitempty"`
	//Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	WebAppName *string `json:"web_app_name,omitempty"`
	//Optional. True, if the access was granted when the bot was added to the attachment or side menu
	FromAttachmentMenu *bool `json:"from_attachment_menu,omitempty"`
}

// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	//Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
	StartDate int `json:"start_date"`
}

// This object represents a service message about a video chat started in the chat.
// Currently holds no information.
type VideoChatStarted struct {
}

// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	//Video chat duration in seconds
	Duration int `json:"duration"`
}

// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	//New members that were invited to the video chat
	Users []User `json:"users"`
}

// This object represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
}

// This object represents a message about a scheduled giveaway.
type Giveaway struct {
	//The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`
	//Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerSelectionDate int `json:"winner_selection_date"`
	//The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int `json:"winner_count"`
	//Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	OnlyNewMembers *bool `json:"only_new_members,omitempty"`
	//Optional. True, if the list of giveaway winners will be visible to everyone
	HasPublicWinners *bool `json:"has_public_winners,omitempty"`
	// /Optional. Description of additional giveaway prize
	PrizeDescription *string `json:"prize_description,omitempty"`
	//Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating
	//the countries from which eligible users for the giveaway must come.
	//If empty, then all users can participate in the giveaway.
	//Users with a phone number that was bought on Fragment can always participate in giveaways.
	CountryCodes *[]string `json:"country_codes,omitempty"`
	//Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
	//Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for;
	//for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount *int `json:"premium_subscription_month_count,omitempty"`
}

// This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	//The chat that created the giveaway
	Chat Chat `json:"chat"`
	//Identifier of the message with the giveaway in the chat
	GiveawayMessageId int `json:"giveaway_message_id"`
	//Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int `json:"winners_selection_date"`
	//Total number of winners in the giveaway
	WinnerCount int `json:"winner_count"`
	//List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`
	//Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	AdditionalChatCount *int `json:"additional_chat_count,omitempty"`
	//Optional. The number of Telegram Stars that were split between giveaway winners;
	//for Telegram Star giveaways only
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
	//Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for;
	//for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount *int `json:"premium_subscription_month_count,omitempty"`
	//Optional. Number of undistributed prizes
	UnclaimedPrizeCount *int `json:"unclaimed_prize_count,omitempty"`
	//Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	OnlyNewMembers *bool `json:"only_new_members,omitempty"`
	//Optional. True, if the giveaway was canceled because the payment for it was refunded
	WasRefunded *bool `json:"was_refunded,omitempty"`
	//Optional. Description of additional giveaway prize
	PrizeDescription *string `json:"prize_description,omitempty"`
}

// This object represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	//Number of winners in the giveaway
	WinnerCount int `json:"winner_count"`
	//Optional. Number of undistributed prizes
	UnclaimedPrizeCount *int `json:"unclaimed_prize_count,omitempty"`
	//Optional. Message with the giveaway that was completed, if it wasn't deleted
	GiveawayMessage *Message `json:"giveaway_message,omitempty"`
	//Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
	IsStarGiveaway *bool `json:"is_star_giveaway,omitempty"`
}

// Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	//Optional. True, if the link preview is disabled
	IsDisabled *bool `json:"is_disabled,omitempty"`
	//Optional. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	UrlFileId *string `json:"url_file_id,omitempty"`
	//Optional. True, if the media in the link preview is supposed to be shrunk;
	//ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferSmallMedia *bool `json:"prefer_small_media,omitempty"`
	//Optional. True, if the media in the link preview is supposed to be enlarged;
	//ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia *bool `json:"prefer_large_media,omitempty"`
	//Optional. True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
	ShowAboveText *bool `json:"show_above_text,omitempty"`
}

func (l LinkPreviewOptions) Validate() error {
	if *l.PreferLargeMedia && *l.PreferSmallMedia {
		return fmt.Errorf("PreferSmallMedia and PreferLargeMedia parameters are mutual exclusive")
	}
	return nil
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	//Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	//Requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}

// This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
type File struct {
	//Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	//Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	//Optional. File size in bytes. It can be bigger than 2^31 and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int64 `json:"file_size,omitempty"`
	//Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath *string `json:"file_path,omitempty"`
}

func (f File) GetFileUrl(botToken string) (string, error) {
	if botToken == "" || strings.TrimSpace(botToken) == "" {
		return "", fmt.Errorf("bot token can't be empty")
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, *f.FilePath), nil
}

// Describes a Web App.
type WebAppInfo struct {
	//An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	Url string `json:"url"`
}

func (w WebAppInfo) Validate() error {
	if strings.TrimSpace(w.Url) == "" {
		return ErrInvalidParam("url paramter can't be empty")
	}
	return nil
}

// The ReplyMarkup interface unifies the possible types of reply markup objects that
// can be used in Telegram Bot API methods for customizing keyboards in messages.
type ReplyMarkup struct {
	ReplyMarkupInterface
}

type ReplyMarkupInterface interface {
	replyKeyboardContract()
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
// Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardMarkup struct {
	//Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]KeyboardButton `json:"keyboard"`
	//Optional. Requests clients to always show the keyboard when the regular keyboard is hidden.
	//Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	IsPersistent *bool `json:"is_persistent,omitempty"`
	//Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons).
	//Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard *bool `json:"resize_keyboard,omitempty"`
	//Optional. Requests clients to hide the keyboard as soon as it's been used.
	//The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat -
	//the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard *bool `json:"one_time_keyboard,omitempty"`
	//Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	//Optional. Use this parameter if you want to show the keyboard to specific users only. Targets:
	//
	//1) users that are @mentioned in the text of the Message object;
	//
	//2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
	//
	//Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
	Selective *bool `json:"selective,omitempty"`
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

// This object represents one button of the reply keyboard. At most one of the optional fields must be used to specify type of the button.
// For simple text buttons, String can be used instead of this object to specify the button text.
type KeyboardButton struct {
	//Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	Text string `json:"text"`
	//Optional. If specified, pressing the button will open a list of suitable users.
	//Identifiers of selected users will be sent to the bot in a “users_shared” service message.
	//Available in private chats only.
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	//Optional. If specified, pressing the button will open a list of suitable chats.
	//Tapping on a chat will send its identifier to the bot in a “chat_shared” service message.
	//Available in private chats only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	//Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
	//Available in private chats only.
	RequestContact *bool `json:"request_contact,omitempty"`
	//Optional. If True, the user's current location will be sent when the button is pressed.
	//Available in private chats only.
	RequestLocation *bool `json:"request_location,omitempty"`
	//Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	//Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	//Optional. If specified, the described Web App will be launched when the button is pressed.
	//The Web App will be able to send a “web_app_data” service message.
	//Available in private chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

func (k KeyboardButton) Validate() error {
	if strings.TrimSpace(k.Text) == "" {
		return ErrInvalidParam("text paramter can't be empty")
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
		return fmt.Errorf("at most one of the optional fields must be used to specify type of the button")
	}

	return nil
}

// This object defines the criteria used to request suitable users.
// Information about the selected users will be shared with the bot when the corresponding button is pressed.
// More about requesting users » https://core.telegram.org/bots/features#chat-and-user-selection
type KeyboardButtonRequestUsers struct {
	//Signed 32-bit identifier of the request that will be received back in the UsersShared object.
	//Must be unique within the message
	RequestId int32 `json:"request_id"`
	//Optional. Pass True to request bots, pass False to request regular users.
	//If not specified, no additional restrictions are applied.
	UserIsBot *bool `json:"user_is_bot,omitempty"`
	//Optional. Pass True to request premium users, pass False to request non-premium users.
	//If not specified, no additional restrictions are applied.
	UserIsPremium *bool `json:"user_is_premium,omitempty"`
	//Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity *int `json:"max_quantity,omitempty"`
	//Optional. Pass True to request the users' first and last names
	RequestName *bool `json:"request_name,omitempty"`
	//Optional. Pass True to request the users' usernames
	RequestUsername *bool `json:"request_username,omitempty"`
	//Optional. Pass True to request the users' photos
	RequestPhoto *bool `json:"request_photo,omitempty"`
}

func (k KeyboardButtonRequestUsers) Validate() error {
	if k.RequestId == 0 {
		return ErrInvalidParam("request_id paramter can't be empty")
	}
	if *k.MaxQuantity < 1 || *k.MaxQuantity > 10 {
		return fmt.Errorf("MaxQuantity parameter must be between 1 and 10")
	}
	return nil
}

// This object defines the criteria used to request a suitable chat.
// Information about the selected chat will be shared with the bot when the corresponding button is pressed.
// The bot will be granted requested rights in the chat if appropriate.
// More about requesting chats » https://core.telegram.org/bots/features#chat-and-user-selection
type KeyboardButtonRequestChat struct {
	RequestId int32 `json:"request_id"`
	//Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`
	//Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat.
	//If not specified, no additional restrictions are applied.
	ChatIsForum *bool `json:"chat_is_forum,omitempty"`
	//Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username.
	//If not specified, no additional restrictions are applied.
	ChatHasUsername *bool `json:"chat_has_username,omitempty"`
	//Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	ChatIsCreated *bool `json:"chat_is_created,omitempty"`
	//Optional. A JSON-serialized object listing the required administrator rights of the user in the chat.
	//The rights must be a superset of bot_administrator_rights.
	//If not specified, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	//Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat.
	//The rights must be a subset of user_administrator_rights.
	//If not specified, no additional restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	//Optional. Pass True to request a chat with the bot as a member.
	//Otherwise, no additional restrictions are applied.
	BotIsMember *bool `json:"bot_is_member,omitempty"`
	//Optional. Pass True to request the chat's title
	RequestTitle *bool `json:"request_title,omitempty"`
	//Optional. Pass True to request the chat's username
	RequestUsername *bool `json:"request_username,omitempty"`
	//Optional. Pass True to request the chat's photo
	RequestPhoto *bool `json:"request_photo,omitempty"`
}

func (k KeyboardButtonRequestChat) Validate() error {
	if k.RequestId == 0 {
		return ErrInvalidParam("request_id paramter can't be empty")
	}
	return nil
}

// This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	//Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	//If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
	Type *string `json:"type,omitempty"`
}

func (k KeyboardButtonPollType) Validate() error {
	if k.Type != nil {
		if *k.Type != "regular" && *k.Type != "quiz" {
			return ErrInvalidParam("type must be reqular or quiz if specified")
		}
	}
	return nil
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and
// display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
// Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardRemove struct {
	//Requests clients to remove the custom keyboard (user will not be able to summon this keyboard;
	//if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`
	//Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets:
	//
	//1) users that are @mentioned in the text of the Message object;
	//
	//2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
	//
	//Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user,
	//while still showing the keyboard with poll options to users who haven't voted yet.
	Selective *bool `json:"selective,omitempty"`
}

func (f ReplyKeyboardRemove) replyKeyboardContract() {}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	//Array of button rows, each represented by an Array of InlineKeyboardButton objects
	Keyboard [][]InlineKeyboardButton `json:"keyboard"`
}

func (f InlineKeyboardMarkup) replyKeyboardContract() {}

func (m InlineKeyboardMarkup) Validate() error {
	for i, row := range m.Keyboard {
		for j, key := range row {
			if err := key.Validate(); err != nil {
				return err
			}
			if key.Pay != nil {
				if i == 0 || j == 0 {
					return ErrInvalidParam("the button with a specified pay parameter must always be the first button at the first row")
				}
			}
			if key.CallbackGame != nil {
				if i == 0 || j == 0 {
					return ErrInvalidParam("the button with a specified callback_game parameter must always be the first button at the first row")
				}
			}
		}
	}
	return nil
}

// This object represents one button of an inline keyboard. Exactly one of the optional fields must be used to specify type of the button.
type InlineKeyboardButton struct {
	//Label text on the button
	Text string `json:"text"`
	//Optional. HTTP or tg:// URL to be opened when the button is pressed.
	//Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username,
	//if this is allowed by their privacy settings.
	Url *string `json:"url,omitempty"`
	//Optional. Data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes
	CallbackData *string `json:"callback_data,omitempty"`
	//Optional. Description of the Web App that will be launched when the user presses the button.
	//The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	//Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	//Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	LoginUrl *LoginUrl `json:"login_url,omitempty"`
	//Optional. If set, pressing the button will prompt the user to select one of their chats,
	//open that chat and insert the bot's username and the specified inline query in the input field.
	//May be empty, in which case just the bot's username will be inserted.
	//Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`
	//Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field.
	//May be empty, in which case only the bot's username will be inserted.
	//
	//This offers a quick way for the user to open your bot in inline mode in the same chat -
	//good for selecting something from multiple options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`
	//Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type,
	//open that chat and insert the bot's username and the specified inline query in the input field.
	//Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	//Optional. Description of the button that copies the specified text to the clipboard.
	CopyText *CopyTextButton `json:"copy_text,omitempty"`
	// 	Optional. Description of the game that will be launched when the user presses the button.
	//
	//IMPORTANT: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	//Optional. Specify True, to send a Pay button. Substrings “⭐” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.
	//
	//IMPORTANT: This type of button must always be the first button in the first row and can only be used in invoice messages.
	Pay *bool `json:"pay,omitempty"`
}

func (b InlineKeyboardButton) Validate() error {
	if strings.TrimSpace(b.Text) == "" {
		return ErrInvalidParam("text paramter can't be empty")
	}
	if b.CallbackData != nil {
		if len([]byte(*b.CallbackData)) > 64 {
			return ErrInvalidParam("callback_data must not be longer than 64 bytes if specified")
		}
	}
	if b.CopyText != nil {
		if err := b.CopyText.Validate(); err != nil {
			return err
		}
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

// This object represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in.
//
// Telegram apps support these buttons as of version 5.7.
type LoginUrl struct {
	//An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed.
	//If the user refuses to provide authorization data, the original URL without information about the user will be opened.
	//The data added is the same as described in Receiving authorization data.
	//
	//IMPORTANT: You must always check the hash of the received data to verify the authentication and
	//the integrity of the data as described in https://core.telegram.org/widgets/login#checking-authorization.
	Url string `json:"url"` // FIXME: should somehow check the hash of the received data
	//Optional. New text of the button in forwarded messages.
	ForwardText *string `json:"forward_text,omitempty"`
	//Optional. Username of a bot, which will be used for user authorization.
	//See Setting up a bot for more details. If not specified, the current bot's username will be assumed.
	//The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	BotUsername *string `json:"bot_username,omitempty"`
	//Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAcess *bool `json:"request_write_acess,omitempty"`
}

func (l LoginUrl) Validate() error {
	if strings.TrimSpace(l.Url) == "" {
		return ErrInvalidParam("url paramter can't be empty")
	}
	return nil
}

// This object represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	//Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	Query *string `json:"query,omitempty"`
	//Optional. True, if private chats with users can be chosen
	AllowUserChats *bool `json:"allow_user_chats,omitempty"`
	// /Optional. True, if private chats with bots can be chosen
	AllowBotChats *bool `json:"allow_bot_chats,omitempty"`
	//Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats *bool `json:"allow_group_chats,omitempty"`
	//Optional. True, if channel chats can be chosen
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty"`
}

// This object represents an inline keyboard button that copies specified text to the clipboard.
type CopyTextButton struct {
	//The text to be copied to the clipboard; 1-256 characters
	Text string
}

func (c CopyTextButton) Validate() error {
	if len(c.Text) < 1 || len(c.Text) > 256 {
		return ErrInvalidParam("text parameter must be between 1 and 256 characters")
	}
	return nil
}

// This object represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
//
// IMPORTANT: After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery.
// It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed
// (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	//Unique identifier for this query
	Id string `json:"id"`
	//Sender
	From User `json:"from"`
	//Optional. Message sent by the bot with the callback button that originated the query
	Message *MaybeInaccessibleMessage `json:"message,omitempty"`
	//Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	//Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent.
	//Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`
	//Optional. Data associated with the callback button.
	//Be aware that the message originated the query can contain no callback buttons with this data.
	Data *string `json:"data,omitempty"`
	//Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName *string `json:"game_short_name,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user
// (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
// Not supported in channels and for messages sent on behalf of a Telegram Business account.
//
// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages and mentions).
// There could be two ways to create a new poll:
//
// - Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2).
// May be appealing for hardcore users but lacks modern day polish.
//
// - Guide the user through a step-by-step process. 'Please send me your question', 'Cool, now let's add the first answer option',
// 'Great. Keep adding answer options, then send /done when you're ready'.
//
// The last option is definitely more attractive. And if you use ForceReply in your bot's questions,
// it will receive the user's answers even if it only receives replies, commands and mentions - without any extra work for the user.
type ForceReply struct {
	//Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`
	//Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	//Optional. Use this parameter if you want to force reply from specific users only. Targets:
	//
	//1) users that are @mentioned in the text of the Message object;
	//
	//2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
	Selective *bool `json:"selective,omitempty"`
}

func (f ForceReply) replyKeyboardContract() {}

func (f ForceReply) Validate() error {
	if f.InputFieldPlaceholder != nil {
		if len(*f.InputFieldPlaceholder) < 1 || len(*f.InputFieldPlaceholder) > 64 {
			return fmt.Errorf("InputFieldPlaceholder parameter must be between 1 and 64 characters")
		}
	}
	return nil
}

// This object represents a chat photo.
type ChatPhoto struct {
	//File identifier of small (160x160) chat photo.
	//This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`
	//Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`
	//File identifier of big (640x640) chat photo.
	//This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`
	//Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	BigFileUniqueId string `json:"big_file_unique_id"`
}

// Represents an invite link for a chat.
type ChatInviteLink struct {
	//The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`
	//Creator of the link
	Creator User `json:"creator"`
	//True, if users joining the chat via the link need to be approved by chat administrators
	CreatesJoinRequest bool `json:"creates_join_request"`
	//True, if the link is primary
	IsPrimary bool `json:"is_primary"`
	//True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`
	//Optional. Invite link name
	Name *string `json:"name,omitempty"`
	//Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	ExpireDate *int `json:"expire_date,omitempty"`
	//Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *bool `json:"member_limit,omitempty"`
	//Optional. Number of pending join requests created using this link
	PendingJoinRequestCount *int `json:"pending_join_request_count,omitempty"`
	//Optional. The number of seconds the subscription will be active for before the next payment
	SubscriptionPeriod *int `json:"subscription_period,omitempty"`
	//Optional. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
	SubscriptionPrice *int `json:"subscription_price,omitempty"`
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	//True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
	//True, if the administrator can access the chat event log, get boost list,
	//see hidden supergroup and channel members, report spam messages and ignore slow mode.
	//Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`
	//True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`
	//True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	//True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members"`
	//True, if the administrator can add new administrators with a subset of their own privileges or
	//demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`
	//True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	//True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	//True, if the administrator can post stories to the chat
	CanPostStories bool `json:"can_post_stories"`
	//True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories bool `json:"can_edit_stories"`
	//True, if the administrator can delete stories posted by other users
	CanDeleteStories bool `json:"can_delete_stories"`
	//Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`
	//Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`
	//Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	//Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	//Chat the user belongs to
	Chat Chat `json:"chat"`
	//Performer of the action, which resulted in the change
	From User `json:"from"`
	//Date the change was done in Unix time
	Date int `json:"date"`
	//Previous information about the chat member
	OldChatMember ChatMember `json:"old_chat_member"`
	//New information about the chat member
	NewChatMember ChatMember `json:"new_chat_member"`
	//Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	//Optional. True, if the user joined the chat after sending a direct join request without using
	//an invite link and being approved by an administrator
	ViaJoinRequest *bool `json:"via_join_request,omitempty"`
	//Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink *bool `json:"via_chat_folder_invite_link,omitempty"`
}

// This object contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
//
// - ChatMemberOwner
//
// - ChatMemberAdministrator
//
// - ChatMemberMember
//
// - ChatMemberRestricted
//
// - ChatMemberLeft
//
// - ChatMemberBanned
type ChatMember struct {
	Status                string  `json:"status"`
	User                  *User   `json:"user,omitempty"`
	IsAnonymous           *bool   `json:"is_anonymous,omitempty"`
	CustomTitle           *string `json:"custom_title,omitempty"`
	CanBeEdited           *bool   `json:"can_be_edited,omitempty"`
	CanManageChat         *bool   `json:"can_manage_chat"`
	CanDeleteMessages     *bool   `json:"can_delete_messages,omitempty"`
	CanManageVideoChats   *bool   `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers    *bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     *bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         *bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool   `json:"can_invite_users,omitempty"`
	CanPostStories        *bool   `json:"can_post_stories,omitempty"`
	CanEditStories        *bool   `json:"can_edit_stories,omitempty"`
	CanDeleteStories      *bool   `json:"can_delete_stories,omitempty"`
	CanPostMessages       *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       *bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages        *bool   `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool   `json:"can_manage_topics,omitempty"`
	UntilDate             *int    `json:"until_date,omitempty,"`
	IsMember              *bool   `json:"is_member,omitempty"`
	CanSendMessages       *bool   `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool   `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     *bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebpagePreviews *bool   `json:"can_add_webpage_previews,omitempty"`
}

func (c ChatMember) OwnerType() (*ChatMemberOwner, error) {
	if c.Status != "creator" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "creator",
			Got:      c.Status,
		}
	}
	return &ChatMemberOwner{
		Status:      c.Status,
		User:        *c.User,
		IsAnonymous: *c.IsAnonymous,
		CustomTitle: c.CustomTitle,
	}, nil
}

func (c ChatMember) AdministratorType() (*ChatMemberAdministrator, error) {
	if c.Status != "“administrator”" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "“administrator”",
			Got:      c.Status,
		}
	}
	return &ChatMemberAdministrator{
		Status:              c.Status,
		User:                *c.User,
		CanBeEdited:         *c.CanBeEdited,
		IsAnonymous:         *c.IsAnonymous,
		CanManageChat:       *c.CanManageChat,
		CanDeleteMessages:   *c.CanDeleteMessages,
		CanManageVideoChats: *c.CanManageVideoChats,
		CanRestrictMembers:  *c.CanRestrictMembers,
		CanPromoteMembers:   *c.CanPromoteMembers,
		CanChangeInfo:       *c.CanChangeInfo,
		CanInviteUsers:      *c.CanInviteUsers,
		CanPostStories:      *c.CanPostStories,
		CanEditStories:      *c.CanEditStories,
		CanDeleteStories:    *c.CanDeleteStories,
		CanPostMessages:     c.CanPostMessages,
		CanEditMessages:     c.CanEditMessages,
		CanPinMessages:      c.CanPinMessages,
		CanManageTopics:     c.CanManageTopics,
		CustomTitle:         c.CustomTitle,
	}, nil
}

func (c ChatMember) MemberType() (*ChatMemberMember, error) {
	if c.Status != "member" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "member",
			Got:      c.Status,
		}
	}
	return &ChatMemberMember{
		Status:    c.Status,
		User:      *c.User,
		UntilDate: c.UntilDate,
	}, nil
}

func (c ChatMember) RestrictedType() (*ChatMemberRestricted, error) {
	if c.Status != "restricted" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "restricted",
			Got:      c.Status,
		}
	}
	return &ChatMemberRestricted{
		Status:                c.Status,
		User:                  c.User,
		IsMember:              *c.IsMember,
		CanSendMessages:       *c.CanSendMessages,
		CanSendAudios:         *c.CanSendAudios,
		CanSendDocuments:      *c.CanSendDocuments,
		CanSendPhotos:         *c.CanSendPhotos,
		CanSendVideos:         *c.CanSendVideos,
		CanSendVideoNotes:     *c.CanSendVideoNotes,
		CanSendVoiceNotes:     *c.CanSendVoiceNotes,
		CanSendPolls:          *c.CanSendPolls,
		CanSendOtherMessages:  *c.CanSendOtherMessages,
		CanAddWebpagePreviews: *c.CanAddWebpagePreviews,
		CanChangeInfo:         *c.CanChangeInfo,
		CanInviteUsers:        *c.CanInviteUsers,
		CanPinMessages:        *c.CanPinMessages,
		CanManageTopics:       *c.CanManageTopics,
		UntilDate:             *c.UntilDate,
	}, nil
}

func (c ChatMember) LeftType() (*ChatMemberLeft, error) {
	if c.Status != "left" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "left",
			Got:      c.Status,
		}
	}
	return &ChatMemberLeft{
		Status: c.Status,
		User:   c.User,
	}, nil
}

func (c ChatMember) BannedType() (*ChatMemberBanned, error) {
	if c.Status != "kicked" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "kicked",
			Got:      c.Status,
		}
	}
	return &ChatMemberBanned{
		Status:    c.Status,
		User:      *c.User,
		UntilDate: *c.UntilDate,
	}, nil
}

// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	//The member's status in the chat, always “creator”
	Status string `json:"status"`
	//Information about the user
	User User `json:"user"`
	//True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
	//Optional. Custom title for this user
	CustomTitle *string `json:"custom_title,omitempty"`
}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	//The member's status in the chat, always “administrator”
	Status string `json:"status"`
	//Information about the user
	User User `json:"user"`
	//True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited"`
	//True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
	//True, if the administrator can access the chat event log, get boost list,
	//see hidden supergroup and channel members, report spam messages and ignore slow mode.
	//Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`
	//True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`
	//True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	//True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members"`
	//True, if the administrator can add new administrators with a subset of their own privileges or
	//demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`
	//True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	//True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	//True, if the administrator can post stories to the chat
	CanPostStories bool `json:"can_post_stories"`
	//True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories bool `json:"can_edit_stories"`
	//True, if the administrator can delete stories posted by other users
	CanDeleteStories bool `json:"can_delete_stories"`
	//Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`
	//Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`
	//Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	//Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
	//Optional. Custom title for this user
	CustomTitle *string `json:"custom_title,omitempty"`
}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	//The member's status in the chat, always “member”
	Status string `json:"status"`
	//Information about the user
	User User `json:"user"`
	//Optional. Date when the user's subscription will expire; Unix time
	UntilDate *int `json:"until_date,omitempty"`
}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	//The member's status in the chat, always “restricted”
	Status string `json:"status"`
	//Information about the user
	User *User `json:"user"`
	//True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`
	//True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages"`
	//True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`
	//True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`
	//True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`
	//True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`
	//True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`
	//True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`
	//True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`
	//True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	//True, if the user is allowed to add web page previews to their messages
	CanAddWebpagePreviews bool `json:"can_add_webpage_previews"`
	//True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	//True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	//True, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages"`
	//True, if the user is allowed to create forum topics
	CanManageTopics bool `json:"can_manage_topics"`
	//Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever
	UntilDate int `json:"until_date"`
}

// Represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	//The member's status in the chat, always “left”
	Status string `json:"status"`
	//Information about the user
	User *User `json:"user"`
}

// Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	//The member's status in the chat, always “kicked”
	Status string `json:"status"`
	//Information about the user
	User User `json:"user"`
	//Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever
	UntilDate int `json:"until_date"`
}

// Represents a join request sent to a chat.
type ChatJoinRequest struct {
	//Chat to which the request was sent
	Chat Chat `json:"chat"`
	//User that sent the join request
	User User `json:"user"`
	//Identifier of a private chat with the user who sent the join request.
	//This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	//The bot can use this identifier for 5 minutes to send messages until the join request is processed,
	//assuming no other administrator contacted the user.
	UserChatId int64 `json:"user_chat_id"`
	//Date the request was sent in Unix time
	Date int `json:"date"`
	//Optional. Bio of the user.
	Bio *string `json:"bio,omitempty"`
	//Optional. Chat invite link that was used by the user to send the join request
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	//True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendMessages *bool `json:"can_send_messages,omitempty"`
	//True, if the user is allowed to send audios
	CanSendAudios *bool `json:"can_send_audios,omitempty"`
	//True, if the user is allowed to send documents
	CanSendDocuments *bool `json:"can_send_documents,omitempty"`
	//True, if the user is allowed to send photos
	CanSendPhotos *bool `json:"can_send_photos,omitempty"`
	//True, if the user is allowed to send videos
	CanSendVideos *bool `json:"can_send_videos,omitempty"`
	//True, if the user is allowed to send video notes
	CanSendVideoNotes *bool `json:"can_send_video_notes,omitempty"`
	//True, if the user is allowed to send voice notes
	CanSendVoiceNotes *bool `json:"can_send_voice_notes,omitempty"`
	//True, if the user is allowed to send polls
	CanSendPolls *bool `json:"can_send_polls,omitempty"`
	//True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty"`
	//True, if the user is allowed to add web page previews to their messages
	CanAddWebpagePreviews *bool `json:"can_add_webpage_previews,omitempty"`
	//True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo *bool `json:"can_change_info,omitempty"`
	//True, if the user is allowed to invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`
	//True, if the user is allowed to pin messages
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	//True, if the user is allowed to create forum topics
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// Describes the birthdate of a user.
type BirthDate struct {
	//Day of the user's birth; 1-31
	Day int `json:"day"`
	//Month of the user's birth; 1-12
	Month int `json:"month"`
	//Optional. Year of the user's birth
	Year *int `json:"year,omitempty"`
}

// Contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	//Optional. Title text of the business intro
	Title string `json:"title,omitempty"`
	//Optional. Message text of the business intro
	Message *string `json:"message,omitempty"`
	//Optional. Sticker of the business intro
	Sticker *Sticker `json:"sticker,omitempty"`
}

// Contains information about the location of a Telegram Business account.
type BusinessLocation struct {
	//Address of the business
	Address string `json:"address"`
	//Optional. Location of the business
	Location *Location `json:"location,omitempty"`
}

// Describes an interval of time during which a business is open.
type BusinessOpeningHoursInterval struct {
	//The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	OpeningMinute int `json:"opening_minute"`
	//The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60
	ClosingMinute int `json:"closing_minute"`
}

// Describes the opening hours of a business.
type BusinessOpeningHours struct {
	//Unique name of the time zone for which the opening hours are defined
	TimeZone string `json:"time_zone"`
	//List of time intervals describing business opening hours
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// Represents a location to which a chat is connected.
type ChatLocation struct {
	//The location to which the supergroup is connected. Can't be a live location.
	Location Location `json:"location"`
	//Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// This object describes the type of a reaction. Currently, it can be one of:
//
// - ReactionTypeEmoji
//
// - ReactionTypeCustomEmoji
//
// - ReactionTypePaid
type ReactionType interface {
	GetReactionType() string
}

// The reaction is based on an emoji.
type ReactionTypeEmoji struct {
	//Type of the reaction, always “emoji”
	Type string `json:"type"`
	//Reaction emoji. Currently, it can be one of
	//"👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢",
	//"🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱", "🥴", "😍", "🐳",
	//"❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨", "😐", "🍓",
	//"🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃", "🙈",
	//"😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿",
	//"🆒", "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"
	Emoji string `json:"emoji"`
}

func (r ReactionTypeEmoji) GetReactionType() string {
	return "emoji"
}

func (r ReactionTypeEmoji) Validate() error {
	if strings.TrimSpace(r.Emoji) == "" {
		return ErrInvalidParam("emoji parameter can't be empty")
	}
	if r.Type != "emoji" {
		return ErrInvalidParam("type must be \"emoji\"")
	}
	return nil
}

// The reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	//Type of the reaction, always “custom_emoji”
	Type string `json:"type"`
	//Custom emoji identifier
	CustomEmojiId string `json:"custom_emoji_id"`
}

func (r ReactionTypeCustomEmoji) GetReactionType() string {
	return "custom_emoji"
}

func (r ReactionTypeCustomEmoji) Validate() error {
	if strings.TrimSpace(r.CustomEmojiId) == "" {
		return ErrInvalidParam("custom_emoji_id parameter can't be empty")
	}
	if r.Type != "custom_emoji" {
		return ErrInvalidParam("type must be \"custom_emoji\"")
	}
	return nil
}

// The reaction is paid.
type ReactionTypePaid struct {
	//Type of the reaction, always “paid”
	Type string `json:"type"`
}

func (r ReactionTypePaid) GetReactionType() string {
	return "paid"
}

func (r ReactionTypePaid) Validate() error {
	if r.Type != "paid" {
		return ErrInvalidParam("type must be\"paid\"")
	}
	return nil
}

// Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	//Type of the reaction
	Type ReactionType `json:"type"`
	//Number of times the reaction was added
	TotalCount int `json:"total_count"`
}

// This object represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	//The chat containing the message the user reacted to
	Chat Chat `json:"chat"`
	//Unique identifier of the message inside the chat
	MessageId int `json:"message_id"`
	//Optional. The user that changed the reaction, if the user isn't anonymous
	User *User `json:"user,omitempty"`
	//Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat *Chat `json:"actor_chat,omitempty"`
	//Date of the change in Unix time
	Date int `json:"date"`
	//Previous list of reaction types that were set by the user
	OldReaction []ReactionType `json:"old_reaction"`
	//New list of reaction types that have been set by the user
	NewReaction []ReactionType `json:"new_reaction"`
}

// This object represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	//The chat containing the message
	Chat Chat `json:"chat"`
	//Unique message identifier inside the chat
	MessageId int `json:"message_id"`
	//Date of the change in Unix time
	Date int `json:"date"`
	//List of reactions that are present on the message
	Reactions []ReactionCount `json:"reactions"`
}

// This object represents a forum topic.
type ForumTopic struct {
	//Unique identifier of the forum topic
	MessageThreadId int `json:"message_thread_id"`
	//Name of the topic
	Name string `json:"name"`
	//Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`
	//Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

// This object represents a bot command.
type BotCommand struct {
	//Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`
	//Description of the command; 1-256 characters.
	Description string `json:"description"`
}

var valid_command = regexp.MustCompile(`^[a-z0-9_]+$`)

func (b BotCommand) Validate() error {
	if len(b.Command) < 1 || len(b.Command) > 32 {
		return ErrInvalidParam("command parameter must be between 1 and 32 characters")
	}
	if !valid_command.MatchString(b.Command) {
		return ErrInvalidParam("command parameter can contain only lowercase English letters, digits and underscores")
	}
	if len(b.Description) < 1 || len(b.Description) > 256 {
		return ErrInvalidParam("description parameter must be between 1 and 256 characters")
	}
	return nil
}

// This object represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
//
// - -BotCommandScopeDefault
//
// - -BotCommandScopeAllPrivateChats
//
// - -BotCommandScopeAllGroupChats
//
// - -BotCommandScopeAllChatAdministrators
//
// - -BotCommandScopeChat
//
// - -BotCommandScopeChatAdministrators
//
// - -BotCommandScopeChatMember
type BotCommandScope interface {
	GetBotCommandScopeType() string
	Validable
}

// Represents the default scope of bot commands. Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	//Scope type, must be default
	Type string `json:"type"`
}

func (b BotCommandScopeDefault) GetBotCommandScopeType() string {
	return "default"
}

func (b BotCommandScopeDefault) Validate() error {
	if b.Type != "default" {
		return ErrInvalidParam("type must be \"default\"")
	}
	return nil
}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	//Scope type, must be all_private_chats
	Type string `json:"type"`
}

func (b BotCommandScopeAllPrivateChats) GetBotCommandScopeType() string {
	return "all_private_chats"
}

func (b BotCommandScopeAllPrivateChats) Validate() error {
	if b.Type != "all_private_chats" {
		return ErrInvalidParam("type must be \"all_private_chats\"")
	}
	return nil
}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	//Scope type, must be all_group_chats
	Type string `json:"type"`
}

func (b BotCommandScopeAllGroupChats) GetBotCommandScopeType() string {
	return "all_group_chats"
}

func (b BotCommandScopeAllGroupChats) Validate() error {
	if b.Type != "all_group_chats" {
		return ErrInvalidParam("type must be \"all_group_chats\"")
	}
	return nil
}

// Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	//Scope type, must be all_chat_administrators
	Type string `json:"type"`
}

func (b BotCommandScopeAllChatAdministrators) GetBotCommandScopeType() string {
	return "all_chat_administrators"
}

func (b BotCommandScopeAllChatAdministrators) Validate() error {
	if b.Type != "all_chat_administrators" {
		return ErrInvalidParam("type must be \"all_chat_administrators\"")
	}
	return nil
}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat[T int | string] struct {
	//Scope type, must be chat
	Type string `json:"type"`
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId T `json:"chat_id"`
}

func (b BotCommandScopeChat[T]) GetBotCommandScopeType() string {
	return "chat"
}

func (b BotCommandScopeChat[T]) Validate() error {
	if b.Type != "chat" {
		return ErrInvalidParam("type must be \"chat\"")
	}
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return ErrInvalidParam("chat_id paramter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c == 0 {
			return ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

// Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators[T int | string] struct {
	//Scope type, must be chat_administrators
	Type string `json:"type"`
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId T `json:"chat_id"`
}

func (b BotCommandScopeChatAdministrators[T]) GetBotCommandScopeType() string {
	return "chat_administrators"
}

func (b BotCommandScopeChatAdministrators[T]) Validate() error {
	if b.Type != "chat_administrators" {
		return ErrInvalidParam("type must be \"chat_administrators\"")
	}
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return ErrInvalidParam("chat_id paramter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c == 0 {
			return ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

// Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember[T int | string] struct {
	//Scope type, must be chat_member
	Type string `json:"type"`
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId T `json:"chat_id"`
	//Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (b BotCommandScopeChatMember[T]) GetBotCommandScopeType() string {
	return "chat_member"
}

func (b BotCommandScopeChatMember[T]) Validate() error {
	if b.Type != "chat_member" {
		return ErrInvalidParam("type must be \"chat_member\"")
	}
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return ErrInvalidParam("chat_id paramter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c == 0 {
			return ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.UserId < 1 {
		return ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

// This object represents the bot's name.
type BotName struct {
	//The bot's name
	Name string
}

// This object represents the bot's description.
type BotDescription struct {
	//The bot's description
	Description string `json:"description"`
}

// This object represents the bot's short description.
type BotShortDescription struct {
	//The bot's short description
	ShortDescription string `json:"short_description"`
}

// This object describes the bot's menu button in a private chat. It should be one of
//
// - MenuButtonCommands
//
// - MenuButtonWebApp
//
// - MenuButtonDefault
//
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat.
// Otherwise the default menu button is applied. By default, the menu button opens the list of bot commands.
type MenuButton interface {
	GetMenuButtonType() string
	Validate() error
}

// MenuButtonResponse represents the generalized form of a bot's menu button in a private chat.
// This type is used exclusively for reading responses from the `getMenuButton` method.
//
// It serves as a base structure containing all possible fields of different menu button types:
//
// - MenuButtonCommands
//
// - MenuButtonWebApp
//
// - MenuButtonDefault
//
// This type provides methods to convert itself into a more specific `MenuButton` implementation
// for further processing.
type MenuButtonResponse struct {
	Type   string      `json:"type"`
	Text   *string     `json:"text,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// NOTE: do i really need it when the only case when this conversion is making sense
	// is when the type is web_app
}

func (m MenuButtonResponse) CommandsType() (*MenuButtonCommands, error) {
	if m.Type != "commands" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "commands",
			Got:      m.Type,
		}
	}
	return &MenuButtonCommands{
		Type: m.Type,
	}, nil
}

func (m MenuButtonResponse) WebAppType() (*MenuButtonWebApp, error) {
	if m.Type != "web_app" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "web_app",
			Got:      m.Type,
		}
	}
	return &MenuButtonWebApp{
		Type:   m.Type,
		Text:   *m.Text,
		WebApp: *m.WebApp,
	}, nil
}

func (m MenuButtonResponse) DefaultType() (*MenuButtonDefault, error) {
	if m.Type != "default" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "default",
			Got:      m.Type,
		}
	}
	return &MenuButtonDefault{
		Type: m.Type,
	}, nil
}

// Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	//Type of the button, must be commands
	Type string `json:"type"`
}

func (m MenuButtonCommands) Validate() error {
	if m.Type != "commands" {
		return ErrInvalidParam("type must be \"commands\"")
	}
	return nil
}

func (m MenuButtonCommands) GetMenuButtonType() string {
	return "commands"
}

// Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	//Type of the button, must be web_app
	Type string `json:"type"`
	//Text on the button
	Text string `json:"text"`
	//Description of the Web App that will be launched when the user presses the button.
	//The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	//Alternatively, a t.me link to a Web App of the bot can be specified in the object instead of the Web App's URL,
	//in which case the Web App will be opened as if the user pressed the link.
	WebApp WebAppInfo `json:"web_app"`
}

func (m MenuButtonWebApp) GetMenuButtonType() string {
	return "web_app"
}

func (m MenuButtonWebApp) Validate() error {
	if m.Type != "web_app" {
		return ErrInvalidParam("type must be \"web_app\"")
	}
	if strings.TrimSpace(m.Text) == "" {
		return ErrInvalidParam("text parameter can't be empty")
	}
	if err := m.WebApp.Validate(); err != nil {
		return err
	}
	return nil
}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	//Type of the button, must be default
	Type string `json:"type"`
}

func (m MenuButtonDefault) Validate() error {
	if m.Type != "default" {
		return ErrInvalidParam("type must be \"default\"")
	}
	return nil
}

func (m MenuButtonDefault) GetMenuButtonType() string {
	return "default"
}

// This object describes the source of a chat boost. It can be one of
//
// - ChatBoostSourcePremium
//
// - ChatBoostSourceGiftCode
//
// - ChatBoostSourceGiveaway
type ChatBoostSource struct {
	Source            string  `json:"source"`
	User              *User   `json:"user,omitempty"`
	GiveawayMessageId *string `json:"giveaway_message_id,omitempty"`
	PrizeStarCount    *int    `json:"prize_star_count,omitempty"`
	IsUnclaimed       *bool   `json:"is_unclaimed,omitempty"`
}

func (c ChatBoostSource) PremiumType() (*ChatBoostSourcePremium, error) {
	if c.Source != "premium" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "premium",
			Got:      c.Source,
		}
	}
	return &ChatBoostSourcePremium{
		Source: c.Source,
		User:   *c.User,
	}, nil
}

func (c ChatBoostSource) GiftCodeType() (*ChatBoostSourceGiftCode, error) {
	if c.Source != "gift_code" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "gift_code",
			Got:      c.Source,
		}
	}
	return &ChatBoostSourceGiftCode{
		Source: c.Source,
		User:   *c.User,
	}, nil
}

func (c ChatBoostSource) GiveawayType() (*ChatBoostSourceGiveaway, error) {
	if c.Source != "giveaway" {
		return nil, ErrInvalidSubtypeConversion{
			Expected: "giveaway",
			Got:      c.Source,
		}
	}
	return &ChatBoostSourceGiveaway{
		Source:            c.Source,
		GiveawayMessageId: *c.GiveawayMessageId,
		User:              c.User,
		PrizeStarCount:    c.PrizeStarCount,
		IsUnclaimed:       c.IsUnclaimed,
	}, nil
}

// The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	//Source of the boost, always “premium”
	Source string `json:"source"`
	//User that boosted the chat
	User User `json:"user"`
}

// The boost was obtained by the creation of Telegram Premium gift codes to boost a chat.
// Each such code boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiftCode struct {
	//Source of the boost, always “gift_code”
	Source string `json:"source"`
	//User for which the gift code was created
	User User `json:"user"`
}

// The boost was obtained by the creation of a Telegram Premium or a Telegram Star giveaway.
// This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription for
// Telegram Premium giveaways and prize_star_count / 500 times for one year for Telegram Star giveaways.
type ChatBoostSourceGiveaway struct {
	//Source of the boost, always “giveaway”
	Source string `json:"source"`
	//Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	GiveawayMessageId string `json:"giveaway_message_id"`
	//Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	User *User `json:"user,omitempty"`
	//Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
	//Optional. True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed *bool `json:"is_unclaimed,omitempty"`
}

// This object contains information about a chat boost.
type ChatBoost struct {
	//Unique identifier of the boost
	BoostId string `json:"boost_id"`
	//Point in time (Unix timestamp) when the chat was boosted
	AddDate int `json:"add_date"`
	//Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	ExpirationDate int `json:"expiration_date"`
	//Source of the added boost
	Source ChatBoostSource `json:"source"`
}

// This object represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	//Chat which was boosted
	Chat Chat `json:"chat"`
	//Information about the chat boost
	Boost ChatBoost `json:"boost"`
}

// This object represents a boost removed from a chat.
type ChatBoostRemoved struct {
	//Chat which was boosted
	Chat Chat `json:"chat"`
	//Unique identifier of the boost
	BoostId string `json:"boost_id"`
	//Point in time (Unix timestamp) when the boost was removed
	RemoveDate int `json:"remove_date"`
	//Source of the removed boost
	Source ChatBoostSource `json:"source"`
}

// This object represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	//The list of boosts added to the chat by the user
	Boosts []ChatBoost `json:"boosts"`
}

// Describes the connection of the bot with a business account.
type BusinessConnection struct {
	//Unique identifier of the business connection
	Id string `json:"id"`
	//Business account user that created the business connection
	User User `json:"user"`
	//Identifier of a private chat with the user who created the business connection.
	//This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserChatId int `json:"user_chat_id"`
	//Date the connection was established in Unix time
	Date int `json:"date"`
	//True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	CanReply bool `json:"can_reply"`
	//True, if the connection is active
	IsEnabled bool `json:"is_enabled"`
}

// /This object is received when messages are deleted from a connected business account.
type BusinessMessagesDeleted struct {
	//Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	//Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	Chat Chat `json:"chat"`
	//The list of identifiers of deleted messages in the chat of the business account
	MessageIds []int `json:"message_ids"`
}

// Describes why a request was unsuccessful.
type ResponseParameters struct {
	//Optional. The group has been migrated to a supergroup with the specified identifier.
	//This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId *int `json:"migrate_to_chat_id,omitempty"`
	//Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter *int `json:"retry_after,omitempty"`
}

// This object represents the content of a media message to be sent. It should be one of
//
// - InputMediaAnimation
//
// - InputMediaDocument
//
// - InputMediaAudio
//
// - InputMediaPhoto
//
// - InputMediaVideo
type InputMedia interface {
	SetInputMedia(media string, isNew bool)
	Validate() error
}

var (
	url_regex        = regexp.MustCompile(`^https?://`)
	attachment_regex = regexp.MustCompile(`^attach://[\w-]+$`)
)

// Represents a photo to be sent.
//
// IMPORTANT: It is strongly recommended to use the SetInputMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputMediaPhoto struct {
	//Type of the result, must be photo
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	//Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the photo caption. See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	isNew      bool  `json:"-"`
}

func (i *InputMediaPhoto) SetInputMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
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
	if i.Type != "photo" {
		return ErrInvalidParam("type must be photo")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}

	return nil
}

// Represents a video to be sent.
//
// IMPORTANT: It is strongly recommended to use the SetInputMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputMediaVideo struct {
	//Type of the result, must be video
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	//Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size.
	//A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	//the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFile `json:"thumbnail,omitempty"`
	//Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the video caption. See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Video width
	Width *int `json:"width,omitempty"`
	//Optional. Video height
	Height *int `json:"height,omitempty"`
	//Optional. Video duration in seconds
	Duration *int `json:"duration,omitempty"`
	//Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
	//Optional. Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	isNew      bool  `json:"-"`
}

func (i *InputMediaVideo) SetInputMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
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
	if i.Type != "video" {
		return ErrInvalidParam("type must be video")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}

	return nil
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//
// IMPORTANT: It is strongly recommended to use the SetInputMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputMediaAnimation struct {
	//Type of the result, must be animation
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	//Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size.
	//A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	//the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFile `json:"thumbnail,omitempty"`
	//Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the animation caption. See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional. Animation width
	Width *int `json:"width,omitempty"`
	//Optional. Animation height
	Height *int `json:"height,omitempty"`
	//Optional. Animation duration in seconds
	Duration *int `json:"duration,omitempty"`
	//Optional. Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	isNew      bool  `json:"-"`
}

func (i *InputMediaAnimation) SetInputMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
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
	if i.Type != "animation" {
		return ErrInvalidParam("type must be animation")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}
	return nil
}

// Represents an audio file to be treated as music to be sent.
//
// IMPORTANT: It is strongly recommended to use the SetInputMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputMediaAudio struct {
	//Type of the result, must be audio
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	//Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size.
	//A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	//the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFile `json:"thumbnail,omitempty"`
	//Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the audio caption. See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Duration of the audio in seconds
	Duration *int `json:"duration,omitempty"`
	//Optional. Performer of the audio
	Performer *string `json:"performer,omitempty"`
	//Optional. Title of the audio
	Title *string `json:"title,omitempty"`
	isNew bool    `json:"-"`
}

func (i *InputMediaAudio) SetInputMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
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
	if i.Type != "audio" {
		return ErrInvalidParam("type must be audio")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}
	return nil
}

// Represents a general file to be sent.
//
// IMPORTANT: It is strongly recommended to use the SetInputMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputMediaDocument struct {
	//Type of the result, must be document
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	//Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size.
	//A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	//the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFile `json:"thumbnail,omitempty"`
	//Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Optional. Mode for parsing entities in the document caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	//Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data.
	//Always True, if the document is sent as part of an album.
	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
	isNew                       bool  `json:"-"`
}

func (i *InputMediaDocument) SetInputMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
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
	if i.Type != "document" {
		return ErrInvalidParam("type must be document")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}
	return nil
}

// FIXME: should be an interface with 3 implementations:
// 1. a byte slice
// 2. a file path
// 3. a file id
type InputFile interface {
	Validable
	Name() string
	Reader() (io.Reader, error)
	IsLocal() bool
}

type InputFileFromPath struct {
	FilePath string
}

func (i InputFileFromPath) Name() string {
	return filepath.Base(i.FilePath)
}

func (i InputFileFromPath) Reader() (io.Reader, error) {
	return os.Open(i.FilePath)
}

func (i InputFileFromPath) Validate() error {
	if i.FilePath == "" {
		return ErrInvalidParam("file path can't be empty")
	}
	if _, err := os.Stat(i.FilePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist at path: %s", i.FilePath)
	}
	return nil
}

func (i InputFileFromPath) IsLocal() bool {
	return true
}

type InputFileFromBytes struct {
	FileName string
	Data     []byte
}

func (i InputFileFromBytes) Name() string {
	return i.FileName
}

func (i InputFileFromBytes) Reader() (io.Reader, error) {
	return bytes.NewReader(i.Data), nil
}

func (i InputFileFromBytes) Validate() error {
	if i.FileName == "" {
		return ErrInvalidParam("file name can't be empty")
	}
	if len(i.Data) == 0 {
		return ErrInvalidParam("data can't be empty")
	}
	return nil
}

func (i InputFileFromBytes) IsLocal() bool {
	return true
}

type InputFileFromString string

func (i InputFileFromString) Validate() error {
	if i == "" {
		return ErrInvalidParam("file id or url can't be empty")
	}
	return nil
}

func (i InputFileFromString) Name() string {
	return ""
}

func (i InputFileFromString) Reader() (io.Reader, error) {
	return nil, fmt.Errorf("remote file can't have reader")
}

func (i InputFileFromString) IsLocal() bool {
	return false
}

// This object describes the paid media to be sent. Currently, it can be one of
//
// - InputPaidMediaPhoto
//
// - InputPaidMediaVideo
type InputPaidMedia interface {
	SetInputPaidMedia(media string, isNew bool)
	Validate() error
}

// The paid media to send is a photo.
//
// IMPORTANT: It is strongly recommended to use the SetInputPaidMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputPaidMediaPhoto struct {
	//Type of the media, must be photo
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	isNew bool   `json:"-"`
}

func (i InputPaidMediaPhoto) Validate() error {
	if i.Type != "photo" {
		return ErrInvalidParam("type must be photo")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}
	return nil
}

func (i *InputPaidMediaPhoto) SetInputPaidMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
			i.Media = media
		} else {
			i.Media = "attach://" + media
		}
	} else {
		i.Media = media
		i.isNew = false
	}
}

// The paid media to send is a video.
//
// IMPORTANT: It is strongly recommended to use the SetInputPaidMedia method to ensure
// proper handling of file attachments, including the use of "attach://" prefixes
// for new files or validation of media URLs.
type InputPaidMediaVideo struct {
	//Type of the media, must be video
	Type string `json:"type"`
	//File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL for Telegram to get a file from the Internet,
	//or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	//Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size.
	//A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file,
	//so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFile `json:"thumbnail,omitempty"`
	//Optional. Video width
	Width *int `json:"width,omitempty"`
	//Optional. Video height
	Height *int `json:"height,omitempty"`
	//Optional. Video duration in seconds
	Duration *int `json:"duration,omitempty"`
	//Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
	isNew             bool  `json:"-"`
}

func (i *InputPaidMediaVideo) SetInputPaidMedia(media string, isNew bool) {
	if isNew {
		if url_regex.MatchString(media) {
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
	if i.Type != "video" {
		return ErrInvalidParam("type must be video")
	}
	if strings.TrimSpace(i.Media) == "" {
		return ErrInvalidParam("media parameter can't be empty")
	}

	if i.isNew {
		if !url_regex.MatchString(i.Media) && !attachment_regex.MatchString(i.Media) {
			return ErrInvalidParam("invalid media parameter. please refer to https://core.telegram.org/bots/api#sending-files")
		}
	}
	return nil
}
