package objects

type ChatShared struct {
	RequestId string
	ChatId    int64
	Title     string
	Username  string
	Photo     []PhotoSize
}
