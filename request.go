package gotely

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ApiResponse represents a response from a request to Telegram Bot API
type ApiResponse[T any] struct {
	//true if request was successful, otherwise false
	Ok bool `json:"ok"`
	//error code of unsuccessful request
	ErrorCode *int `json:"error_code,omitempty"`
	//a human-readable description of the result
	Description *string `json:"description,omitempty"`
	//Describes why a request was unsuccessful.
	Parameters *ResponseParameters `json:"parameters,omitempty"`
	//result of request
	Result *T `json:"result,omitempty"`
}

// ResponseParameters describes why a request was unsuccessful.
type ResponseParameters struct {
	//Optional. The group has been migrated to a supergroup with the specified identifier.
	//This number may have more than 32 significant bits and
	//some programming languages may have difficulty/silent defects in interpreting it.
	//But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId *int `json:"migrate_to_chat_id,omitempty"`
	//Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter *int `json:"retry_after,omitempty"`
}

// Method represents a Telegram Bot API method as a structured set of request parameters.
type Method interface {
	//Should return an API endpoint of this method as a string.
	//For example, in case with `SendMessage` it will return "sendMessage".
	Endpoint() string
	//Should return an error if request contains invalid data or a nil
	Validate() error
	//Should return `io.Reader` of this method body
	Reader() (io.Reader, error)
	//Returns "application/json" for regular methods
	//and "multipart/form-data" with pre-generated boundary
	ContentType() string
}

// SendRequest is a wrapper around `SendRequestWith` with no `RequestOption` specified,
// and is used to send a set of request parameters `body` using API Token `token` and HTTP Method `method`.
// Returns a pointer to specified expected object T or an error if something is went wrong
func SendRequest[T any](body Method, token string, method string) (*T, error) {
	return SendRequestWith[T](body, token, method)
}

// SendPostRequest is a wrapper around SendRequest that is using `http.MethodPost` as HTTP Method
// and is used to send a set of request parameters `body` using API Token `token`
// Returns a pointer to specified expected object T or an error if something is went wrong
func SendPostRequest[T any](body Method, token string) (*T, error) {
	return SendRequest[T](body, token, http.MethodPost)
}

// SendGetRequest is a wrapper around SendRequest that is using `http.MethodGet` as HTTP Method
// and is used to send a set of request parameters `body` using API Token `token`
// Returns a pointer to specified expected object T or an error if something is went wrong
func SendGetRequest[T any](body Method, token string) (*T, error) {
	return SendRequest[T](body, token, http.MethodGet)
}

// SendRequestWith is used to send a set of request parameters `body` using API Token `token`, HTTP Method `method`
// and a list of request options `opts`.
// Returns a pointer to specified expected object T or an error if something is went wrong
func SendRequestWith[T any](body Method, token string, method string, opts ...RequestOption) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	//its important to call Reader() before using ContentType()
	//since content-type boundary is generated inside Reader() and stored inside of a struct
	r, err := body.Reader()
	if err != nil {
		return nil, err
	}

	cfg := RequestConfig{
		Client:         http.DefaultClient,
		RequestBaseUrl: "https://api.telegram.org/bot%s/%s",
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	url := fmt.Sprintf(cfg.RequestBaseUrl, token, body.Endpoint())
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", body.ContentType())

	resp, err := cfg.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ApiResponse[T]
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, fmt.Errorf("bad request: %s", *result.Description)
	}
	return result.Result, nil
}

// SendPostRequest is a wrapper around SendRequestWith that is using `http.MethodPost` as HTTP Method
// and is used to send a set of request parameters `body` using API Token `token` and a list of request options `opts`
// Returns a pointer to specified expected object T or an error if something is went wrong
func SendPostRequestWith[T any](body Method, token string, opts ...RequestOption) (*T, error) {
	return SendRequestWith[T](body, token, http.MethodPost, opts...)
}

// SendGetRequest is a wrapper around SendRequestWith that is using `http.MethodPost` as HTTP Method
// and is used to send a set of request parameters `body` using API Token `token` and a list of request options `opts`
// Returns a pointer to specified expected object T or an error if something is went wrong
func SendGetRequestWith[T any](body Method, token string, opts ...RequestOption) (*T, error) {
	return SendRequestWith[T](body, token, http.MethodGet, opts...)
}

// RequestOption is a type based on a function that is accepting a pointer to a `RequestConfig`
// and is used when sending a request using `SendRequestWith` with specified options
type RequestOption func(*RequestConfig)

// RequestConfig is a structured configuration that will be used while sending a request to Telegram Bot API
type RequestConfig struct {
	//custom client. Defaults to `http.DefaultClient`
	Client *http.Client
	//Custom Telegram Bot API URL. Defaults to "https://api.telegram.org/bot%s/%s",
	//where the first placeholder is used for API token and the second one is for API endpoint.
	//Please use %s placeholders to properly use Bot API token and API endpoint
	RequestBaseUrl string
}

// WithClient is a functional option that is used when sending a request using `SendRequestWith`.
// Use it to provide a custom `http.Client` that will be used when doing the request
func WithClient(c *http.Client) RequestOption {
	return func(rc *RequestConfig) {
		rc.Client = c
	}
}

// WithUrl is a functional option that is used when sending a request using `SendRequestWith`.
// Use it to provide your own Telegram Bot API URL in case you're running it locally
func WithUrl(url string) RequestOption {
	return func(rc *RequestConfig) {
		rc.RequestBaseUrl = url
	}
}

// WithConfig is a functional option that is used when sending a request using `SendRequestWith`.
// Use it to configure the request all at once
func WithConfig(cfg RequestConfig) RequestOption {
	return func(rc *RequestConfig) {
		rc = &cfg
	}
}
