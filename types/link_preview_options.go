package types

import "fmt"

type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`
	UrlFileId        *string `json:"url_file_id,omitempty"`
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`
}

func (l LinkPreviewOptions) Validate() error {
	if *l.PreferLargeMedia && *l.PreferSmallMedia {
		return fmt.Errorf("PreferSmallMedia and PreferLargeMedia parameters are mutual exclusive")
	}
	return nil
}
