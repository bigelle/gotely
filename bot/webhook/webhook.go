package webhook

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
)

// Use this method to specify a URL and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL,
// containing a JSON-serialized [objects.Update].
// In case of an unsuccessful request (a request with response HTTP status code different from 2XY),
// we will repeat the request and give up after a reasonable amount of attempts.
// Returns True on success.
//
// If you'd like to make sure that the webhook was set by you,
// you can specify secret data in the parameter secret_token.
// If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content
type SetWebhook struct {
	// HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Url string `json:"url"`
	// Upload your public key certificate so that the root certificate in use can be checked.
	// See our self-signed guide (https://core.telegram.org/bots/self-signed) for details.
	Certificate objects.InputFile `json:"certificate,omitempty"`
	// The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	IpAddress *string `json:"ip_address,omitempty"`
	// The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40.
	// Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections *int `json:"max_connections,omitempty"`
	// A JSON-serialized list of the update types you want your bot to receive.
	// For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See [objects.Update] for a complete list of available update types.
	// Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default).
	// If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to [GetUpdates],
	// so unwanted updates may be received for a short period of time.
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
	// Pass True to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters.
	// Only characters A-Z, a-z, 0-9, _ and - are allowed.
	//  The header is useful to ensure that the request comes from a webhook set by you.
	SecretToken *string `json:"secret_token,omitempty"`

	contentType string
}

func (s SetWebhook) Endpoint() string {
	return "setWebhook"
}

func (s SetWebhook) Validate() error {
	if s.Url == "" {
		return gotely.ErrInvalidParam("url can't be empty")
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
			if err := gotely.WriteJSONToForm(mw, "allowed_updates", *s.AllowedUpdates); err != nil {
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
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates.
// Returns True on success.
type DeleteWebhook struct {
	// Pass True to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
}

func (d DeleteWebhook) Validate() error {
	return nil
}

func (d DeleteWebhook) Reader() io.Reader {
	return gotely.EncodeJSON(d)
}

func (d DeleteWebhook) Endpoint() string {
	return "deleteWebhook"
}

func (d DeleteWebhook) ContentType() string {
	return "application/json"
}

// Use this method to get current webhook status.
// Requires no parameters.
// On success, returns a [WebhookInfo] object.
// If the bot is using [longpolling.GetUpdates}, will return an object with the url field empty.
type GetWebhookInfo struct{}

func (g GetWebhookInfo) Validate() error {
	return nil
}

func (g GetWebhookInfo) Reader() io.Reader {
	return gotely.EncodeJSON(g)
}

func (g GetWebhookInfo) Endpoint() string {
	return "getWebhookInfo"
}

func (g GetWebhookInfo) ContentType() string {
	return "application/json"
}

// Describes the current status of a webhook.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`
	// Optional. Currently used webhook IP address
	IpAddress *string `json:"ip_address,omitempty"`
	// Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate *int `json:"last_error_date,omitempty"`
	// Optional. Error message in human-readable format for the most recent error
	// that happened when trying to deliver an update via webhook
	LastErrorMessage *string `json:"last_error_message,omitempty"`
	// Optional. Unix time of the most recent error
	// that happened when trying to synchronize available updates with Telegram datacenters
	LastSynchronizationErrorDate *int `json:"last_synchronization_error_date,omitempty"`
	// Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections *int `json:"max_connections,omitempty"`
	// Optional. A list of update types the bot is subscribed to.
	// Defaults to all update types except chat_member
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}
