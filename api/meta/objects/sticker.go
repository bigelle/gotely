package objects

type Sticker struct {
	fileId           string
	fileUniqueId     string
	Type             string
	width            int
	height           int
	thumbnail        PhotoSize
	fileSize         int
	emoji            string
	setName          string
	maskPosition     MaskPosition
	isAnimated       bool
	isVideo          bool
	premiumAnimation File
	customEmojiId    string
	needsRepainting  bool
}
