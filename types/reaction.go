package types

import (
	"encoding/json"
	"fmt"

	"github.com/bigelle/tele.go/internal/assertions"
)

type MessageReactionCountUpdated struct {
	Chat      Chat           `json:"chat"`
	MessageId int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type MessageReactionUpdated struct {
	Chat        Chat          `json:"chat"`
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

// TODO: refactor and add marshal/demarshal funcs
type ReactionType struct {
	ReactionTypeInterface
}

type ReactionTypeInterface interface {
	reactionTypeEmojiContract()
}

func (r *ReactionType) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "emoji":
		r.ReactionTypeInterface = new(ReactionTypeEmoji)
	case "custom_emoji":
		r.ReactionTypeInterface = new(ReactionTypeCustomEmoji)
	case "paid":
		r.ReactionTypeInterface = new(ReactionTypePaid)
	default:
		return fmt.Errorf("Unrecognized type: %t", r.ReactionTypeInterface)
	}
	return json.Unmarshal(raw.Attributes, &r.ReactionTypeInterface)
}

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"` // = CustomEmojiType
	CustomEmojiId string `json:"custom_emoji_id"`
}

func (r ReactionTypeCustomEmoji) reactionTypeEmojiContract() {}

func (r ReactionTypeCustomEmoji) Vaidate() error {
	if assertions.IsStringEmpty(r.CustomEmojiId) {
		return fmt.Errorf("CustomEmojiId parameter cant' be empty")
	}
	if r.Type != "custom_emoji" {
		return fmt.Errorf("Type must be \"custom_emoji\"")
	}
	return nil
}

type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

func (r ReactionTypeEmoji) reactionTypeEmojiContract() {}

func (r ReactionTypeEmoji) Validate() error {
	if assertions.IsStringEmpty(r.Emoji) {
		return fmt.Errorf("Emoji parameter can't be empty")
	}
	if r.Type != "emoji" {
		return fmt.Errorf("Type must be \"emoji\"")
	}
	return nil
}

type ReactionTypePaid struct {
	Type string `json:"type"`
}

func (r ReactionTypePaid) reactionTypeEmojiContract() {}

func (r ReactionTypePaid) Validate() error {
	if r.Type != "paid" {
		return fmt.Errorf("Type must be\"paid\"")
	}
	return nil
}
