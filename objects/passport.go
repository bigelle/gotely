package objects

import (
	"fmt"

	"github.com/bigelle/gotely"
)

// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Array with information about documents and other Telegram Passport elements that was shared with the bot
	Data []EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data
	Credentials EncryptedCredentials `json:"credentials"`
}

// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// File size in bytes
	FileSize int `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate string `json:"file_date"`
}

// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	//Element type.
	//One of “personal_details”, “passport”, “driver_license”, “identity_card”,
	//“internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”,
	//“passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`
	//Optional. Base64-encoded encrypted Telegram Passport element data provided by the user;
	//available only for “personal_details”, “passport”, “driver_license”,
	//“identity_card”, “internal_passport” and “address” types.
	//Can be decrypted and verified using the accompanying EncryptedCredentials.
	Data *string `json:"data,omitempty"`
	// Optional. User's verified phone number; available only for “phone_number” type
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Optional. User's verified email address; available only for “email” type
	Email *string `json:"email,omitempty"`
	//Optional. Array of encrypted files with documents provided by the user;
	//available only for “utility_bill”, “bank_statement”, “rental_agreement”,
	//“passport_registration” and “temporary_registration” types.
	//Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Files *[]PassportFile `json:"files,omitempty"`
	// Optional. Encrypted file with the front side of the document, provided by the user;
	// available only for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`
	// Optional. Encrypted file with the reverse side of the document, provided by the user;
	// available only for “driver_license” and “identity_card”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	// Optional. Encrypted file with the selfie of the user holding a document, provided by the user;
	// available if requested for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`
	//Optional. Array of encrypted files with translated versions of documents provided by the user;
	//available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”,
	//“utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types.
	//Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Translations *[]PassportFile `json:"translations,omitempty"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

// Describes data required for decrypting and authenticating EncryptedPassportElement.
// See https://core.telegram.org/passport#receiving-information
// for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload,
	// data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

// This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user.
// It should be one of:
//
// - PassportElementErrorDataField
//
// - PassportElementErrorFrontSide
//
// - PassportElementErrorReverseSide
//
// - PassportElementErrorSelfie
//
// - PassportElementErrorFile
//
// - PassportElementErrorFiles
//
// - PassportElementErrorTranslationFile
//
// - PassportElementErrorTranslationFiles
//
// - PassportElementErrorUnspecified
type PassportElementError interface {
	GetPassportElementErrorSource() string
	Validate() error
}

// Represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Error source, must be data
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the error,
	// one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`
	// Name of the data field which has the error
	FieldName string `json:"field_name"`
	// Base64-encoded data hash
	DataHash string `json:"data_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorDataField) GetPassportElementErrorSource() string {
	return "data"
}

func (p PassportElementErrorDataField) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "data" {
		err = append(err, fmt.Errorf("source must be 'data'"))
	}
	types := map[string]struct{}{
		"personal_details":  {},
		"passport":          {},
		"driver_license":    {},
		"identity_card":     {},
		"internal_passport": {},
		"address":           {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf(
			"type parameter accepts only one of one of 'personal_details', 'passport”, 'driver_license”, 'identity_card”, 'internal_passport”, 'address'",
		))
	}
	if p.DataHash == "" {
		err = append(err, fmt.Errorf("data_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if p.FieldName == "" {
		err = append(err, fmt.Errorf("field_name parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Error source, must be front_side
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorFrontSide) GetPassportElementErrorSource() string {
	return "front_side"
}

func (p PassportElementErrorFrontSide) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "front_side" {
		err = append(err, fmt.Errorf("source must be 'front_side'"))
	}
	types := map[string]struct{}{
		"passport":          {},
		"driver_license":    {},
		"identity_card":     {},
		"internal_passport": {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf("type parameter accepts only one of 'passport', 'driver_license', 'identity_card', 'internal_passport'"))
	}
	if p.FileHash == "" {
		err = append(err, fmt.Errorf("file_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Error source, must be reverse_side
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorReverseSide) GetPassportElementErrorSource() string {
	return "reverse_side"
}

func (p PassportElementErrorReverseSide) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "reverse_side" {
		err = append(err, fmt.Errorf("source must be 'reverse_side'"))
	}
	types := map[string]struct{}{
		"driver_license": {},
		"identity_card":  {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf("type parameter accepts only one of 'driver_license', 'identity_card'"))
	}
	if p.FileHash == "" {
		err = append(err, fmt.Errorf("file_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Error source, must be selfie
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue,
	// one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorSelfie) GetPassportElementErrorSource() string {
	return "selfie"
}

func (p PassportElementErrorSelfie) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "selfie" {
		err = append(err, fmt.Errorf("source must be 'selfie'"))
	}
	types := map[string]struct{}{
		"passport":          {},
		"driver_license":    {},
		"identity_card":     {},
		"internal_passport": {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf("type parameter accepts only one of 'passport', 'driver_license', 'identity_card', 'internal_passport'"))
	}
	if p.FileHash == "" {
		err = append(err, fmt.Errorf("file_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	// Error source, must be file
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue,
	// one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// Base64-encoded file hash
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorFile) GetPassportElementErrorSource() string {
	return "file"
}

func (p PassportElementErrorFile) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "file" {
		err = append(err, fmt.Errorf("source must be 'file'"))
	}
	types := map[string]struct{}{
		"utility_bill":           {},
		"bank_statement":         {},
		"rental_agreement":       {},
		"passport_registration":  {},
		"temporary_registration": {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf(
			"type parameter accepts only one of 'utility_bill', 'bank_statement', 'rental_agreement', 'passport_registration', 'temporary_registration'",
		))
	}
	if p.FileHash == "" {
		err = append(err, fmt.Errorf("file_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// Error source, must be files
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue,
	// one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorFiles) GetPassportElementErrorSource() string {
	return "files"
}

func (p PassportElementErrorFiles) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "files" {
		err = append(err, fmt.Errorf("source must be 'files'"))
	}
	types := map[string]struct{}{
		"utility_bill":           {},
		"bank_statement":         {},
		"rental_agreement":       {},
		"passport_registration":  {},
		"temporary_registration": {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf(
			"type parameter accepts only one of 'utility_bill', 'bank_statement', 'rental_agreement', 'passport_registration', 'temporary_registration'",
		))
	}
	if len(p.FileHashes) == 0 {
		err = append(err, fmt.Errorf("file_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Error source, must be translation_file
	Source string `json:"source"`
	//Type of element of the user's Telegram Passport which has the issue,
	//one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”,
	//“bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// Base64-encoded file hash
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorTranslationFile) GetPassportElementErrorSource() string {
	return "translation_file"
}

func (p PassportElementErrorTranslationFile) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "translation_file" {
		err = append(err, fmt.Errorf("source must be 'translation_file'"))
	}
	types := map[string]struct{}{
		"passport":               {},
		"driver_license":         {},
		"identity_card":          {},
		"internal_passport":      {},
		"utility_bill":           {},
		"bank_statement":         {},
		"rental_agreement":       {},
		"passport_registration":  {},
		"temporary_registration": {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf(
			"type parameter accepts only one of 'passport', 'driver_license', 'identity_card', 'internal_passport', 'utility_bill', "+
				"'bank_statement', 'rental_agreement', 'passport_registration', 'temporary_registration'",
		))
	}
	if p.FileHash == "" {
		err = append(err, fmt.Errorf("file_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Error source, must be translation_files
	Source string `json:"source"`
	//Type of element of the user's Telegram Passport which has the issue,
	//one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”,
	//“bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorTranslationFiles) GetPassportElementErrorSource() string {
	return "translation_files"
}

func (p PassportElementErrorTranslationFiles) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "translation_files" {
		err = append(err, fmt.Errorf("source must be 'translation_files'"))
	}
	types := map[string]struct{}{
		"passport":               {},
		"driver_license":         {},
		"identity_card":          {},
		"internal_passport":      {},
		"utility_bill":           {},
		"bank_statement":         {},
		"rental_agreement":       {},
		"passport_registration":  {},
		"temporary_registration": {},
	}
	if _, ok := types[p.Type]; !ok {
		err = append(err, fmt.Errorf(
			"type parameter accepts only one of 'passport', 'driver_license', 'identity_card', 'internal_passport', 'utility_bill', "+
				"'bank_statement', 'rental_agreement', 'passport_registration', 'temporary_registration'",
		))
	}
	if len(p.FileHashes) == 0 {
		err = append(err, fmt.Errorf("file_hashes parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	// Error source, must be unspecified
	Source string `json:"source"`
	// Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type"`
	// Base64-encoded element hash
	ElementHash string `json:"element_hash"`
	// Error message
	Message string `json:"message"`
}

func (p PassportElementErrorUnspecified) GetPassportElementErrorSource() string {
	return "unspecified"
}

func (p PassportElementErrorUnspecified) Validate() error {
	var err gotely.ErrFailedValidation
	if p.Source != "unspecified" {
		err = append(err, fmt.Errorf("source must be 'unspecified'"))
	}
	if p.Type == "" {
		err = append(err, fmt.Errorf("type parameter can't be empty"))
	}
	if p.ElementHash == "" {
		err = append(err, fmt.Errorf("element_hash parameter can't be empty"))
	}
	if p.Message == "" {
		err = append(err, fmt.Errorf("message parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}
