package types 

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Performer    *string    `json:"performer,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
}
