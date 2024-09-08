package objects

const (
	EmojiType       = "emoji"
	CustomEmojiType = "custom_emoji"
	PaidType        = "paid"
)

type ReactionType interface {
	GetType() string
}

type Reaction struct {
	Type string `json:"type"`
}
