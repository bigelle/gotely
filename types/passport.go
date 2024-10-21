package types

import (
	"encoding/json"
	"fmt"

	"github.com/bigelle/tele.go/internal/assertions"
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

func (p PassportElementError) UnmarshalJSON(data []byte) error {
	var raw struct {
		Source     string          `json:"source"`
		Attributes json.RawMessage `json:"attributes"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Source {
	case "data":
		p.PassportElementErrorInterface = new(PassportElementErrorDataField)
	case "file":
		p.PassportElementErrorInterface = new(PassportElementErrorFile)
	case "files":
		p.PassportElementErrorInterface = new(PassportElementErrorFiles)
	case "front_side":
		p.PassportElementErrorInterface = new(PassportElementErrorFrontSide)
	case "reverse_side":
		p.PassportElementErrorInterface = new(PassportElementErrorReverseSide)
	default:
		return fmt.Errorf("Unrecognized type: %T", p.PassportElementErrorInterface)
	}
	return json.Unmarshal(raw.Attributes, p.PassportElementErrorInterface)
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
	if assertions.IsStringEmpty(p.FieldName) {
		return fmt.Errorf("FieldName parameter can't be empty")
	}
	if assertions.IsStringEmpty(p.DataHash) {
		return fmt.Errorf("DataHash parameter can't be empty")
	}
	if assertions.IsStringEmpty(p.Message) {
		return fmt.Errorf("Message parameter can't be empty")
	}
	if assertions.IsStringEmpty(p.Type) {
		return fmt.Errorf("Type parameter can't be empty")
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
	if err := assertions.ParamNotEmpty(p.FileHash, "FileHash"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if assertions.IsSliceEmpty(p.FileHashes) {
		return fmt.Errorf("FileHashes parameter can't be empty")
	}
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.FileHash, "FileHash"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.FileHash, "FileHash"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.FileHash, "FileHash"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.FileHash, "FileHash"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if assertions.IsSliceEmpty(p.FileHashes) {
		return fmt.Errorf("FileHashes parameter can't be empty")
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
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
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.ElementHash, "ElementHash"); err != nil {
		return err
	}
	if err := assertions.ParamNotEmpty(p.Message, "Message"); err != nil {
		return err
	}
	return nil
}
