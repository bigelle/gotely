package webhook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
)

func Connect(url, token string, opts ...Option) error {
	swh := setWebhook{
		Url:    url,
		client: http.DefaultClient,
		apiUrl: "https://api.telegram.org/bot%s/%s",
	}
	for _, opt := range opts {
		opt(&swh)
	}
	_, err := gotely.SendPostRequestWith[bool](
		&swh,
		token,
		gotely.WithClient(swh.client),
		gotely.WithUrl(swh.apiUrl),
	)
	return err
}

type Option func(*setWebhook)

// TODO opts

type setWebhook struct {
	Url                string            `json:"url"`
	Certificate        objects.InputFile `json:"certificate,omitzero"`
	IpAddress          string            `json:"ip_address,omitzero"`
	MaxConnections     int               `json:"max_connections,omitzero"`
	AllowedUpdates     *[]string         `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool              `json:"drop_pending_updates,omitzero"`
	SecretToken        string            `json:"secret_token,omitzero"`
	contentType        string
	client             *http.Client
	apiUrl             string
}

func (s setWebhook) Endpoint() string {
	return "setWebhook"
}

func (s setWebhook) Validate() error {
	//FIXME
	return nil
}

func (s *setWebhook) Reader() (io.Reader, error) {
	//FIXME multipart
	return nil, nil
}

func (s setWebhook) ContentType() string {
	return s.contentType
}

func Disconnect(token string, dropPendingUpdates ...bool) error {
	_, err := gotely.SendPostRequestWith[bool](
		deleteWebhook{
			DropPendingUpdates: dropPendingUpdates[0],
		},
		token,
	)
	return err
}

type deleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitzero"`
}

func (d deleteWebhook) Endpoint() string {
	return "deleteWebhook"
}

func (d deleteWebhook) Validate() error {
	return nil
}

func (d deleteWebhook) Reader() (io.Reader, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (d deleteWebhook) ContentType() string {
	return "application/json"
}

// FIXME maybe should make a WebhookBot struct and make all of this methods part of the struct
func GetWebhookInfo(token string) (*WebHookInfo, error) {
	return gotely.SendGetRequestWith[WebHookInfo](getWebhookInfo{}, token)
}

type getWebhookInfo struct {
}

func (g getWebhookInfo) Endpoint() string {
	return "getWebhookInfo"
}

func (g getWebhookInfo) Validate() error {
	return nil
}

func (g getWebhookInfo) Reader() (io.Reader, error) {
	return io.NopCloser(strings.NewReader("")), nil
}

func (g getWebhookInfo) ContentType() string {
	return ""
}

type WebHookInfo struct {
	Url                          string    `json:"url"`
	HasCustomCertificate         bool      `json:"has_custom_certificate"`
	PendingUpdateCount           int       `json:"pending_update_count"`
	IpAddress                    *string   `json:"ip_address,omitempty"`
	LastErrorDate                *int      `json:"last_error_date,omitempty"`
	LastErrorMessage             *string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate *int      `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               *int      `json:"max_connections,omitempty"`
	AllowedUpdates               *[]string `json:"allowed_updates,omitempty"`
}
