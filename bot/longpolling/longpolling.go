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
	// IMPORTANT:

	// Telegram Bot API access token
	token string
	// A function called on every incoming update
	onUpdate bot.OnUpdateFunc

	// CONFIGURABLE:

	// A list of middleware functions called before `OnUpdate`.
	// Useful for logging, caching, etc.
	middleware []bot.MiddlewareFunc
	// An HTTP client used for making API requests.
	// Defaults to http.DefaultClient
	client *http.Client
	// Telegram Bot API URL
	apiUrl string
	// Limits the number of updates to retrieve (1-100).
	// Defaults to 100.
	limit int
	// Timeout in seconds for long polling.
	// Should be positive; short polling is for testing purposes only.
	// Defaults to 30.
	timeout int
	// A list of update types the bot should receive.
	// See https://core.telegram.org/bots/api#update for available types.
	// An empty list receives all types except chat_member, message_reaction, and message_reaction_count (default).
	// If unspecified, the previous setting is used.
	allowedUpdates *[]string
	// Logger used for logging incoming updates, errors, and responses.
	logger slog.Logger

	// AUTO:

	// Used asynchronously to respond to incoming updates
	chContext chan *bot.Context
	// Automatically calculated offset for polling new updates
	offset *int
	// Cancel function for gracefully stopping goroutines
	cancel context.CancelFunc
}

// New creates a LongPollingBot with the given Telegram Bot API token
// and an update handler function (`OnUpdateFunc`).
//
// Additional configuration can be provided via functional options (`opts`).
func New(token string, onUpdate bot.OnUpdateFunc, opts ...Option) *LongPollingBot {
	bot := LongPollingBot{
		token:    token,
		onUpdate: onUpdate,
		// no middleware
		client:  http.DefaultClient,
		apiUrl:  "https://api.telegram.org/bot<token>/<method>",
		limit:   100,
		timeout: 30,
		// all updates are allowed
		logger:    *slog.Default(),
		chContext: make(chan *bot.Context),
		// offset initially not specified
	}

	for _, opt := range opts {
		opt(&bot)
	}

	return &bot
}

// Option defines a functional option for configuring a LongPollingBot.
type Option func(*LongPollingBot)

// Use adds middleware functions to the bot.
// Middleware functions are executed before the update handler (`OnUpdate`).
func (l *LongPollingBot) Use(mw ...bot.MiddlewareFunc) {
	l.middleware = append(l.middleware, mw...)
}

// WithClient sets a custom HTTP client for API requests.
func WithClient(c *http.Client) Option {
	return func(lpb *LongPollingBot) {
		lpb.client = c
	}
}

// WithUrl sets a custom Telegram Bot API URL.
func WithUrl(url string) Option {
	return func(lpb *LongPollingBot) {
		lpb.apiUrl = url
	}
}

// WithLimit sets the maximum number of updates to be retrieved per request (1-100).
func WithLimit(l int) Option {
	return func(lpb *LongPollingBot) {
		lpb.limit = l
	}
}

// WithTimeout sets the long polling timeout in seconds.
// A positive value is required; short polling should be used for testing only.
func WithTimeout(t int) Option {
	return func(lpb *LongPollingBot) {
		lpb.timeout = t
	}
}

// WithAllowedUpdates sets the list of update types the bot should receive.
// If nil, Telegram's default or previously used settings will be applied.
func WithAllowedUpdates(upds *[]string) Option {
	return func(lpb *LongPollingBot) {
		lpb.allowedUpdates = upds
	}
}

// WithLogger sets a custom logger for logging incoming updates, errors, and responses.
func WithLogger(l slog.Logger) Option {
	return func(lpb *LongPollingBot) {
		lpb.logger = l
	}
}

// Start begins polling for updates using the current configuration.
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
	// FIXME
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
	b := GetUpdates{
		Offset:         l.offset,
		Timeout:        &l.timeout,
		Limit:          &l.limit,
		AllowedUpdates: l.allowedUpdates,
	}
	var result []objects.Update
	err := gotely.SendRequestWith(
		b,
		&result,
		l.token,
		gotely.WithClient(l.client),
		gotely.WithUrl(l.apiUrl),
	)
	return &result, err
}

type GetUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g GetUpdates) Endpoint() string {
	return "getUpdates"
}

func (g GetUpdates) Validate() error {
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
	// FIXME allowed updates validation
	return nil
}

func (g GetUpdates) Reader() (io.Reader, error) {
	b, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (g GetUpdates) ContentType() string {
	return "application/json"
}

func (g GetUpdates) HttpMethod() string {
	return http.MethodGet
}

func (l *LongPollingBot) respond(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			l.logger.Info("bot stopped, exiting responding loop")
			return
		case upd := <-l.chContext:
			handler := l.onUpdate
			for i := len(l.middleware) - 1; i >= 0; i-- {
				handler = l.middleware[i](handler)
			}

			if err := handler(*upd); err != nil {
				l.logger.Error("failed to respond to an update", "update_id", upd.Update.UpdateId, "error", err)
				continue
			}
			l.logger.Info("done responding to update", "update_id", upd.Update.UpdateId)
		}
	}
}

// Stop gracefully shuts down the bot, stopping all active polling and background processes.
func (l *LongPollingBot) Stop() {
	if l.cancel != nil {
		l.cancel()
	}
}
