package objects

type Document struct {
	fileId       string
	fileUniqueId string
	thumbnail    PhotoSize
	fileName     string
	mimeType     string
	fileSize     int64
}
