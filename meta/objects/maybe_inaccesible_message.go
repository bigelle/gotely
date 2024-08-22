package objects

type MaybeInaccessibleMessage interface {
	IsUserMessage() bool
	IsGroupMessage() bool
	IsSuperGroupMessage() bool
	GetChatId() int64
	GetChat() Chat 
	GetMessageId() int
	GetDate() int
}
