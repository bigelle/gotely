package objects

import (
	"encoding/json"

	"github.com/bigelle/tele.go/api/meta/interfaces"
)

type ChatBoost struct {
	BoostId        string
	AddDate        int
	ExpirationDate int
	Source         ChatBoostSource
}

type ChatBoostAdded struct {
	BoostCount int
}

type ChatBoostRemoved struct {
	Chat       Chat
	BoostId    string
	RemoveDate int
	Source     ChatBoostSource
}

const (
	PremiumType  = "premium"
	GiftCodeType = "gift_code"
	GiveawayType = "giveaway"
)

type ChatBoostSource interface {
	interfaces.BotApiObject
}

type ChatBoostSourceType struct {
	// TODO: connect it with other types
}

type ChatBoostSourceGiftCode struct {
	Source string
	User   User
}

func (c ChatBoostSourceGiftCode) MarshalJSON() ([]byte, error) {
	type Alias ChatBoostSourceGiftCode
	return json.Marshal(&struct {
		Source string
		*Alias
	}{
		Source: GiftCodeType,
		Alias:  (*Alias)(&c),
	})
}

type ChatBoostSourceGiveaway struct {
	Source            string
	GiveawayMessageId string
	User              User
	IsUnclaimed       bool
}

func (c ChatBoostSourceGiveaway) MarshalJSON() ([]byte, error) {
	type Alias ChatBoostSourceGiveaway
	return json.Marshal(&struct {
		Source string
		*Alias
	}{
		Source: GiveawayType,
		Alias:  (*Alias)(&c),
	})
}

type ChatBoostSourcePremium struct {
	Source string
	User   User
}

func (c ChatBoostSourcePremium) MarshalJSON() ([]byte, error) {
	type Alias ChatBoostSourcePremium
	return json.Marshal(&struct {
		Source string
		*Alias
	}{
		Source: PremiumType,
		Alias:  (*Alias)(&c),
	})
}

type ChatBoostUpdated struct {
	Chat  Chat
	Boost ChatBoost
}

type UserChatBoosts struct {
	Boosts []ChatBoost
}
