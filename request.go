package gotely

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ApiResponse represents a response from a request to Telegram Bot API
type ApiResponse[T any] struct {
	Ok          bool                `json:"ok"`
	ErrorCode   *int                `json:"error_code,omitempty"`
	Description *string             `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      *T                  `json:"result,omitempty"`
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
	//Returns an API endpoint of this method as a string.
	//For example, `SendMessage` will return "sendMessage".
	Endpoint() string
	//Returns an error if request contains invalid data or a nil
	Validate() error
}

func SendRequest[T any](body Method, token string, method string) (*T, error) {
	return SendRequestWith[T](body, token, method)
}

func SendPostRequest[T any](body Method, token string) (*T, error) {
	return SendRequest[T](body, token, http.MethodPost)
}

func SendGetRequest[T any](body Method, token string) (*T, error) {
	return SendRequest[T](body, token, http.MethodGet)
}

func SendRequestWith[T any](body Method, token string, method string, opts ...RequestOption) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)

	cfg := RequestConfig{
		Client:         *&http.DefaultClient,
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

	resp, err := cfg.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
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

func SendPostRequestWith[T any](body Method, token string, opts ...RequestOption) (*T, error) {
	return SendRequestWith[T](body, token, http.MethodPost, opts...)
}

func SendGetRequestWith[T any](body Method, token string, opts ...RequestOption) (*T, error) {
	return SendRequestWith[T](body, token, http.MethodGet, opts...)
}

type RequestOption func(*RequestConfig)

type RequestConfig struct {
	Client         *http.Client
	RequestBaseUrl string
}
