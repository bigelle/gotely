package objects

type Invoice struct {
	Title          string
	Description    string
	StartParameter string
	Currency       string
	TotalAmount    string
	Photo          PhotoSize
}
