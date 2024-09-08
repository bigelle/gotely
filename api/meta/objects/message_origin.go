package objects

import "encoding/json"

const (
	UserType       = "user"
	HiddenUserType = "hidden_user"
	ChatType       = "chat"
	ChannelType    = "channel"
)

type MessageOrigin interface {
	// TODO: possibly should have MarshalJSON method from json.Marshaler
}

type MessageOriginType struct {
	Type string `json:"type"` // constant
}

type MessageOriginChannel struct {
	MessageOriginType 
	Date            int
	Chat            Chat
	MessageId       int
	AuthorSignature string
}

// NOTE: need to test
func (m MessageOriginChannel) MarshalJSON() ([]byte, error) {
	type Alias MessageOriginChannel
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  ChannelType,
		Alias: (*Alias)(&m),
	})
}

type MessageOriginChat struct {
	MessageOriginType
	Date            int
	SenderChat      Chat
	AuthorSignature string
}

func (m MessageOriginChat) MarshalJSON() ([]byte, error) {
	type Alias MessageOriginChat
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  ChatType,
		Alias: (*Alias)(&m),
	})
}

type MessageOriginHiddenUser struct {
	MessageOriginType
	Date           int
	SenderUsername string
}

func (m MessageOriginHiddenUser) MarshalJSON() ([]byte, error) {
	type Alias MessageOriginHiddenUser
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  HiddenUserType,
		Alias: (*Alias)(&m),
	})
}

type MessageOriginUser struct {
	MessageOriginType
	Date           int
	SenderUsername string
}

func (m MessageOriginUser) MarshalJSON() ([]byte, error) {
	type Alias MessageOriginUser
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  UserType,
		Alias: (*Alias)(&m),
	})
}
