package types

import "encoding/json"

type ChatBoost struct {
	BoostId        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
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

type ChatBoostSource struct {
	ChatBoostSourceInterface
}

type ChatBoostSourceInterface interface {
	chatBoostSourceContract()
}

func (c *ChatBoostSource) UnmarshalJSON(data []byte) error {
	var raw struct {
		Source     string
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Source {
	case "gift_code":
		c.ChatBoostSourceInterface = new(ChatBoostSourceGiftCode)
	case "giveaway":
		c.ChatBoostSourceInterface = new(ChatBoostSourceGiveaway)
	case "premium":
		c.ChatBoostSourceInterface = new(ChatBoostSourcePremium)
	}
	return json.Unmarshal(raw.Attributes, c.ChatBoostSourceInterface)
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

func (c ChatBoostSourceGiftCode) chatBoostSourceContract() {}

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageId string `json:"giveaway_message_id"`
	User              *User  `json:"user,omitempty"`
	IsUnclaimed       *bool  `json:"is_unclaimed,omitempty"`
	PrizeStarCount    *int   `json:"prize_star_count,omitempty"`
}

func (c ChatBoostSourceGiveaway) chatBoostSourceContract() {}

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

func (c ChatBoostSourcePremium) chatBoostSourceContract() {}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}
