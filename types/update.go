package types

type Update struct {
	UpdateId               int                          `json:"update_id"`
	Message                *Message                     `json:"message,omitempty"`
	InlineQuery            *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineQuery      *ChosenInlineResult          `json:"chosen_inline_query,omitempty"`
	CallbackQuery          *CallbackQuery               `json:"callback_query,omitempty"`
	EditedMessage          *Message                     `json:"edited_message,omitempty"`
	ChannelPost            *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost      *Message                     `json:"edited_channel_post,omitempty"`
	ShippingQuery          *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery       *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	Poll                   *Poll                        `json:"poll,omitempty"`
	PollAnswer             *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember           *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember             *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest        *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	MessageReaction        *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount   *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	ChatBoost              *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost       *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	BusinessConnection     *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage        *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage  *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessage *BusinessMessagesDeleted     `json:"deleted_business_message,omitempty"`
	PaidMediaPurchased     *PaidMediaPurchased          `json:"paid_media_purchased,omitempty"`
}
