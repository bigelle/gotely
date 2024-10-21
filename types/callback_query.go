package types

type CallbackQuery struct {
	Id              string                    `json:"id"`
	From            User                      `json:"from"`
	ChatInstance    string                    `json:"chat_instance"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageId *string                   `json:"inline_message_id,omitempty"`
	Data            *string                   `json:"data,omitempty"`
	GameShortName   *string                   `json:"game_short_name,omitempty"`
}
