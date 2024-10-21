package types

type TextQuote struct {
	Text     string           `json:"text"`
	Position int              `json:"position"`
	Entities *[]MessageEntity `json:"entities,omitempty"`
	IsManual *bool            `json:"is_manual,omitempty"`
}
