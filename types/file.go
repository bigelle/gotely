package types

import (
	"fmt"
	"strings"
)

type File struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FileSize     *int64  `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

func (f File) GetFileUrl(botToken string) (string, error) {
	if botToken == "" || strings.TrimSpace(botToken) == "" {
		return "", fmt.Errorf("bot token can't be empty")
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, *f.FilePath), nil
}
