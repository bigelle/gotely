package objects

type VideoNote struct {
	fileId       string
	fileUniqueId string
	length       int
	duration     int
	thumbnail    PhotoSize
	fileSize     int
}
