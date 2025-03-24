package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/bigelle/gotely/api/objects"
)

type SetWebhook struct {
	Url                string            `json:"url"`
	Certificate        objects.InputFile `json:"certificate,omitempty"`
	IpAddress          *string           `json:"ip_address,omitempty"`
	MaxConnections     *int              `json:"max_connections,omitempty"`
	AllowedUpdates     *[]string         `json:"allowed_updates,omitempty"`
	DropPendingUpdates *bool             `json:"drop_pending_updates,omitempty"`
	SecretToken        *string           `json:"secret_token,omitempty"`
	contentType        string
}

func (s SetWebhook) Endpoint() string {
	return "setWebhook"
}

func (s SetWebhook) Validate() error {
	if s.Url == "" {
		return fmt.Errorf("url can't be empty")
	}
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SetWebhook) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("url", s.Url); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Certificate != nil {
			if err := s.Certificate.WriteTo(mw, "certificate"); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.IpAddress != nil {
			if err := mw.WriteField("ip_address", *s.IpAddress); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MaxConnections != nil {
			if err := mw.WriteField("max_connections", fmt.Sprint(s.MaxConnections)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowedUpdates != nil {
			b, err := json.Marshal(s.AllowedUpdates)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			if err := mw.WriteField("allowed_updates", string(b)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DropPendingUpdates != nil {
			if err := mw.WriteField("drop_pending_updates", fmt.Sprint(s.DropPendingUpdates)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.SecretToken != nil {
			if err := mw.WriteField("secret_token", *s.SecretToken); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()

	return pr
}

func (s SetWebhook) ContentType() string {
	return s.contentType
}

func (s SetWebhook) HttpMethod() string {
	return http.MethodPost
}
