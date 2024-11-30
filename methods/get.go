package methods

import (
	"encoding/json"
	"fmt"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
	iso6391 "github.com/emvi/iso-639-1"
)

type GetUserProfilePhotos struct {
	UserId int
	Offset *int
	Limit  *int
}

func (g GetUserProfilePhotos) Validate() error {
	if g.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return errors.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

func (g GetUserProfilePhotos) MarshalJSON() ([]byte, error) {
	type alias GetUserProfilePhotos
	return json.Marshal(alias(g))
}

func (g GetUserProfilePhotos) Execute() (*types.UserProfilePhotos, error) {
	return internal.MakeGetRequest[types.UserProfilePhotos](telego.GetToken(), "getUserProfilePhotos", g)
}

type GetFile struct {
	FileId string
}

func (g GetFile) Validate() error {
	if strings.TrimSpace(g.FileId) == "" {
		return errors.ErrInvalidParam("file_id parameter can't be empty")
	}
	return nil
}

func (g GetFile) MarshalJSON() ([]byte, error) {
	type alias GetFile
	return json.Marshal(alias(g))
}

func (g GetFile) Execute() (*types.File, error) {
	return internal.MakeGetRequest[types.File](telego.GetToken(), "getFile", g)
}

type GetChat[T ChatId] struct {
	ChatId T
}

func (p GetChat[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p GetChat[T]) MarshalJSON() ([]byte, error) {
	type alias GetChat[T]
	return json.Marshal(alias(p))
}

func (p GetChat[T]) Execute() (*types.ChatFullInfo, error) {
	return internal.MakeGetRequest[types.ChatFullInfo](telego.GetToken(), "getChat", p)
}

type GetChatAdministrators[T ChatId] struct {
	ChatId T
}

func (p GetChatAdministrators[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p GetChatAdministrators[T]) MarshalJSON() ([]byte, error) {
	type alias GetChatAdministrators[T]
	return json.Marshal(alias(p))
}

func (p GetChatAdministrators[T]) Execute() (*[]types.ChatMember, error) {
	return internal.MakeGetRequest[[]types.ChatMember](telego.GetToken(), "getChatAdministrators", p)
}

type GetChatMemberCount[T ChatId] struct {
	ChatId T
}

func (p GetChatMemberCount[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p GetChatMemberCount[T]) MarshalJSON() ([]byte, error) {
	type alias GetChatMemberCount[T]
	return json.Marshal(alias(p))
}

func (p GetChatMemberCount[T]) Execute() (*int, error) {
	return internal.MakeGetRequest[int](telego.GetToken(), "getChatMemberCount", p)
}

type GetChatMember[T ChatId] struct {
	ChatId T
	UserId int
}

func (p GetChatMember[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (p GetChatMember[T]) MarshalJSON() ([]byte, error) {
	type alias GetChatMember[T]
	return json.Marshal(alias(p))
}

func (p GetChatMember[T]) Execute() (*types.ChatMember, error) {
	return internal.MakeGetRequest[types.ChatMember](telego.GetToken(), "getChatMember", p)
}

type GetForumTopicIconStickers struct {
}

// always nil
func (g GetForumTopicIconStickers) Validate() error {
	return nil
}

// alwways empty json
func (g GetForumTopicIconStickers) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{}{})
}

func (g GetForumTopicIconStickers) Execute() (*[]types.Sticker, error) {
	return internal.MakeGetRequest[[]types.Sticker](telego.GetToken(), "getForumTopicStickers", g)
}

type GetUserChatBoosts[T ChatId] struct {
	ChatId T
	UserId int
}

func (g GetUserChatBoosts[T]) Validate() error {
	if c, ok := any(g.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(g.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if g.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (g GetUserChatBoosts[T]) MarshalJSON() ([]byte, error) {
	type alias GetUserChatBoosts[T]
	return json.Marshal(alias(g))
}

func (g GetUserChatBoosts[T]) Execute() (*types.UserChatBoosts, error) {
	return internal.MakeGetRequest[types.UserChatBoosts](telego.GetToken(), "getUserChatBoosts", g)
}

type GetBusinessConnection struct {
	BusinessConnectionId string
}

func (g GetBusinessConnection) Validate() error {
	if strings.TrimSpace(g.BusinessConnectionId) == "" {
		return errors.ErrInvalidParam("business_connection_id parameter can't be empty")
	}
	return nil
}

func (g GetBusinessConnection) MarshalJSON() ([]byte, error) {
	type alias GetBusinessConnection
	return json.Marshal(alias(g))
}

func (g GetBusinessConnection) Execute() (*types.BusinessConnection, error) {
	return internal.MakeGetRequest[types.BusinessConnection](telego.GetToken(), "getBusinessConnection", g)
}

type GetMyCommands struct {
	Scope        *types.BotCommandScope
	LanguageCode *string
}

func (s GetMyCommands) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyCommands) MarshalJSON() ([]byte, error) {
	type alias DeleteMyCommands
	return json.Marshal(alias(s))
}

func (s GetMyCommands) Execute() (*[]types.BotCommand, error) {
	return internal.MakeGetRequest[[]types.BotCommand](telego.GetToken(), "getMyCommands", s)
}

type GetMyName struct {
	LanguageCode *string
}

func (s GetMyName) Validate() error {

	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyName) MarshalJSON() ([]byte, error) {
	type alias GetMyName
	return json.Marshal(alias(s))
}

func (s GetMyName) Execute() (*types.BotName, error) {
	return internal.MakeGetRequest[types.BotName](telego.GetToken(), "getMyName", s)
}

type GetMyDescription struct {
	LanguageCode *string
}

func (s GetMyDescription) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyDescription) MarshalJSON() ([]byte, error) {
	type alias GetMyDescription
	return json.Marshal(alias(s))
}

func (s GetMyDescription) Execute() (*types.BotDescription, error) {
	return internal.MakeGetRequest[types.BotDescription](telego.GetToken(), "getMyDescription", s)
}

type GetMyShortDescription struct {
	LanguageCode *string
}

func (s GetMyShortDescription) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyShortDescription) MarshalJSON() ([]byte, error) {
	type alias GetMyShortDescription
	return json.Marshal(alias(s))
}

func (s GetMyShortDescription) Execute() (*bool, error) {
	return internal.MakeGetRequest[bool](telego.GetToken(), "getMyShortDescription", s)
}

type GetChatMenuButton struct {
	ChatId *int
}

func (s GetChatMenuButton) Validate() error {
	if s.ChatId != nil {
		if *s.ChatId < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (s GetChatMenuButton) MarshalJSON() ([]byte, error) {
	type alias GetChatMenuButton
	return json.Marshal(alias(s))
}

func (s GetChatMenuButton) Execute() (*types.MenuButton, error) {
	return internal.MakeGetRequest[types.MenuButton](telego.GetToken(), "setChatMenuButton", s)
}

type GetMyDefaultAdministratorRights struct {
	ForChannels *bool
}

// always nil
func (s GetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s GetMyDefaultAdministratorRights) MarshalJSON() ([]byte, error) {
	type alias GetMyDefaultAdministratorRights
	return json.Marshal(alias(s))
}

func (s GetMyDefaultAdministratorRights) Execute() (*types.ChatAdministratorRights, error) {
	return internal.MakePostRequest[types.ChatAdministratorRights](telego.GetToken(), "getMyDefaultAdministratorRights", s)
}
