package objects

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type InputFile struct {
	AttachName     string
	MediaName      string
	newMediaFile   *os.File
	newMediaStream io.Reader // FIXME: possibly not work
	isNew          bool
}

func (f *InputFile) SetMediaName(fileName string) *InputFile {
	f.MediaName = fileName
	return f
}

func (f *InputFile) SetMediaFile(file *os.File) *InputFile {
	f.newMediaFile = file
	if f.MediaName == "" {
		f.AttachName = "attach://" + file.Name()
	} else {
		f.AttachName = "attach://" + f.MediaName
	}
	f.isNew = true
	return f
}

func (f *InputFile) SetMediaStream(stream io.Reader, fileName string) *InputFile {
	f.newMediaStream = stream
	f.MediaName = fileName
	f.AttachName = "attach://" + fileName
	f.isNew = true
	return f
}

func (f *InputFile) SetAttachName(attachName string) *InputFile {
	f.AttachName = attachName
	f.isNew = false
	return f
}

func (f InputFile) Validate() error {
	if f.isNew {
		if f.MediaName == "" || strings.TrimSpace(f.MediaName) == "" {
			return fmt.Errorf("media name can't be empty")
		}
		if f.newMediaFile == nil && f.newMediaStream == nil {
			return fmt.Errorf("media can't be empty")
		}
	} else {
		if f.AttachName == "" || strings.TrimSpace(f.AttachName) == "" {
			return fmt.Errorf("file_id can't be empty")
		}
	}
	return nil
}
