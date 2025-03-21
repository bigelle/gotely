package gotely

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Reader() (io.Reader, error)

	// ContentType returns the appropriate content type:
	// - "application/json" for standard requests
	// - "multipart/form-data" (with a pre-generated boundary) for file uploads
	ContentType() string
}

// SendRequest sends a request using the given HTTP `method` and `token` with `body` as parameters.
// It is a wrapper around `SendRequestWith`, without additional request options (`RequestOption`).
// Returns a pointer to the expected response object `T` or an error if the request fails.
func SendRequest(body Method, dest any, token string) error {
	return SendRequestWith(body, dest, token)
}

// SendRequestWith sends a request using the given API `method`, `token`, request parameters `body`,
// and additional request options (`opts`).
// Returns a pointer to the expected response object `T` or an error if the request fails.
func SendRequestWith(body Method, dest any, token string, opts ...RequestOption) error {
	if err := body.Validate(); err != nil {
		return err
	}

	// its important to call Reader() before using ContentType()
	// since content-type boundary is generated inside Reader() and stored inside of a struct
	r, err := body.Reader()
	if err != nil {
		return err
	}

	cfg := RequestConfig{
		Client: http.DefaultClient,
		// TODO: replace %s placeholders with <token> and <method>
		ApiUrl: "https://api.telegram.org/bot%s/%s",
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	url := fmt.Sprintf(cfg.ApiUrl, token, body.Endpoint())
	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", body.ContentType())

	resp, err := cfg.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result ApiResponse
	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("bad request: %s", *result.Description)
	}
	// TODO: maybe use decoder
	return json.NewDecoder(bytes.NewReader(result.Result)).Decode(dest)
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
