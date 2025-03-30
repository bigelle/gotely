package webhook

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/objects"
	"github.com/bigelle/gotely/tgbot"
)

// WebhookBot creates a simple webhook server
// that responds to updates from the Telegram Bot API.
type WebhookBot struct {
	Bot tgbot.Bot

	s               *http.Server
	path            string
	addr            string
	middleware      []func(next http.Handler) http.Handler
	shutdownTimeout time.Duration
	l               *slog.Logger

	certFile string
	keyFile  string
	useTLS   bool
}

// New creates a new instance of [WebhookBot] using the specified options.
func New(bot tgbot.Bot, opts ...Option) WebhookBot {
	b := WebhookBot{
		Bot: bot,

		addr:            ":8080",
		path:            "/webhook",
		middleware:      []func(next http.Handler) http.Handler{RecoveryMiddleware, LoggingMiddleware},
		shutdownTimeout: 5 * time.Second,
		l:               slog.Default(),
	}
	for _, opt := range opts {
		opt(&b)
	}

	if b.s == nil {
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
	}
	return b
}

// Use adds middleware that wraps the bot's update handler.
func (b *WebhookBot) Use(m ...func(http.Handler) http.Handler) {
	b.middleware = append(b.middleware, m...)
}

// SetMiddleware is used to completely change the list of bot's middleware
func (b *WebhookBot) SetMiddleware(m []func(http.Handler) http.Handler) {
	b.middleware = m
}

// Start launches the bot's [http.Server].
func (b *WebhookBot) Start() error {
	b.l.Info("webhook server is listening and serving on", "addr", b.addr, "path", b.path)
	if b.useTLS {
		b.l.Debug("starting HTTPS server with TLS", "cert", b.certFile, "key", b.keyFile)
		return b.s.ListenAndServeTLS(b.certFile, b.keyFile)
	}
	return b.s.ListenAndServe()
}

// Stop shuts down the bot's [http.Server], allowing the time specified in the bot's settings for active requests to complete.
func (b WebhookBot) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), b.shutdownTimeout)
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

// WithReadsTimeout sets the timeout for the bot's [http.Server].
func WithReadTimeout(t time.Duration) Option {
	return func(wb *WebhookBot) {
		wb.s.ReadTimeout = t
	}
}

// WithWriteTimeout sets the write timeout for the bot's [http.Server].
func WithWriteTimeout(t time.Duration) Option {
	return func(wb *WebhookBot) {
		wb.s.WriteTimeout = t
	}
}

// WithIdleTimeout sets the idle timeout for the bot's [http.Server].
func WithIdleTimeout(t time.Duration) Option {
	return func(wb *WebhookBot) {
		wb.s.IdleTimeout = t
	}
}

// WithTLS sets the certFile and keyFile to enable HTTPS.
func WithTLS(certFile, keyFile string) Option {
	return func(wb *WebhookBot) {
		wb.useTLS = true
		wb.certFile = certFile
		wb.keyFile = keyFile
	}
}

// WithAddress sets the address for the [http.Server].
// Defaults to ":8080".
func WithAddress(addr string) Option {
	return func(wb *WebhookBot) {
		wb.addr = addr
	}
}

// WithPath sets the path for registering the bot's [http.Handler].
// Defaults to "/webhook".
func WithPath(p string) Option {
	return func(wb *WebhookBot) {
		wb.path = p
	}
}

// WithLogger replaces the default [slog.Logger] used
// for reporting errors, warnings, debug information, etc.
func WithLogger(l *slog.Logger) Option {
	return func(wb *WebhookBot) {
		wb.l = l
	}
}

// WithLogLevel sets the logging level for the bot's logger.
func WithLogLevel(l slog.Level) Option {
	return func(wb *WebhookBot) {
		wb.l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: l}))
	}
}

// WithMiddleware replaces the bot's middleware with m.
func WithMiddleware(m []func(http.Handler) http.Handler) Option {
	return func(wb *WebhookBot) {
		wb.SetMiddleware(m)
	}
}

// WithCustomHandler replaces the bot's default [http.Handler].
func WithCustomHandler(h http.Handler) Option {
	return func(wb *WebhookBot) {
		wb.s.Handler = h
	}
}

// WithShutdownTimeout sets the time the bot will wait before
// aborting unanswered requests when calling Stop().
func WithShutdownTimeout(t time.Duration) Option {
	return func(wb *WebhookBot) {
		wb.shutdownTimeout = t
	}
}

// RecoveryMiddleware recovers from panics in the request handling pipeline.
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

// LoggingMiddleware logs each request and its duration.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info(r.Pattern, "took", time.Since(start))
	})
}
