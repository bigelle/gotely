package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bigelle/tele.go/assertions"
)

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
	if err := assertions.ParamNotEmpty(r.Emoji, "emoji"); err != nil {
		return err
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
