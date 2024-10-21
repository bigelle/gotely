package types

type Giveaway struct {
	Chats                         []Chat    `json:"chats"`
	WinnerSelectionDate           int       `json:"winner_selection_date"`
	WinnerCount                   int       `json:"winner_count"`
	OnlyNewMembers                *bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              *bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              *string   `json:"prize_description,omitempty"`
	CountryCodes                  *[]string `json:"country_codes,omitempty"`
	PrizeStarCount                *int      `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      *bool    `json:"is_star_giveaway,omitempty"`
}

type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`
	GiveawayMessageId             int     `json:"giveaway_message_id"`
	WinnersSelectionDate          int     `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []User  `json:"winners"`
	AdditionalChatCount           *int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                *int    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`
	PrizeDescription              *string `json:"prize_description,omitempty"`
}
