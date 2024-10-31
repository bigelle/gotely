package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/methods"
)

var bot *telego.Bot

func TestSendMessage(t *testing.T) {
	msg, err := methods.SendMessage[int]{}.New(446182219, "deez nuts").Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", msg)
}

func TestGetUpdates(t *testing.T) {
	upds, err := methods.GetUpdates{}.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", upds)
}

func init() {
	godotenv.Load("../.env")
	token := os.Getenv("BOT_TOKEN")
	bot = telego.GetBot()
	bot.SetToken(token)
	fmt.Printf("running with token: %s\n", token)
}
