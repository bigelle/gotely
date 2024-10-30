package types

import "github.com/bigelle/tele.go/internal/assertions"

type ApiResponse[T any] struct {
	Ok          bool                `json:"ok"`
	ErrorCode   int                 `json:"error_code"`
	Description *string             `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      T                   `json:"result"`
}

type ResponseParameters struct {
	MigrateToChatId *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int   `json:"retry_after,omitempty"`
}

type ReplyParameters struct {
	MessageId                int              `json:"message_id"`
	ChatId                   *string          `json:"chat_id,omitempty"`
	AllowSendingWithoutReply *bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    *string          `json:"quote,omitempty"`
	QuoteParseMode           *string          `json:"quote_parse_mode,omitempty"`
	QuoteEntities            *[]MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            *int             `json:"quote_position,omitempty"`
}

func (r ReplyParameters) Validate() error {
	if err := assertions.ParamNotEmpty(*r.ChatId, "ChatId"); err != nil {
		return err
	}
	return nil
}
