package objects

type PassportData struct {
	Data        []EncryptedPassportElement
	Credentials EncryptedCredentials
}

type EncryptedPassportElement struct {
	Type         string
	Data         string
	PhoneNumber  string
	Email        string
	Files        []PassportFile
	FrontSide    PassportFile
	ReverseSide  PassportFile
	Selfie       PassportFile
	Translations []PassportFile
	Hash         string
}

type EncryptedCredentials struct {
	Data   string
	Hash   string
	Secret string
}

type PassportFile struct {
	FileUniqueId string
	FileSize     int
	FileDate     string
}

//TODO: dataerrors
