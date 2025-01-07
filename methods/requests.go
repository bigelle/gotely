package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/objects"
)

type ApiResponse[T any] struct {
	Ok          bool                        `json:"ok"`
	ErrorCode   int                         `json:"error_code"`
	Description *string                     `json:"description,omitempty"`
	Parameters  *objects.ResponseParameters `json:"parameters,omitempty"`
	Result      T                           `json:"result"`
}

type Sendable interface { // NOTE probably should rename it to make it more obviously separated from multipart objects
	objects.Validable
	ToRequestBody() ([]byte, error)
}

type MultipartSendable interface {
	objects.Validable
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

func MakeRequest[T any](httpMethod, endpoint string, body Sendable) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, fmt.Errorf("can't make request: %w", err)
	}

	settings := telego.GetBotSettings()
	token := settings.Token
	if token == "" {
		return nil, fmt.Errorf("API token can't be empty")
	}

	data, err := body.ToRequestBody()
	if err != nil {
		return nil, fmt.Errorf("can't marshal request body: %w", err)
	}

	url := fmt.Sprintf(settings.BaseUrl, token, endpoint)

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("can't form a request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := settings.Client.Do(req)
	if err != nil {
		e := err.Error()
		return nil, ErrFailedRequest{
			Code:    &resp.StatusCode,
			Message: &e,
		}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	var apiResp ApiResponse[T]
	if err := json.Unmarshal(b, &apiResp); err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}
	if !apiResp.Ok {
		return nil, ErrFailedRequest{
			Code:    &apiResp.ErrorCode,
			Message: apiResp.Description,
		}
	}
	return &apiResp.Result, nil
}

func MakeGetRequest[T any](endpoint string, body Sendable) (*T, error) {
	return MakeRequest[T]("GET", endpoint, body)
}

func MakePostRequest[T any](endpoint string, body Sendable) (*T, error) {
	return MakeRequest[T]("POST", endpoint, body)
}

func MakeMultipartRequest[T any](endpoint string, body MultipartSendable) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	settings := telego.GetBotSettings()

	buf, w, err := body.ToMultipartBody()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(settings.BaseUrl, settings.Token, endpoint)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := settings.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp ApiResponse[T]
	if err := json.Unmarshal(b, &apiResp); err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}
	if !apiResp.Ok {
		return nil, ErrFailedRequest{
			Code:    &apiResp.ErrorCode,
			Message: apiResp.Description,
		}
	}
	return &apiResp.Result, nil
}
