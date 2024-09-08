package objects

const (
	MENTION     = "mention"
	HASHTAG     = "hashtag"
	CASHTAG     = "cashtag"
	BOTCOMMAND  = "bot_command"
	URL         = "url"
	EMAIL       = "email"
	PHONENUMBER = "phone_number"
	BOLD        = "bold"
	ITALIC      = "italic"
	CODE        = "code"
	PRE         = "pre"
	TEXTLINK    = "text_link"
	TEXTMENTION = "text_mention"
	SPOILER     = "spoiler"
)

type MessageEntity struct {
	Type          string
	Offset        int
	Length        int
	Url           string
	User          User
	Language      string
	CustomEmojiId string
	Text          string
}

func (m MessageEntity) ComputeText(message string) {
	if message != "" {
		m.Text = message[m.Offset : m.Offset+m.Length]
	}
}
