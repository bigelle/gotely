package types

type BusinessConnection struct {
	Id         string `json:"id"`
	User       *User  `json:"user"`
	UserChatId int64  `json:"user_chat_id"`
	Date       int    `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message *string  `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionId string `json:"business_connection_id"`
	Chat                 Chat   `json:"chat"`
	MessageIds           []int  `json:"message_ids"`
}

type BusinessOpeningHours struct {
	TimeZone     string                         `json:"time_zone"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}
