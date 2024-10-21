package types

import (
	"fmt"
	"regexp"
)

type InputFile string

func (i InputFile) Validate() error {
	urlRegex := regexp.MustCompile(`^https?://`)
	attachmentRegex := regexp.MustCompile(`^attach://[\w-]+$`)
	switch {
	case urlRegex.MatchString(string(i)):
		return nil
	case attachmentRegex.MatchString(string(i)):
		return nil
	default:
		return fmt.Errorf(
			"Invalid media parameter. Please refer to: https://core.telegram.org/bots/api#sending-files",
		)
	}
}
