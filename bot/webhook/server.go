package webhook

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/bot"
)

// TODO: simple default server
type WebhookBot struct {
	http.Server
	Client     *http.Client
	Endpoint   string
	WebhookUrl string
	ApiUrl     string
	Token      string
	OnUpdate   bot.OnUpdateFunc
	Middleware []bot.MiddlewareFunc
	port       int
}

func New(addr, token string, onUpdate bot.OnUpdateFunc, opts ...Option) *WebhookBot {
	bot := WebhookBot{
		Server:   http.Server{Addr: addr},
		Endpoint: "/webhook",
		Token:    token,
		OnUpdate: onUpdate,
		// no middleware by default
		port: 80,
		//no cert by default
	}
	for _, opt := range opts {
		opt(&bot)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(bot.Endpoint, bot.handleUpdate)
	bot.Handler = mux

	return &bot
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
	if err := whb.OnUpdate(ctx); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
		//FIXME check error type
	}
	w.WriteHeader(http.StatusOK)
}

func (whb *WebhookBot) Start() {
	go func() {
		if whb.WebhookUrl != "" {
			if err := SetWebhook(whb.Token, whb.WebhookUrl); err != nil {
				log.Fatalf("failed to set webhook: %s", err.Error())
			}
		}
		if err := whb.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("internal error: %s", err.Error())
		}
	}()
}

func (whb *WebhookBot) Stop(ctx context.Context) {
	if err := whb.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %s", err.Error())
	}
}

type Option func(*WebhookBot)

//TODO options
