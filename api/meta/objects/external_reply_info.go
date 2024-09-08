package objects

type ExternalReplyInfo struct {
	Origin             MessageOrigin
	Chat               Chat
	MessageId          int
	LinkPreviewOptions LinkPreviewOptions
	Animation          Animation
	Document           Document
	Photo              []PhotoSize
	Sticker            Sticker
	Story              Story
	Video              Video
	VideoNote          VideoNote
	Voice              Voice
	Contact            Contact
	Dice               Dice
	Game               Game
	Giveaway           Giveaway
	GiveawayWinners    GiveawayWinners
	Invoice            Invoice
	Poll               Poll
	Venue              Venue
	PaidMediaInfo      PaidMediaInfo
}
