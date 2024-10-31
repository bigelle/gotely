package telego

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Bot struct {
	Token  string
	ApiUrl string
}

var (
	bot  Bot
	once sync.Once
)

const defaultApiUrl = "https://api.telegram.org/bot"

func GetBot() *Bot {
	once.Do(func() {
		bot = Bot{
			ApiUrl: defaultApiUrl,
		}
	})
	return &bot
}

func (b *Bot) SetToken(t string) *Bot {
	b.Token = t
	return b
}

func (b *Bot) SetApiUrl(url string) *Bot {
	b.ApiUrl = url
	return b
}

func (b Bot) MakeRequest(reqMethod string, apiMethod string, body []byte) ([]byte, error){
	// TODO: probably bot validating method
	if b.Token == "" {
		return nil, errors.New("API token can't be empty")
	}
	url := fmt.Sprintf("%s%s/%s", bot.ApiUrl, bot.Token, apiMethod)
	req, err := http.NewRequest(reqMethod, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// sending request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// reading response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (b Bot) MakePostRequest(method string, body []byte) ([]byte, error) {
	return b.MakeRequest("POST", method, body)
}

func (b Bot) MakeGetRequest(method string, body []byte)([]byte, error){
	return b.MakeRequest("GET", method, body)
}
