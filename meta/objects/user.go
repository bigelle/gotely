package objects

type User struct {
	id                      int64
	firstName               string
	isBot                   bool
	lastName                string
	userName                string
	languageCode            string
	canJoinGroups           bool
	canReadAllGroupMessages bool
	supportInlineQueries    bool
	isPremium               bool
	addedToAttachmentMenu   bool
	canConnectToBusiness    bool
	hasMainWebApp           bool
}
