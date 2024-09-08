package objects

type Giveaway struct {
	Chats                         []Chat
	WinnerSelectionDate           int
	WinnerCount                   int
	OnlyNewMembers                bool
	HasPublicWinners              bool
	PrizeDescription              string
	CountryCodes                  []string
	PremiumSubscriptionMonthCount int
}

type GiveawayCompleted struct {
	WinnerCount         int
	UnclaimedPrizeCount int
	GiveawayMessage     Message
}

// event, placeholder
type GiveawayCreated struct{}

type GiveawayWinners struct {
	Chat                          Chat
	GiveawayMessageId             int
	WinnerSelectionDate           int
	WinnerCount                   int
	Winners                       []User
	AdditionalChatCount           int
	PremiumSubscriptionMonthCount int
	UnclaimedPrizeCount           int
	OnlyNewMembers                bool
	WasRefunded                   bool
	PrizeDescription              string
}
