package methods

import (
	"os"
	"testing"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/longpolling"
	"github.com/bigelle/tele.go/types"
	"github.com/joho/godotenv"
)

func TestSendMessage(t *testing.T) {
	b, err := telego.NewBot(os.Getenv("BOT_TOKEN"), func(u types.Update) error { return nil })
	if err != nil {
		t.Fatal(err)
	}
	longpolling.Connect(b)
	sm := SendMessage[int]{}.New(446182219, "DEEZ NUTS")
	m, err := sm.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", m)
}

func init() {
	godotenv.Load("../.env")
}
