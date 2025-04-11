package webhook_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/methods"
	"github.com/bigelle/gotely/objects"
	"github.com/bigelle/gotely/tgbot"
	"github.com/bigelle/gotely/tgbot/webhook"
)

type fakeRoundTripper func(*http.Request) (*http.Response, error)

func (f fakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type TestNoErrBot struct {
	token string
	tgbot.DefaultBot
}

func (t TestNoErrBot) Token() string {
	return t.token
}

func (t TestNoErrBot) Client() *http.Client {
	return &http.Client{
		Transport: fakeRoundTripper(func(r *http.Request) (*http.Response, error) {
			msg := objects.Message{
				MessageId: 42,
				Date:      int(time.Now().Unix()),
			}
			b, err := json.Marshal(msg)
			if err != nil {
				return nil, err
			}
			resp := gotely.ApiResponse{
				Ok:     true,
				Result: json.RawMessage(b),
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(gotely.EncodeJSON(resp)),
			}, nil
		}),
	}
}

func (t TestNoErrBot) OnUpdate(upd objects.Update) error {
	if upd.Message != nil {
		id := upd.Message.From.Id
		text := upd.Message.Text
		err := gotely.SendRequestWith(
			methods.SendMessage{
				ChatId: fmt.Sprint(id),
				Text:   *text,
			},
			t.Token(),
			nil,
			gotely.WithClient(t.Client()),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestNoErr(t *testing.T) {
	bot := TestNoErrBot{
		token: "MOCK_TOKEN",
	}
	hook := webhook.New(bot)
	defer hook.Stop()
	go hook.Start()

	time.Sleep(3 * time.Second)

	resp, err := http.Post("http://localhost:8080/webhook", "application/json", bytes.NewBuffer([]byte(`{"update_id": 123, "message": {"text": "hello", "from":{"id":42}}}`)))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}
