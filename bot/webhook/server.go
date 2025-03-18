package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/bot"
)

// WebhookBot is an `http.Server` designed to handle updates
// arriving at the "/webhook" WebhookEndpoint.
type WebhookBot struct {
	//configurable:

	http.Server
	// Used to send Telegram Bot API requests.
	Client *http.Client
	// The endpoint used when registering an `http.HandlerFunc` to handle new updates.
	// Defaults to "/webhook".
	WebhookEndpoint string
	// The Telegram Bot API URL. You can replace this value if running the Bot API locally.
	ApiUrl string
	// The Telegram Bot API token.
	Token string
	// A function called for every incoming update.
	OnUpdate bot.OnUpdateFunc
	// A list of middleware functions executed on every call to OnUpdate.
	// Defaults to a list containing only the recovery middleware.
	Middleware []bot.MiddlewareFunc

	//sending in SetWebhook:

	//HTTPS URL to send updates to. Use an empty string to remove webhook integration
	WebhookUrl string `json:"url"`
	//Upload your public key certificate so that the root certificate in use can be checked.
	//See https://core.telegram.org/bots/self-signed for details.
	Certificate objects.InputFile `json:"certificate,omitempty"`
	//The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	IpAddress *string `json:"ip_address,omitempty"`
	//The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	//Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections *int `json:"max_connections,omitempty"`
	//A list of the update types you want your bot to receive.
	// For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See https://core.telegram.org/bots/api#update for a complete list of available update types.
	// Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
	//Pass True to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
	//A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters.
	// Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
	SecretToken *string `json:"secret_token,omitempty"`

	//service

	ctx    context.Context
	cancel context.CancelFunc
}

// New creates a WebhookBot with the given Webhook URL and Telegram Bot API token.
func New(url, token string, onUpdate bot.OnUpdateFunc, opts ...Option) *WebhookBot {
	bot := WebhookBot{
		Server:          http.Server{Addr: ":80"},
		Client:          http.DefaultClient,
		WebhookEndpoint: "/webhook",
		ApiUrl:          "https://api.telegram.org/bot%s/%s", //FIXME add support for <token> <method> placeholders
		Token:           token,
		OnUpdate:        onUpdate,
		Middleware:      []bot.MiddlewareFunc{recoveryMiddleware},
		WebhookUrl:      url,
		//no cert by default
		//no ip address
		//not specifying max connections
		// not specifying allowed updates
		//not dropping pending updates
		//no secret token
	}

	for _, opt := range opts {
		opt(&bot)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(bot.WebhookEndpoint, bot.handleUpdate)
	bot.Handler = mux

	return &bot
}

func recoveryMiddleware(next bot.OnUpdateFunc) bot.OnUpdateFunc {
	return func(ctx bot.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
				err = fmt.Errorf("internal error")
			}
		}()
		return next(ctx)
	}
}

func (whb *WebhookBot) handleUpdate(w http.ResponseWriter, r *http.Request) {
	var upd objects.Update
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ctx := bot.Context{
		Token:  whb.Token,
		Update: upd,
		ApiUrl: whb.ApiUrl,
		Client: whb.Client,
	}

	handler := whb.OnUpdate
	for i := len(whb.Middleware) - 1; i >= 0; i-- {
		handler = whb.Middleware[i](handler)
	}

	if err := handler(ctx); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Start sends a SetWebhook request to the Telegram Bot API.
// On success, it begins listening and serving with the current configuration.
func (whb *WebhookBot) Start() {
	var result bool
	err := gotely.SendRequestWith(
		&SetWebhook{
			Url:                whb.WebhookUrl + whb.WebhookEndpoint,
			Certificate:        whb.Certificate,
			AllowedUpdates:     whb.AllowedUpdates,
			MaxConnections:     whb.MaxConnections,
			IpAddress:          whb.IpAddress,
			DropPendingUpdates: whb.DropPendingUpdates,
			SecretToken:        whb.SecretToken,
		},
		result,
		whb.Token,
		gotely.WithClient(whb.Client),
		gotely.WithUrl(whb.ApiUrl),
	)
	if err != nil {
		log.Fatalf("unable to set up webhook: %s", err.Error())
	}
	whb.ctx, whb.cancel = context.WithCancel(context.Background())

	serverErr := make(chan error, 1)

	go func() {
		if err := whb.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
		close(serverErr)
	}()

	select {
	case <-whb.ctx.Done():
		log.Println("Shutting down server...")
		whb.Stop()
	case err := <-serverErr:
		if err != nil {
			log.Fatalf("server error: %s", err)
		}
	}
}

// Stop gracefully shuts down the currently running WebhookBot server.
func (whb *WebhookBot) Stop() {
	if whb.cancel != nil {
		whb.cancel()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := whb.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %s", err.Error())
	}
}

// Use appends middleware to the current WebhookBot
func (wb *WebhookBot) Use(middleware ...bot.MiddlewareFunc) {
	wb.Middleware = append(wb.Middleware, middleware...)
}

type Option func(*WebhookBot)

// WithServerConfig is a functional option used to configure
// the `http.Server` of the WebhookBot.
func WithServerConfig(s *http.Server) Option {
	return func(wb *WebhookBot) {
		wb.Addr = s.Addr
		wb.ReadTimeout = s.ReadTimeout
		wb.WriteTimeout = s.WriteTimeout
		wb.IdleTimeout = s.IdleTimeout
		wb.MaxHeaderBytes = s.MaxHeaderBytes
		wb.TLSConfig = s.TLSConfig
		wb.TLSNextProto = s.TLSNextProto
		wb.ConnContext = s.ConnContext
		wb.ConnState = s.ConnState
		wb.ErrorLog = s.ErrorLog
		wb.Handler = s.Handler
	}
}

// WithWebhookEndpoint is a functional option used to configure
// the endpoint for handling webhooks.
func WithWebhookEndpoint(e string) Option {
	return func(wb *WebhookBot) {
		wb.WebhookEndpoint = e
	}
}

// WithCertificate is a functional option used to
// set the HTTPS certificate.
func WithCertificate(c objects.InputFile) Option {
	return func(wb *WebhookBot) {
		wb.Certificate = c
	}
}

// WithIpAddress is a functional option used to
// set the IP address for sending webhook requests
// instead of the DNS-resolved address.
func WithIpAddress(a string) Option {
	return func(wb *WebhookBot) {
		wb.IpAddress = &a
	}
}

// WithMaxConnections is a functional option used to
// set the maximum number of simultaneous HTTPS connections
// to the webhook for update delivery.
func WithMaxConnections(c int) Option {
	return func(wb *WebhookBot) {
		wb.MaxConnections = &c
	}
}

// WithAllowedUpdates is a functional option used to
// set the list of updates the bot will receive through the webhook.
func WithAllowedUpdates(upds *[]string) Option {
	return func(wb *WebhookBot) {
		wb.AllowedUpdates = upds
	}
}

// WithDropPendingUpdates is a functional option used to
// determine whether the SetWebhook method should drop
// all currently pending updates.
func WithDropPendingUpdates(d bool) Option {
	return func(wb *WebhookBot) {
		wb.DropPendingUpdates = &d
	}
}

// WithSecretToken is a functional option used to
// set the secret token sent through the webhook
// inside the "X-Telegram-Bot-Api-Secret-Token" header.
// It can be used to ensure that the webhook was set by you.
func WithSecretToken(t string) Option {
	return func(wb *WebhookBot) {
		wb.SecretToken = &t
	}
}
