package methods

import (
	"encoding/json"
	"net/http"

	"github.com/bigelle/gotely/objects"
)

type SetPassportDataErrors struct {
	UserId  int
	Errors  []objects.PassportElementError
	client  *http.Client
	baseUrl string
}

func (s *SetPassportDataErrors) WithClient(c *http.Client) *SetPassportDataErrors {
	s.client = c
	return s
}

func (s SetPassportDataErrors) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetPassportDataErrors) WithApiBaseUrl(u string) *SetPassportDataErrors {
	s.baseUrl = u
	return s
}

func (s SetPassportDataErrors) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s SetPassportDataErrors) Validate() error {
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(s.Errors) < 1 {
		return objects.ErrInvalidParam("objects.parameter can't be empty")
	}
	return nil
}

func (s SetPassportDataErrors) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetPassportDataErrors) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setPassportDataErrors", s)
}
