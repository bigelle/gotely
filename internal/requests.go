package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/types"
)

const api_url = "https://api.telegram.org/bot%s/%s"

func MakeRequest[T any](httpMethod, token, endpoint string, body Executable) (*T, error) {
	if token == "" {
		return nil, errors.ErrInvalidParam("token can't be empty. did you connect the bot?")
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
		return nil, errors.ErrFailedRequest{
			Code:    resp.StatusCode,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	var apiResp types.ApiResponse[T]
	if err := json.Unmarshal(b, &apiResp); err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}
	if !apiResp.Ok {
		return nil, errors.ErrFailedRequest{
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
