package gotely

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ApiResponse represents a response from Telegram Bot API
type ApiResponse struct {
	// true if request was successful, otherwise false
	Ok bool `json:"ok"`
	// error code of unsuccessful request
	ErrorCode *int `json:"error_code,omitempty"`
	// a human-readable description of the result
	Description *string `json:"description,omitempty"`
	// Describes why a request was unsuccessful.
	Parameters *ResponseParameters `json:"parameters,omitempty"`
	// result of request
	Result json.RawMessage `json:"result,omitempty"`
}

// ResponseParameters describes why a request was unsuccessful.
type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and
	// some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId *int `json:"migrate_to_chat_id,omitempty"`
	// Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter *int `json:"retry_after,omitempty"`
}

// Method represents a Telegram Bot API method as a structured set of request parameters.
type Method interface {
	// Endpoint returns the API method name as a string.
	// For example, `SendMessage` should return "sendMessage".
	Endpoint() string

	// Validate checks if the request contains valid data.
	// Returns an error if the request is invalid, otherwise nil.
	Validate() error

	// Reader returns an `io.Reader` representing the request body.
	Reader() io.Reader

	// ContentType returns the appropriate content type:
	// - "application/json" for standard requests
	// - "multipart/form-data" (with a pre-generated boundary) for file uploads
	ContentType() string
}

// RequestOption represents a function that modifies `RequestConfig`.
// It is used to customize request settings when calling `SendRequestWith`.
type RequestOption func(*RequestConfig)

// RequestConfig defines configuration options for sending a request to the Telegram Bot API.
type RequestConfig struct {
	// Client is the HTTP client used to send the request.
	// Defaults to `http.DefaultClient`.
	Client *http.Client
	// ApiUrl is the Telegram Bot API base URL.
	// Defaults to "https://api.telegram.org/bot%s/%s",
	// where the first placeholder is replaced by the bot token and the second by the API method.
	// Use `%s` placeholders to properly insert the token and API method.
	ApiUrl string

	Context context.Context
}

// WithClient sets a custom HTTP client for `SendRequestWith`.
// Use this option to provide a specific `http.Client` for making requests.
func WithClient(c *http.Client) RequestOption {
	return func(rc *RequestConfig) {
		rc.Client = c
	}
}

// WithUrl sets a custom Telegram Bot API URL for `SendRequestWith`.
// Use this option if running a local instance of the Telegram Bot API.
func WithUrl(url string) RequestOption {
	return func(rc *RequestConfig) {
		rc.ApiUrl = url
	}
}

func WithContext(ctx context.Context) RequestOption {
	return func(rc *RequestConfig) {
		rc.Context = ctx
	}
}

var defaultReqCfg = RequestConfig{
	Client:  http.DefaultClient,
	ApiUrl:  DEFAULT_URL_TEMPLATE,
	Context: context.Background(),
}

func makeReqCfg(opts ...RequestOption) RequestConfig {
	cfg := defaultReqCfg
	for _, opt := range opts {
		opt(&cfg)
	}

	// fallback
	if cfg.Client == nil {
		cfg.Client = http.DefaultClient
	}
	if !IsCorrectUrlTemplate(cfg.ApiUrl) {
		cfg.ApiUrl = defaultReqCfg.ApiUrl
	}
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}
	return cfg
}

// SendRequestWith sends a request to the Telegram Bot API using the provided token,
// and with parameters described in body.
// If dest is not nil, the response content is written to it.
// Pass nil as dest to ignore the response content.
func SendRequest(body Method, token string, dest any) error {
	return SendRequestWith(body, token, dest)
}

// SendRequestWith sends a request to the Telegram Bot API using the provided token,
// with parameters described in body and optional request options opts.
// If dest is not nil, the response content is written to it.
// Pass nil as dest to ignore the response content.
func SendRequestWith(body Method, token string, dest any, opts ...RequestOption) error {
	if err := body.Validate(); err != nil {
		return err
	}
	if token == "" {
		return fmt.Errorf("API token can't be empty")
	}

	cfg := makeReqCfg(opts...)

	url := formatUrl(cfg.ApiUrl, token, body.Endpoint())
	// its important to call Reader() before using ContentType()
	// since content-type boundary is generated inside Reader() and stored inside of a struct
	req, err := http.NewRequestWithContext(cfg.Context, http.MethodPost, url, body.Reader())
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", body.ContentType())

	resp, err := cfg.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result ApiResponse
	if err := DecodeJSON(resp.Body, &result); err != nil {
		return err
	}

	if !result.Ok {
		return ErrTelegramAPIFailedRequest{
			Code:               *result.ErrorCode,
			Description:        *result.Description,
			ResponseParameters: &ResponseParameters{},
		}
	}
	// not writing any results if destination is nil
	// not returning any errors because the request itself was successful
	if dest == nil {
		return nil
	}
	return json.NewDecoder(bytes.NewReader(result.Result)).Decode(dest)
}

func formatUrl(template, token, method string) string {
	url := strings.Replace(template, "<token>", token, 1)
	url = strings.Replace(url, "<method>", method, 1)
	return url
}
