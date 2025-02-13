package webhook

type WebHookInfo struct {
	Url                          string    `json:"url"`
	HasCustomCertificate         bool      `json:"has_custom_certificate"`
	PendingUpdatesCount          int       `json:"pending_updates_count"`
	IpAddress                    *string   `json:"ip_address,omitempty"`
	LastErrorDate                *int      `json:"last_error_date,omitempty"`
	LastErrorMessage             *string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate *int      `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               *int      `json:"max_connections,omitempty"`
	AllowedUpdates               *[]string `json:"allowed_updates,omitempty"`
}
