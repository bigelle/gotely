package webhook

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/tgbot"
)

type WebhookBot struct {
	Bot tgbot.Bot

	s          *http.Server
	path       string
	addr       string
	middleware []func(next http.Handler) http.Handler
	l          slog.Logger
}

func New(bot tgbot.Bot, opts ...Option) WebhookBot {
	whb := WebhookBot{
		Bot: bot,

		addr: ":8080",
		path: "/webhook",
		l:    *slog.Default(),
	}
	for _, opt := range opts {
		opt(&whb)
	}
	return whb
}

func (b *WebhookBot) Use(m ...func(http.Handler) http.Handler) {
	b.middleware = append(b.middleware, m...)
}

func (b *WebhookBot) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc(b.path, b.handleFunc)

	h := http.Handler(mux)
	for i := len(b.middleware) - 1; i >= 0; i-- {
		h = b.middleware[i](h)
	}

	b.s = &http.Server{
		Addr:    b.addr,
		Handler: h,
	}
	b.l.Info("webhook server is listening and serving on", "addr", b.addr)
	return b.s.ListenAndServe()
}

func (b *WebhookBot) handleFunc(w http.ResponseWriter, r *http.Request) {
	var upd objects.Update
	if err := gotely.DecodeJSON(r.Body, &upd); err != nil {
		b.l.Error("can't read JSON", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := b.Bot.OnUpdate(upd)
	if err != nil {
		var validationErr gotely.ErrFailedValidation
		if errors.As(err, &validationErr) {
			b.l.Error("response body failed validation", "err", err, "update ID", upd.UpdateId)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var apiErr *gotely.ErrTelegramAPIFailedRequest
		if errors.As(err, &apiErr) {
			b.l.Error("failed request", "err", err, "update ID", upd.UpdateId)
			switch apiErr.Code {
			case 400:
				w.WriteHeader(http.StatusBadRequest)

			case 401:
				w.WriteHeader(http.StatusUnauthorized)

			case 403:
				w.WriteHeader(http.StatusForbidden)

			case 500:
				w.WriteHeader(http.StatusInternalServerError)

			default:
				w.WriteHeader(http.StatusServiceUnavailable)
			}
			return
		}

		b.l.Error("unexpected error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b.l.Info("succeeded response", "update ID", upd.UpdateId)
	w.WriteHeader(http.StatusOK)
}

type Option func(*WebhookBot)

// TODO opts

// TODO some default middlewares
