package types

type User struct {
	Id                      int64   `json:"id"`
	FirstName               string  `json:"first_name"`
	IsBot                   bool    `json:"is_bot"`
	LastName                *string `json:"last_name,omitempty"`
	UserName                *string `json:"user_name,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportInlineQueries    *bool   `json:"support_inline_queries,omitempty"`
	IsPremium               *bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Plotos     [][]PhotoSize `json:"plotos"`
}

type SharedUser struct {
	UserId    int64        `json:"user_id"`
	FirstName *string      `json:"first_name,omitempty"`
	LastName  *string      `json:"last_name,omitempty"`
	Username  *string      `json:"username,omitempty"`
	Photo     *[]PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestId string       `json:"request_id"`
	Users     []SharedUser `json:"users"`
}
