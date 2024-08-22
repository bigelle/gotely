package objects

import (
	"fmt"
	"strings"
)

type File struct {
	fileId       string
	fileUniqueId string
	fileSize     int64
	filePath     string
}

func (f File) GetFileUrl(botToken string) (string, error) {
	if botToken == "" || strings.TrimSpace(botToken) == "" {
		return "", fmt.Errorf("bot token can't be empty")
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, f.filePath), nil
}
