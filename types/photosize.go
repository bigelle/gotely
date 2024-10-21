package types

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FilePath     string `json:"file_path"`
	FileSize     *int   `json:"file_size,omitempty"`
}
