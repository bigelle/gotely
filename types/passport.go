package types

import (
	"encoding/json"
	"errors"
	"strings"

	errors1 "github.com/bigelle/tele.go/errors"
)

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type EncryptedPassportElement struct {
	Type         string          `json:"type"`
	Hash         string          `json:"hash"`
	Data         *string         `json:"data,omitempty"`
	PhoneNumber  *string         `json:"phone_number,omitempty"`
	Email        *string         `json:"email,omitempty"`
	Files        *[]PassportFile `json:"files,omitempty"`
	FrontSide    *PassportFile   `json:"front_side,omitempty"`
	ReverseSide  *PassportFile   `json:"reverse_side,omitempty"`
	Selfie       *PassportFile   `json:"selfie,omitempty"`
	Translations *[]PassportFile `json:"translations,omitempty"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     string `json:"file_date"`
}

type PassportElementError struct {
	PassportElementErrorInterface `json:"passport_element_error_interface"`
}

type PassportElementErrorInterface interface {
	passportElementErrorContract()
}

func (p PassportElementError) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.PassportElementErrorInterface)
}

func (p *PassportElementError) UnmarshalJSON(data []byte) error {
	var raw struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Source {
	case "data":
		tmp := PassportElementErrorDataField{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PassportElementErrorInterface = tmp
	case "file":
		tmp := PassportElementErrorFile{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PassportElementErrorInterface = tmp
	case "files":
		tmp := PassportElementErrorFiles{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PassportElementErrorInterface = tmp
	case "front_side":
		tmp := PassportElementErrorFrontSide{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PassportElementErrorInterface = tmp
	case "reverse_side":
		tmp := PassportElementErrorReverseSide{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		p.PassportElementErrorInterface = tmp
	default:
		return errors.New("type must be data, file, files, front_side or reverse_side")
	}
	return nil
}

type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

func (p PassportElementErrorDataField) passportElementErrorContract() {}

func (p PassportElementErrorDataField) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.DataHash) == "" {
		return errors1.ErrInvalidParam("data_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	if strings.TrimSpace(p.FieldName) == "" {
		return errors1.ErrInvalidParam("field_name parameter can't be empty")
	}
	return nil
}

type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (p PassportElementErrorFile) passportElementErrorContract() {}

func (p PassportElementErrorFile) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.FileHash) == "" {
		return errors1.ErrInvalidParam("file_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (p PassportElementErrorFiles) passportElementErrorContract() {}

func (p PassportElementErrorFiles) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if len(p.FileHashes) == 0 {
		return errors1.ErrInvalidParam("file_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (p PassportElementErrorFrontSide) passportElementErrorContract() {}

func (p PassportElementErrorFrontSide) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.FileHash) == "" {
		return errors1.ErrInvalidParam("file_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (p PassportElementErrorReverseSide) passportElementErrorContract() {}

func (p PassportElementErrorReverseSide) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.FileHash) == "" {
		return errors1.ErrInvalidParam("file_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (p PassportElementErrorSelfie) passportElementErrorContract() {}

func (p PassportElementErrorSelfie) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.FileHash) == "" {
		return errors1.ErrInvalidParam("file_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (p PassportElementErrorTranslationFile) passportElementErrorContract() {}

func (p PassportElementErrorTranslationFile) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.FileHash) == "" {
		return errors1.ErrInvalidParam("file_hash parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (p PassportElementErrorTranslationFiles) passportElementErrorContract() {}

func (p PassportElementErrorTranslationFiles) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if len(p.FileHashes) == 0 {
		return errors1.ErrInvalidParam("file_hashes parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}

type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

func (p PassportElementErrorUnspecified) passportElementErrorContract() {}

func (p PassportElementErrorUnspecified) Validate() error {
	if strings.TrimSpace(p.Type) == "" {
		return errors1.ErrInvalidParam("type parameter can't be empty")
	}
	if strings.TrimSpace(p.ElementHash) == "" {
		return errors1.ErrInvalidParam("element_hashe parameter can't be empty")
	}
	if strings.TrimSpace(p.Message) == "" {
		return errors1.ErrInvalidParam("message parameter can't be empty")
	}
	return nil
}
