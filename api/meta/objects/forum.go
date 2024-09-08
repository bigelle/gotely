package objects

type ForumTopic struct {
	MessageThreadId   int
	Name              string
	IconColor         int
	IconCustomEmojiId string
}

// placeholder for event
type ForumTopicClosed struct{}

type ForumTopicCreated struct {
	Name              string
	IconColor         int
	IconCustomEmojiId string
}

type ForumTopicEdited struct {
	Name              string
	IconCustomEmojiId string
}

// placeholder for event
type ForumTopicReopened struct{}

// placeholder for event
type GeneralForumTopicHidden struct{}

// placeholder for event
type GeneralForumTopicUnhidden struct{}
