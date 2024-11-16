package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const api_url = "https://api.telegram.org/bot%s/%s"

func MakeRequest(httpMethod, token, endpoint string, body []byte) ([]byte, error) {
	url := fmt.Sprintf(api_url, token, endpoint)

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(body))
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
	return b, nil
}

func MakeGetRequest(token, endnpoint string, body []byte) ([]byte, error) {
	return MakeRequest("GET", token, endnpoint, body)
}

func MakePostRequest(token, endpoint string, body []byte) ([]byte, error) {
	return MakeRequest("POST", token, endpoint, body)
}
