package objects

type Game struct {
	title       string
	description string
	photo       []PhotoSize
	text        string
	entities    []MessageEntity
	animation   Animation
}

func (g Game) HasEntities() bool {
	return g.entities != nil && len(g.entities) != 0
}

// placeholder
type CallbackGame struct {
}

// placeholder
func (g CallbackGame) Validate() error {
	return nil
}
