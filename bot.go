package telego

import "sync"

type Bot struct {
	Token  string
	ApiUrl string
}

var (
	bot  Bot
	once sync.Once
)

const defaultApiUrl = "https://api.telegram.org/bot"

func GetBot() Bot {
	once.Do(func() {
		bot = Bot{
			ApiUrl: defaultApiUrl,
		}
	})
	return bot
}

func (b *Bot) SetToken(t string) *Bot {
	b.Token = t
	return b
}

func (b *Bot) SetApiUrl(url string) *Bot {
	b.ApiUrl = url
	return b
}
