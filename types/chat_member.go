package types

import (
	"encoding/json"
	"errors"
)

type ChatMember struct {
	ChatMemberInterface
}

type ChatMemberInterface interface {
	chatMemberContract()
}

func (c ChatMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ChatMemberInterface)
}

func (c *ChatMember) UnmarshalJSON(data []byte) error {
	var raw struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Status {
	case "administrator":
		tmp := ChatMemberAdministrator{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "member":
		tmp := ChatMemberMember{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "owner":
		tmp := ChatMemberOwner{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "restricted":
		tmp := ChatMemberRestricted{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "banned":
		tmp := ChatMemberBanned{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "left":
		tmp := ChatMemberLeft{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	case "updated":
		tmp := ChatMemberUpdated{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		c.ChatMemberInterface = tmp
	default:
		return errors.New(
			"status must be administrator, member, owner, restricted, banned, left or updated",
		)
	}

	return nil
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
