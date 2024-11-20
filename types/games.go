package types

type CallbackGame struct {
}

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []PhotoSize      `json:"photo"`
	Text         *string          `json:"text,omitempty"`
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}
