package types

type ForumTopic struct {
	MessageThreadId   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

// placeholder for event
type ForumTopicClosed struct {
}

type ForumTopicCreated struct {
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicEdited struct {
	Name              *string `json:"name,omitempty"`
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

// placeholder for event
type ForumTopicReopened struct {
}

// placeholder for event
type GeneralForumTopicHidden struct {
}

// placeholder for event
type GeneralForumTopicUnhidden struct {
}
