package objects

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
	//FIXME
	//Optional. New incoming message of any kind - text, photo, sticker, etc.
	//	Message *Message `json:"message,omitempty"`
	//Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	//	ChannelPost *Message `json:"channel_post,omitempty"`
	//Optional. New version of a channel post that is known to the bot and was edited.
	//This update may at times be triggered by changes to message fields
	//that are either unavailable or not actively used by your bot.
	//	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	//Optional. New version of a message that is known to the bot and was edited.
	//This update may at times be triggered by changes to message fields
	//that are either unavailable or not actively used by your bot.
	//	EditedMessage *Message `json:"edited_message,omitempty"`
	//Optional. The bot was connected to or disconnected from a business account,
	//or a user edited an existing connection with the bot
	//	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`
	//Optional. New message from a connected business account
	//	BusinessMessage *Message `json:"business_message,omitempty"`
	//Optional. New version of a message from a connected business account
	//	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`
	//Optional. Messages were deleted from a connected business account
	//	DeletedBusinessMessage *BusinessMessagesDeleted `json:"deleted_business_message,omitempty"`
	//Optional. A reaction to a message was changed by a user.
	//The bot must be an administrator in the chat and must explicitly specify "message_reaction"
	//in the list of allowed_updates to receive these updates.
	//The update isn't received for reactions set by bots.
	//	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`
	//Optional. Reactions to a message with anonymous reactions were changed.
	//The bot must be an administrator in the chat and must explicitly specify "message_reaction_count"
	//in the list of allowed_updates to receive these updates.
	//The updates are grouped and can be sent with delay up to a few minutes.
	//	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	//Optional. New incoming inline query
	//	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	//Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	//Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	//	ChosenInlineQuery *ChosenInlineResult `json:"chosen_inline_query,omitempty"`
	//Optional. New incoming callback query
	//	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	//Optional. New incoming shipping query. Only for invoices with flexible price
	//	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	//Optional. A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	//	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`
	//Optional. New incoming pre-checkout query. Contains full information about checkout
	//	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	//Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	//	Poll *Poll `json:"poll,omitempty"`
	//Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	//	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	//Optional. The bot's chat member status was updated in a chat.
	//For private chats, this update is received only when the bot is blocked or unblocked by the user.
	//	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	//Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and
	//must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	//	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	//Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	//	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
	//Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	//	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`
	//Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	//	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}
