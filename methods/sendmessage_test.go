package methods

import (
	"fmt"
	"testing"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/longpolling"
	"github.com/bigelle/tele.go/types"
)

func TestSendMessage(t *testing.T) {
	OnUpdate := func(upd types.Update) error {
		if upd.Message != nil {
			chatId := upd.Message.Chat.Id
			txt := upd.Message.Text

			msg, err := SendMessage[int]{
				ChatId: int(chatId),
				Text:   *txt,
			}.Execute()

			if err != nil {
				return err
			}
			fmt.Println(*msg.Text)
		}
		return nil
	}
	bot, err := telego.NewBot("6470269136:AAG_UUjSJhbH89AIabYy4hp4n3hxAjB1sGs", OnUpdate)
	if err != nil {
		panic(err)
	}
	err = longpolling.Connect(bot)
	if err != nil {
		panic(err)
	}
}
