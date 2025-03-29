package webhook

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/tgbot"
)

// WebhookBot is used to create a simple webhook server
// that will respond to updates coming from Telegram Bot API.
type WebhookBot struct {
	Bot tgbot.Bot

	s          *http.Server
	path       string
	addr       string
	middleware []func(next http.Handler) http.Handler
	l          *slog.Logger
}

// New is used to create a new instance of [WebhookBot] using specified options
func New(bot tgbot.Bot, opts ...Option) WebhookBot {
	b := WebhookBot{
		Bot: bot,

		addr:       ":8080",
		path:       "/webhook",
		middleware: []func(next http.Handler) http.Handler{RecoveryMiddleware, LoggingMiddleware},
		l:          slog.Default(),
	}

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

	for _, opt := range opts {
		opt(&b)
	}
	return b
}

// Use is used to add a middleware that will be wrapped around bot's update handler.
func (b *WebhookBot) Use(m ...func(http.Handler) http.Handler) {
	b.middleware = append(b.middleware, m...)
}

// Start is launching bot's [http.Server]
func (b *WebhookBot) Start() error {
	b.l.Info("webhook server is listening and serving on", "addr", b.addr, "path", b.path)
	return b.s.ListenAndServe()
}

// Stop is shutting down bot's [http.Server] giving it 5 seconds to complete current requests
func (b WebhookBot) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return b.s.Shutdown(ctx)
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
	b.l.Debug("succeeded response", "update ID", upd.UpdateId)
	w.WriteHeader(http.StatusOK)
}

type Option func(*WebhookBot)

// TODO opts

// WithServer is used to completely setup bot's [http.Server]
func WithServer(s *http.Server) Option {
	return func(wb *WebhookBot) {
		wb.s = s
	}
}

// WithAddress is used to set address that [http.Server] will be working on.
// Defaults to ":8080"
func WithAddress(addr string) Option {
	return func(wb *WebhookBot) {
		wb.addr = addr
	}
}

// WithPath is used to set path that will be used when registering bot's [http.Handler].
// Defaults to "/webhook"
func WithPath(p string) Option {
	return func(wb *WebhookBot) {
		wb.path = p
	}
}

// WithLogger is used to replace default [slog.Logger] that is used
// to report about any occurred errors, warnings, debug info, etc.
func WithLogger(l *slog.Logger) Option {
	return func(wb *WebhookBot) {
		wb.l = l
	}
}

// RecoveryMiddleware is used to simply recover from panic
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic", r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware is used to report about every request and how long it took
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info(r.Pattern, "took", time.Since(start))
	})
}
