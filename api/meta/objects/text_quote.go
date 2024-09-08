package objects

type TextQuote struct {
	Text     string
	Entities []MessageEntity
	Position int
	IsManual bool
}
