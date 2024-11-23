package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bigelle/tele.go/types"
)

const api_url = "https://api.telegram.org/bot%s/%s"

func MakeRequest[T any](httpMethod, token, endpoint string, body Executable) (*T, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(api_url, token, endpoint)

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp types.ApiResponse[T]
	if err := json.Unmarshal(b, &apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Ok {
		return nil, fmt.Errorf("%d: %s", apiResp.ErrorCode, *apiResp.Description)
	}
	return &apiResp.Result, nil
}

func MakeGetRequest[T any](token, endnpoint string, body Executable) (*T, error) {
	return MakeRequest[T]("GET", token, endnpoint, body)
}

func MakePostRequest[T any](token, endpoint string, body Executable) (*T, error) {
	return MakeRequest[T]("POST", token, endpoint, body)
}
