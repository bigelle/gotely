package types

import (
	"encoding/json"
	"fmt"
)

type ChatMember struct {
	ChatMemberInterface
}

type ChatMemberInterface interface {
	chatMemberContract()
}

func (c *ChatMember) UnmarshalJSON(data []byte) error {
	var raw struct {
		Status     string `json:"status"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Status {
	case "administrator":
		c.ChatMemberInterface = new(ChatMemberAdministrator)
	case "member":
		c.ChatMemberInterface = new(ChatMemberMember)
	case "owner":
		c.ChatMemberInterface = new(ChatMemberOwner)
	case "restricted":
		c.ChatMemberInterface = new(ChatMemberRestricted)
	case "banned":
		c.ChatMemberInterface = new(ChatMemberBanned)
	case "left":
		c.ChatMemberInterface = new(ChatMemberLeft)
	case "updated":
		c.ChatMemberInterface = new(ChatMemberUpdated)
	default:
		return fmt.Errorf(
			"Status must be administrator, member, owner, restricted, banned, left or updated",
		)
	}

	return json.Unmarshal(raw.Attributes, c.ChatMemberInterface)
}

type ChatMemberAdministrator struct {
	Status      string `json:"status"`
	User        User   `json:"user"`
	CanBeEdited bool   `json:"can_be_edited"`
	ChatAdministratorRights
	CustomTitle *string `json:"custom_title,omitempty"`
}

func (c ChatMemberAdministrator) chatMemberContract() {}

type ChatMemberMember struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate *int   `json:"until_date,omitempty"`
}

func (c ChatMemberMember) chatMemberContract() {}

type ChatMemberOwner struct {
	Status      string  `json:"status"`
	User        *User   `json:"user"`
	IsAnonymous bool    `json:"is_anonymous"`
	CustomTitle *string `json:"custom_title,omitempty"`
}

func (c ChatMemberOwner) chatMemberContract() {}

type ChatMemberRestricted struct {
	Status   string `json:"status"`
	User     *User  `json:"user"`
	IsMember bool   `json:"is_member"`
	ChatPermissions
	UntilDate int `json:"until_date"`
}

func (c ChatMemberRestricted) chatMemberContract() {}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate int    `json:"until_date"`
}

func (c ChatMemberBanned) chatMemberContract() {}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

func (c ChatMemberLeft) chatMemberContract() {}

type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          *bool           `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"`
}

func (c ChatMemberUpdated) chatMemberContract() {}
