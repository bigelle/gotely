package types

type Voice struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     *int    `json:"file_size,omitempty"`
}
