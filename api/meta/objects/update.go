package objects

type Update struct {
	UpdateId               int
	Message                *Message
	InlineQuery            *InlineQuery       //TODO: implement it
	ChosenInlineQuery      *ChosenInlineQuery //TODO: implement it
	CallbackQuery          *CallbackQuery     //TODO: implement it
	EditedMessage          *Message
	ChannelPost            *Message
	EditedChannelPost      *Message
	ShippingQuery          *ShippingQuery               //TODO: implement it
	PreCheckoutQuery       *PreCheckoutQuery            //TODO: implement it
	Poll                   *Poll                        //TODO: implement it
	PollAnswer             *PollAnswer                  //TODO: implement it
	MyChatMember           *ChatMemberUpdated           //TODO: implement it
	ChatMember             *ChatMemberUpdated           //TODO: implement it
	ChatJoinRequest        *ChatJoinRequest             //TODO: implement it
	MessageReaction        *MessageReactionUpdated      //TODO: implement it
	MessageReactionCount   *MessageReactionCountUpdated //TODO: implement it
	ChatBoost              *ChatBoostUpdated
	RemovedChatBoost       *ChatBoostRemoved
	BusinessConnection     *BusinessConnection //TODO: implement it
	BusinessMessage        *Message
	EditedBusinessMessage  *Message
	DeletedBusinessMessage *BusinessMessageDeleted //TODO: implement it
}

func (u Update) HasMessage() bool {
	return u.Message != nil
}

func (u Update) HasInlineQuery() bool {
	return u.InlineQuery != nil
}

func (u Update) HasChosenInlineQuery() bool {
	return u.ChosenInlineQuery != nil
}

func (u Update) HasCallbackQuery() bool {
	return u.CallbackQuery != nil
}

func (u Update) HasEditedMessage() bool {
	return u.EditedMessage != nil
}

func (u Update) HasChannelPost() bool {
	return u.ChannelPost != nil
}

func (u Update) HasEditedChannelPost() bool {
	return u.EditedChannelPost != nil
}

func (u Update) HasShippingQuery() bool {
	return u.ShippingQuery != nil
}

func (u Update) HasPreCheckoutQuery() bool {
	return u.PreCheckoutQuery != nil
}

func (u Update) HasPoll() bool {
	return u.Poll != nil
}

func (u Update) HasPollAnswer() bool {
	return u.PollAnswer != nil
}

func (u Update) HasMyChatMember() bool {
	return u.MyChatMember != nil
}

func (u Update) HasChatMember() bool {
	return u.ChatMember != nil
}

func (u Update) HasChatJoinRequest() bool {
	return u.ChatJoinRequest != nil
}

func (u Update) HasBusinessConnection() bool {
	return u.BusinessConnection != nil
}

func (u Update) HasBusinessMessage() bool {
	return u.BusinessMessage != nil
}

func (u Update) HasEditedBusinessMessage() bool {
	return u.BusinessMessage != nil
}

func (u Update) HasDeletedBusinessMessage() bool {
	return u.DeletedBusinessMessage != nil
}
