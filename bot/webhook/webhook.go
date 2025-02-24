package webhook

import (
	"io"
	"mime/multipart"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
)

type setWebhook struct {
	Url                string             `json:"url"`
	Certificate        *objects.InputFile `json:"certificate,omitempty"`
	IpAddress          *string            `json:"ip_address,omitempty"`
	MaxConnections     *int               `json:"max_connections,omitempty"`
	AllowedUpdates     *[]string          `json:"allowed_updates,omitempty"`
	DropPendingUpdates *bool              `json:"drop_pending_updates,omitempty"`
	SecretToken        *string            `json:"secret_token,omitempty"`
	contentType        string
}

func (s setWebhook) Endpoint() string {
	return "setWebhook"
}

func (s setWebhook) Validate() error {
	//TODO
	return nil
}

func (s *setWebhook) Reader() (io.Reader, error) {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		if err := mw.WriteField("url", s.Url); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Certificate != nil {
			if c, ok := (*s.Certificate).(objects.InputFileFromRemote); ok {
				part, err := mw.CreateFormFile("certificate", c.Name())
				if err != nil {
					pw.CloseWithError(err)
					return
				}
				r, err := c.Reader()
				if err != nil {
					pw.CloseWithError(err)
					return
				}
				if _, err := io.Copy(part, r); err != nil {
					pw.CloseWithError(err)
					return
				}
			}
		}
		if s.IpAddress != nil {
			if err := mw.WriteField("ip_address", *s.IpAddress); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		//TODO other fields
	}()

	mw.Close()
	return pr, nil
}

func (s setWebhook) ContentType() string {
	return s.contentType
}

func SetWebhook(token, toUrl string) error {
	sw := setWebhook{
		Url: toUrl,
	}

	_, err := gotely.SendPostRequestWith[bool](&sw, token)
	return err
}
