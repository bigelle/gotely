package longpolling

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/bot"
)

// LongPollingBot represents a Telegram bot using long polling.
//
// Default configuration (without options):
//
//   - HTTP client: http.DefaultClient
//
//   - API URL: "https://api.telegram.org/bot%s/%s"
//     (first placeholder is the bot token, second is the API method)
//
//   - Update limit: 100
//
//   - Timeout: 30s
//
//   - Logger: slog.Default()
//
//   - Allowed updates: nil (use Telegram defaults or previously used settings)
//
// The bot requires an update handler (`OnUpdateFunc`) to process incoming updates.
// Middlewares can be added using `Use()`.
type LongPollingBot struct {
	//IMPORTANT:

	//Telegram bot API access token
	token string
	//A function which is called on every incoming update
	onUpdate bot.OnUpdateFunc

	//CONFIGURABLE:

	//A list of functions that will be called before `OnUpdate`.
	// Good spot for additional logging, caching, etc.
	middleWare []bot.MiddlewareFunc
	//A client that will be used when making any API request.
	//Defaults to default http.Client
	client *http.Client
	apiUrl string
	//Limits the number of updates to be retrieved. Values between 1-100 are accepted.
	//Defaults to 100.
	limit int
	//Timeout in seconds for long polling. Should be positive, short polling should be used for testing purposes only.
	//Defaults to 30
	timeout int
	//A list of the update types you want your bot to receive.
	//See https://core.telegram.org/bots/api#update for a complete list of available update types.
	//Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default).
	//If not specified, the previous setting will be used.
	allowedUpdates *[]string
	//Logger that will be used to display information about any incoming updates, errors, responses, etc
	logger slog.Logger

	//AUTO:

	//will be used to asynchronously respond to every incoming update
	chContext chan *bot.Context
	//calculated automatically, used to poll for new updates
	offset *int
	//cancel func to gracefully stop go-routines
	cancel context.CancelFunc
}

func New(token string, onUpdate bot.OnUpdateFunc, opts ...Option) *LongPollingBot {
	bot := LongPollingBot{
		token:    token,
		onUpdate: onUpdate,
		//no middleware
		client:  http.DefaultClient,
		apiUrl:  "https://api.telegram.org/bot%s/%s",
		limit:   100,
		timeout: 30,
		//all updates are allowed
		logger:    *slog.Default(),
		chContext: make(chan *bot.Context),
		//offset initially not specified
	}

	for _, opt := range opts {
		opt(&bot)
	}

	return &bot
}

type Option func(*LongPollingBot)

func (l *LongPollingBot) Use(mw ...bot.MiddlewareFunc) {
	l.middleWare = append(l.middleWare, mw...)
}

func WithClient(c *http.Client) Option {
	return func(lpb *LongPollingBot) {
		lpb.client = c
	}
}

func WithUrl(url string) Option {
	return func(lpb *LongPollingBot) {
		lpb.apiUrl = url
	}
}

func WithLimit(l int) Option {
	return func(lpb *LongPollingBot) {
		lpb.limit = l
	}
}

func WithTimeout(t int) Option {
	return func(lpb *LongPollingBot) {
		lpb.timeout = t
	}
}

func WithAllowedUpdates(upds *[]string) Option {
	return func(lpb *LongPollingBot) {
		lpb.allowedUpdates = upds
	}
}

func WithLogger(l slog.Logger) Option {
	return func(lpb *LongPollingBot) {
		lpb.logger = l
	}
}

func (l *LongPollingBot) Start() {
	if err := l.Validate(); err != nil {
		panic(fmt.Errorf("can't start bot because it failed the validation: %w", err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	l.cancel = cancel

	wg := &sync.WaitGroup{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		l.logger.Info("shutting down...")
		l.Stop()
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.poll(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.respond(ctx)
	}()

	l.logger.Info("bot started")
	wg.Wait()
	l.logger.Info("bot stopped")
}

func (l *LongPollingBot) Validate() error {
	//FIXME
	return nil
}

func (l *LongPollingBot) poll(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			l.logger.Info("bot stopped, exiting polling loop")
			return
		default:
			upds, err := l.getUpdates()
			if err != nil {
				l.logger.Error("failed to get new updates", "error", err)
				continue
			}
			if len(*upds) > 0 {
				for _, upd := range *upds {
					select {
					case l.chContext <- &bot.Context{
						Token:  l.token,
						Update: upd,
						Client: l.client,
						ApiUrl: l.apiUrl,
					}:
						l.logger.Info("new incoming update", "update_id", upd.UpdateId)
						newOffset := (*upds)[len(*upds)-1].UpdateId + 1
						l.offset = &newOffset
					case <-ctx.Done():
						l.logger.Info("bot stopped, exiting polling loop")
						return
					}
				}
			}
		}
	}
}

func (l *LongPollingBot) getUpdates() (*[]objects.Update, error) {
	b := getUpdates{
		Offset:         l.offset,
		Timeout:        &l.timeout,
		Limit:          &l.limit,
		AllowedUpdates: l.allowedUpdates,
	}
	return gotely.SendGetRequestWith[[]objects.Update](
		b,
		l.token,
		gotely.WithClient(l.client),
		gotely.WithUrl(l.apiUrl),
	)
}

type getUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g getUpdates) Endpoint() string {
	return "getUpdates"
}

func (g getUpdates) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return fmt.Errorf("limit must be between 1 and 100")
		}
	}
	if g.Timeout != nil {
		if *g.Timeout < 0 {
			return fmt.Errorf("timeout must be positive")
		}
	}
	//FIXME allowed updates validation
	return nil
}

func (g getUpdates) Reader() (io.Reader, error) {
	b, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (g getUpdates) ContentType() string {
	return "application/json"
}

func (l *LongPollingBot) respond(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			l.logger.Info("bot stopped, exiting responding loop")
			return
		case upd := <-l.chContext:
			//FIXME: middleware
			// for i, mw := range l.middleWare {
			// 	if err := mw(*upd); err != nil {
			// 		//							telling which middleware failed, 0 based
			// 		l.logger.Error("failed middleware", "middleware", i, "error", err)
			// 	}
			// }

			if err := l.onUpdate(*upd); err != nil {
				l.logger.Error("failed to respond to an update", "update_id", upd.Update.UpdateId, "error", err)
				continue
			}
			l.logger.Info("done responding to update", "update_id", upd.Update.UpdateId)
		}
	}
}

func (l *LongPollingBot) Stop() {
	if l.cancel != nil {
		l.cancel()
	}
}
