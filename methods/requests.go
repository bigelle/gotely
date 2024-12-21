package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bigelle/tele.go/types"
)

type Executable interface {
	types.Validable
	ToRequestBody() ([]byte, error)
}

const api_url = "https://api.telegram.org/bot%s/%s"

type ErrFailedRequest struct {
	Code    int
	Message string
}

func (e ErrFailedRequest) Error() string {
	return fmt.Sprintf("request failed with code %d: %s", e.Code, e.Message)
}

func MakeRequest[T any](httpMethod, token, endpoint string, body Executable) (*T, error) {
	if token == "" {
		return nil, types.ErrInvalidParam("token can't be empty. did you connect the bot?")
	}

	if err := body.Validate(); err != nil {
		return nil, fmt.Errorf("can't make request: %w", err)
	}

	data, err := body.ToRequestBody()
	if err != nil {
		return nil, fmt.Errorf("can't marshal request body: %w", err)
	}

	url := fmt.Sprintf(api_url, token, endpoint)

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("can't form a request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func MakeGetRequest[T any](token, endnpoint string, body Executable) (*T, error) {
	return MakeRequest[T]("GET", token, endnpoint, body)
}

func MakePostRequest[T any](token, endpoint string, body Executable) (*T, error) {
	return MakeRequest[T]("POST", token, endpoint, body)
}

type ApiResponse[T any] struct {
	Ok          bool                      `json:"ok"`
	ErrorCode   int                       `json:"error_code"`
	Description *string                   `json:"description,omitempty"`
	Parameters  *types.ResponseParameters `json:"parameters,omitempty"`
	Result      T                         `json:"result"`
}
