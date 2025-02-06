package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/bigelle/gotely/objects"
)

type ApiResponse[T any] struct {
	Ok          bool                        `json:"ok"`
	ErrorCode   *int                        `json:"error_code,omitempty"`
	Description *string                     `json:"description,omitempty"`
	Parameters  *objects.ResponseParameters `json:"parameters,omitempty"`
	Result      *T                          `json:"result,omitempty"`
}

type Sendable interface { // NOTE probably should rename it to make it more obviously separated from multipart objects
	objects.Validable
	Client() *http.Client
	ApiBaseUrl() string
	ToRequestBody() ([]byte, error)
}

type MultipartSendable interface {
	objects.Validable
	Client() *http.Client
	ApiBaseUrl() string
	ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error)
}

// TODO: make it better
type ErrFailedRequest struct {
	Code    *int
	Message *string
	Cause   *string
}

func (e ErrFailedRequest) Error() string {
	return fmt.Sprintf("request failed with code %d: %s", *e.Code, *e.Message)
}

func SendTelegramRequest[T any](httpMethod, token, endpoint string, body Sendable) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, fmt.Errorf("can't make request: %w", err)
	}

	data, err := body.ToRequestBody()
	if err != nil {
		return nil, fmt.Errorf("can't marshal request body: %w", err)
	}

	url := fmt.Sprintf(body.ApiBaseUrl(), token, endpoint)

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("can't form a request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := body.Client().Do(req)
	if err != nil {
		e := err.Error()
		return nil, ErrFailedRequest{
			Code:    &resp.StatusCode,
			Message: &e,
		}
	}
	defer resp.Body.Close()

	return readResult[T](resp.Body)
}

func SendTelegramGetRequest[T any](token, endpoint string, body Sendable) (*T, error) {
	return SendTelegramRequest[T]("GET", token, endpoint, body)
}

func SendTelegramPostRequest[T any](token, endpoint string, body Sendable) (*T, error) {
	return SendTelegramRequest[T]("POST", token, endpoint, body)
}

func SendTelegramMultipartRequest[T any](token, endpoint string, body MultipartSendable) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	buf, w, err := body.ToMultipartBody()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(body.ApiBaseUrl(), token, endpoint)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := body.Client().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return readResult[T](resp.Body)
}

func readResult[T any](r io.ReadCloser) (*T, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var apiResp ApiResponse[T]
	if err := json.Unmarshal(b, &apiResp); err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}
	if !apiResp.Ok {
		return nil, ErrFailedRequest{
			Code:    apiResp.ErrorCode,
			Message: apiResp.Description,
		}
	}
	if apiResp.Description != nil {
		fmt.Println(*apiResp.Description)
	}
	return apiResp.Result, nil
}
