package test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bigelle/tele.go/types"
)

func TestDeserializeMenuButton(t *testing.T) {
	var testcases = []struct {
		CaseName  string
		Case      types.MenuButton
		ExpectErr bool
	}{
		{
			CaseName: "Commands",
			Case: types.MenuButton{
				MenuButtonInterface: types.MenuButtonCommands{Type: "commands"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "WebApp",
			Case: types.MenuButton{
				MenuButtonInterface: types.MenuButtonWebApp{Type: "web_app"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "Default",
			Case: types.MenuButton{
				MenuButtonInterface: types.MenuButtonDefault{Type: "default"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "UnknownType",
			Case: types.MenuButton{
				MenuButtonInterface: types.MenuButtonDefault{Type: "UnknownType"},
			},
			ExpectErr: true,
		},
	}

	for _, c := range testcases {
		t.Run(c.CaseName, func(t *testing.T) {
			original := c.Case

			data, err := json.Marshal(original)
			if err != nil && !c.ExpectErr {
				t.Fatalf("marshal error: %v", err)
			}

			var result types.MenuButton
			if err := json.Unmarshal(data, &result); err != nil && !c.ExpectErr {
				t.Fatalf("unmarshal error: %v, %s", err, data)
			}

			if !reflect.DeepEqual(original, result) && !c.ExpectErr {
				t.Errorf("expected: %+v, got: %+v", original, result)
			}
		})
	}
}

func TestDeserializeMessageOrigin(t *testing.T) {
	var testcases = []struct {
		CaseName  string
		Case      types.MessageOrigin
		ExpectErr bool
	}{
		{
			CaseName: "User",
			Case: types.MessageOrigin{
				MessageOriginInterface: types.MessageOriginUser{Type: "user"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "Chat",
			Case: types.MessageOrigin{
				MessageOriginInterface: types.MessageOriginChat{Type: "chat"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "HiddenUser",
			Case: types.MessageOrigin{
				MessageOriginInterface: types.MessageOriginHiddenUser{Type: "hidden_user"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "Channel",
			Case: types.MessageOrigin{
				MessageOriginInterface: types.MessageOriginChannel{Type: "channel"},
			},
			ExpectErr: false,
		},
		{
			CaseName: "UnknownType",
			Case: types.MessageOrigin{
				MessageOriginInterface: types.MessageOriginUser{Type: "UnknownType"},
			},
			ExpectErr: true,
		},
	}

	for _, c := range testcases {
		t.Run(c.CaseName, func(t *testing.T) {
			original := c.Case

			data, err := json.Marshal(original)
			if err != nil && !c.ExpectErr {
				t.Fatal(err)
			}

			var result types.MessageOrigin
			if err = json.Unmarshal(data, &result); err != nil && !c.ExpectErr {
				t.Fatalf("unmarshal error: %v, %s", err, data)
			}

			if !reflect.DeepEqual(original, result) && !c.ExpectErr {
				t.Errorf("expected: %+v, got: %+v", original, result)
			}
		})
	}
}

func TestDeserealizeBotCommandScope(t *testing.T) {
	var testcases = []struct {
		CaseName  string
		Case      types.BotCommandScope
		ExpectErr bool
	}{
		{
			CaseName: "AllChatMembers",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeAllChatAdministrators{
					Type: "all_chat_administrators",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "AllGroupChats",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeAllGroupChats{
					Type: "all_group_chats",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "AllPrivateChats",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeAllPrivateChats{
					Type: "all_private_chats",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "Chat",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeChat{
					Type: "chat",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "ChatAdministrators",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeChatAdministrators{
					Type: "chat_administrators",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "ChatMember",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeChatMember{
					Type: "chat_member",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "Default",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeDefault{
					Type: "default",
				},
			},
			ExpectErr: false,
		},
		{
			CaseName: "UnknownType",
			Case: types.BotCommandScope{
				BotCommandScopeInterface: types.BotCommandScopeAllChatAdministrators{
					Type: "UnknownType",
				},
			},
			ExpectErr: true,
		},
	}

	for _, c := range testcases {
		t.Run(c.CaseName, func(t *testing.T) {
			original := c.Case

			data, err := json.Marshal(original)
			if err != nil && !c.ExpectErr {
				t.Fatal(err)
			}

			var result types.BotCommandScope
			if err := json.Unmarshal(data, &result); err != nil && !c.ExpectErr {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(original, result) && !c.ExpectErr {
				t.Errorf("expected %v, got %v", original, result)
			}
		})
	}
}

//TODO: more testcases
