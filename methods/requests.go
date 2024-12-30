package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/objects"
)

type Executable interface {
	objects.Validable
	ToRequestBody() ([]byte, error)
}

type ErrFailedRequest struct {
	Code    int
	Message string
}

func (e ErrFailedRequest) Error() string {
	return fmt.Sprintf("request failed with code %d: %s", e.Code, e.Message)
}

func MakeRequest[T any](httpMethod, endpoint string, body Executable) (*T, error) {
	settings := telego.GetBotSettings()
	token := settings.Token
	if token == "" {
		return nil, fmt.Errorf("API token can't be empty")
	}

	if err := body.Validate(); err != nil {
		return nil, fmt.Errorf("can't make request: %w", err)
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
		return nil, ErrFailedRequest{
			Code:    resp.StatusCode,
			Message: err.Error(),
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
			Code:    apiResp.ErrorCode,
			Message: *apiResp.Description,
		}
	}
	return &apiResp.Result, nil
}

func MakeGetRequest[T any](endnpoint string, body Executable) (*T, error) {
	return MakeRequest[T]("GET", endnpoint, body)
}

func MakePostRequest[T any](endpoint string, body Executable) (*T, error) {
	return MakeRequest[T]("POST", endpoint, body)
}

type ApiResponse[T any] struct {
	Ok          bool                        `json:"ok"`
	ErrorCode   int                         `json:"error_code"`
	Description *string                     `json:"description,omitempty"`
	Parameters  *objects.ResponseParameters `json:"parameters,omitempty"`
	Result      T                           `json:"result"`
}
