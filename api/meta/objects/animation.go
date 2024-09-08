package objects

type Animation struct {
	fileId       string
	fileUniqueId string
	width        int
	height       int
	duration     int
	thumbnail    PhotoSize
	fileName     string
	mimeType     string
	fileSize     int64
}
